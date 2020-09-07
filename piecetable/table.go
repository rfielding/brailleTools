package piecetable

import (
	"os"
	"io"
)

/*
 Original - immutable part of the file that is on disk
 Add - append-only edit data
 Table - mutable structure to organize pieces

 Use ints, which will typically store runes

 Represent it as a:

 */

// Seeking, handling linebreaks
type Cursor interface {
	CursorUp() (int64, error)
	CursorDown() (int64, error)
	CursorLeft() (int64, error)
	CursorRight() (int64, error)
}

// Inserting data before or after
type Insertable interface {
	Insert(b []byte, isAfter bool) error
}

// Load and store
type LoadStore interface {
	Load(name string) error
	Store(name string) (int64, error)
}

type Piecetable interface {
	io.ReadCloser
	io.Seeker
	Insertable
	Cursor
	LoadStore
}

type Pt struct {
	// Assume that the original is enormous
	// Do not hold it on disk.
	Loaded string
	Original *os.File
	// Uncommitted edits are similar,
	// and only needed if edits are made
	Modifications string
	Added *os.File
	// pieces to weave modifictaions into original
	ModificationPieces string
	Pieces *os.File
}

func New() *Pt {
	return &Pt{}
}

func (p *Pt) Load(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	p.Loaded = name
	p.Original = f
	return nil
}

func (p *Pt) Read(b []byte) (int, error) {
	return p.Original.Read(b)
}

func (p *Pt) Store(name string) (int64, error) {
	// Save where we are at in the stream
	at, err := p.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	// Open file to create
	f, err := os.Create(name)
	if err != nil {
		return 0,err
	}
	defer f.Close()

	// Rewind and write
	p.Seek(0, io.SeekStart)
	written, err := io.Copy(f, p)
	if err != nil {
		return written,err
	}

	// Restore cursur position
	at, err = p.Seek(at, io.SeekStart)
	if err != nil {
		return at, err
	}
	return written, nil
}

func (p *Pt) Close() error {
	return p.Original.Close()
}

func (p *Pt) Seek(n int64, whence int) (int64, error) {
	return p.Original.Seek(n, whence)
}
