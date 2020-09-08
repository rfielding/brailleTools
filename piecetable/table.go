package piecetable

import (
	"io"
	"os"
	"fmt"
)

/*
 Original - immutable part of the file that is on disk
 Add - append-only edit data
 Table - mutable structure to organize pieces

 This is effectively a character device:

 - It behaves like a read/write/seek device
 - Writes will insert before the current byte
 - We can seek past the last byte to append after

 Looking for carriage returns, we can implement things
 such as cursor down or up one line, which are convenient
 seek methods. 
*/

// Seeking, handling linebreaks
type Cursor interface {
	CursorUp() (int64, error)
	CursorDown() (int64, error)
	CursorLeft() (int64, error)
	CursorRight() (int64, error)
}

// Load and store
type LoadStore interface {
	// Wipe out all state (including undo capability) and load from file
	Load(name string) error
	// Temporarily seek to beginning and write file out, restoring seek
	Store(name string) (int64, error)
}

// Make this a valid character stream
type Piecetable interface {
	io.ReadCloser
	io.WriteCloser
	io.Seeker
	Cursor
	LoadStore
}

// Organize all changes IN MEMORY
type Piece struct {
	IsOriginal bool
	Start      int64
	Size       int64
}

type Pt struct {
	// Assume that the original is enormous
	// Do not hold it in memory.  ASSUME
	// that this file is NOT concurrently edited with us.
	Loaded   string
	Original *os.File
	Appended []byte // Store and Load to flush this out of memory!
	// Use in-memory uncommitted modifications,
	// which assumes small inserts, and frequent saves
	Pieces   []Piece
	SeekedTo int64
}

func New() *Pt {
	return &Pt{}
}

func (p *Pt) Load(name string) error {
	// Freshly open the original
	if p.Original != nil {
		p.Close()
	}
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	p.Loaded = name
	p.Original = f
	// Clean out data structure 
	s, err := os.Stat(name)
	if err != nil {
		return err
	}
	p.Pieces = nil
	p.Pieces = append(p.Pieces, Piece{
		IsOriginal: true,
		Start:      0,
		Size:       s.Size(),
	})
	p.Appended = nil
	p.SeekedTo = 0
	return nil
}

// How large will our file be when we Store it?
func (p *Pt) Size() int64 {
	total := int64(0)
	for i := 0; i < len(p.Pieces); i++ {
		total += p.Pieces[i].Size
	}
	return total
}

// Write is an insert-before at the point we are seeked to.
// To append to the file, p.Seek(p.Size(), io.SeekStart)
// to place the cursor after the last byte
//
// To write at the beginning of the file: p.Seek(0, io.SeekStart)
// before writing.
//
// Writing into the middle of the file will split pieces.
func (p *Pt) Write(b []byte) (int, error) {
	// Any new piece appends edit data
	newPiece := Piece{
		IsOriginal: false,
		Start: int64(len(p.Appended)),
		Size: int64(len(b)),
	}
	if p.Size() == p.SeekedTo {
		fmt.Printf("seek to the end!")
		// If we are seeked just past the end of the file
		p.Pieces = append(p.Pieces, newPiece)
	} else if 0 == p.SeekedTo {
		// If we are at first byte
		p.Pieces = append([]Piece{newPiece},p.Pieces...)
	} else {
		// If we seek to the middle, we have to split something
		lo := int64(0)
		for i := 0; i < len(p.Pieces); i++ {
			sz := p.Pieces[i].Size
			hi := lo + sz
			if lo <= p.SeekedTo && p.SeekedTo < hi {
				// Insert split piece
				p.Pieces = append(
					p.Pieces[0:i],
					append(
						[]Piece{
							p.Pieces[i],
							newPiece,Piece{},
						},
						p.Pieces[i:]...,
					)...,
				)
				// Split i around newPiece
				n := p.SeekedTo - lo
				p.Pieces[i+2].IsOriginal = p.Pieces[i].IsOriginal
				// Split the start
				p.Pieces[i+2].Start = p.Pieces[i].Start + n
				// Split the size
				p.Pieces[i+2].Size = p.Pieces[i].Size - n
				p.Pieces[i].Size = n
			}
			break
		}
	}
	// Write the data out
	p.Appended = append(p.Appended, b...)
	p.SeekedTo += int64(len(b))
	return len(b), nil
}

func (p *Pt) Read(b []byte) (int, error) {
	if len(p.Pieces) == 0 {
		return 0, io.EOF
	}
	// Find the right piece and spot to seek to
	remaining := p.SeekedTo
	for i := range p.Pieces {
		if p.Pieces[i].Size < remaining {
			// Skip to next piece
			remaining -= p.Pieces[i].Size
		} else {
			// Note: Size is int64 for Original, but int for appended
			if p.Pieces[i].IsOriginal {
				p.Original.Seek(remaining+p.Pieces[i].Start, io.SeekStart)
				read, err := p.Original.Read(b)
				if err == nil {
					p.SeekedTo += int64(read)
					return read, nil
				}
				return read, err
			} else {
				// Read as many bytes as possible
				copyMax := len(b)
				remainder := int(remaining)
				sz := int(p.Pieces[i].Size)
				st := int(p.Pieces[i].Start)
				if sz-remainder < copyMax {
					copyMax = sz - remainder
				}
				if copyMax == 0 {
					return 0, io.EOF
				}
				// Copy bytes into the read
				for n := 0; n < copyMax; n++ {
					b[n] = p.Appended[st+i+remainder]
				}
				p.SeekedTo += int64(copyMax)
				return copyMax, nil
			}
		}
	}
	return p.Original.Read(b)
}

func (p *Pt) Store(name string) (int64, error) {
	at := p.SeekedTo

	// Open file to create
	f, err := os.Create(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Rewind and write
	p.Seek(0, io.SeekStart)
	written, err := io.Copy(f, p)
	if err != nil {
		return written, err
	}

	// Restore cursur position
	at, err = p.Seek(at, io.SeekStart)
	if err != nil {
		return at, err
	}
	return written, nil
}

func (p *Pt) Close() error {
	err := p.Original.Close()
	p.Original = nil
	p.Pieces = nil
	p.Appended = nil
	return err
}

func (p *Pt) Seek(n int64, whence int) (int64, error) {
	p.SeekedTo = n
	return n, nil
}
