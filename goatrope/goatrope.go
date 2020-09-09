package goatrope

import (
	"io"
	"os"
)

/*
 A GoatRope is my variant of a PieceTable for bytes

 This is the hard, and most general component
 of making a text editor.  It is kind of like
 extending the normal Unix file interface to
 handle inserts and deletions efficiently.

 It is like a standard *io.File that
 allows for efficient inserts and deletes.

 It is also possible to use the Write method
 to delete from it, by writing in escape
 sequences that perform deletes and backspaces
 on the file rather than writing them IN
 to the file.
*/

// File puts down an interface for what passes as a File for us
type File interface {
	io.Reader
	io.Seeker
	io.Closer
	io.Writer
	Stat() (os.FileInfo, error)
}

// Ensure that indeed a *os.File does implement File
var _ File = &os.File{}

// GoatRope extends File with extra behaviors
type GoatRope interface {
	File
}

// Piece describes how to include data in the stream
type Piece struct {
	Original bool
	Start    int64
	Size     int64
}

// PieceTable structure of a GoatRope
// is the bit that needs manipulation
// and testing, but not actually any of the
// file bytes
type PieceTable struct {
	Pieces   []Piece
	Index    int64
	Original File
	Mods     File
	ModsSize int64
}

// Load original file to represent it in PieceTable
// Prequisite: a default PieceTable value
func (pt *PieceTable) Load(size int64) {
	pt.Pieces = append(
		pt.Pieces,
		Piece{
			Original: true,
			Start:    0,
			Size:     size,
		},
	)
	pt.Index += size
}

// Size is the total size of pieces
func (pt *PieceTable) Size() int64 {
	total := int64(0)
	for i := 0; i < len(pt.Pieces); i++ {
		total += pt.Pieces[i].Size
	}
	return total
}

// Cut from the data, inverse of Insert
// Prerequisites: set Index to lowest byte to cut
func (pt *PieceTable) Cut(psize int64) {
	if len(pt.Pieces) == 0 {
		return
	}
	lo := int64(0)
	for i := 0; i < len(pt.Pieces); i++ {
		hi := lo + pt.Pieces[i].Size
		if lo == pt.Index && pt.Index+psize == hi {
			// it's an exact range match
			if len(pt.Pieces) == 1 {
				// just one piece... get rid of it
				pt.Pieces = nil
			} else if i+1 == len(pt.Pieces) {
				// last piece... chop off last one
				pt.Pieces = pt.Pieces[0 : len(pt.Pieces)-1]
			} else if i == 0 {
				// first piece... chop it off
				pt.Pieces = pt.Pieces[1:]
			} else {
				// More than two pieces, not first or last
				pt.Pieces = append(
					pt.Pieces[0:i],
					pt.Pieces[i+1:]...,
				)
			}
			break
		} else if lo == pt.Index && pt.Index+psize < hi {
			// matches lower boundary, but less than upper
			pt.Pieces[i].Size -= psize
			pt.Pieces[i].Start += (pt.Pieces[i].Size - psize)
			break
		} else if lo < pt.Index && pt.Index+psize == hi {
			// matches upper boundary
			pt.Pieces[i].Size -= psize
			break
		}
		lo = lo + pt.Pieces[i].Size
	}
}

// Insert calculates pieces,
// Prerequisite: Load
// Prerequisite: set Index to insert-before location,
// where index is set past the last byte to append
func (pt *PieceTable) Insert(psize int64) {
	sz := pt.Size()
	// This is the append-to-end case
	if sz == pt.Index {
		pt.Pieces = append(
			pt.Pieces,
			Piece{
				Original: false,
				Start:    pt.ModsSize,
				Size:     psize,
			},
		)
	} else {
		// Insert for all other cases
		lo := int64(0)
		newPiece := Piece{
			Original: false,
			Start:    pt.ModsSize,
			Size:     psize,
		}
		for i := 0; i < len(pt.Pieces); i++ {
			hi := lo + pt.Pieces[i].Size
			// If we are on a piece boundary (lo or hi), then no split required
			if lo == pt.Index {
				// insert-before position i
				pt.Pieces = append(
					pt.Pieces[0:i],
					append(
						[]Piece{newPiece},
						pt.Pieces[i:]...,
					)...,
				)
				break
			} else if lo < pt.Index && pt.Index < hi {
				n := pt.Index - lo
				// SplitPiece invariants:
				// sizes total original size
				// top start ahead of original start by n
				topPiece := Piece{
					Original: pt.Pieces[i].Original,
					Start:    pt.Pieces[i].Start + n,
					Size:     pt.Pieces[i].Size - n,
				}
				pt.Pieces[i].Size = n
				pt.Pieces = append(
					pt.Pieces[0:i+1],
					append(
						[]Piece{newPiece, topPiece},
						pt.Pieces[i+1:]...,
					)...,
				)
				break
			}
			lo += pt.Pieces[i].Size
		}
	}
	// Only tracked internally and set here
	pt.ModsSize += psize
	// May be modified by caller
	pt.Index += psize
}

// MemoryFile always expects
// to be seeked to the end before
// Write
type MemoryFile struct {
}
