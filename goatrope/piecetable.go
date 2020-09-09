package goatrope

import (
)


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
func (pt *PieceTable) Load(cutsize int64) {
	pt.Pieces = append(
		pt.Pieces,
		Piece{
			Original: true,
			Start:    0,
			Size:     cutsize,
		},
	)
	pt.Index += cutsize
}

// Size is the total size of pieces
func (pt *PieceTable) Size() int64 {
	totalsz := int64(0)
	for i := 0; i < len(pt.Pieces); i++ {
		totalsz += pt.Pieces[i].Size
	}
	return totalsz
}


// Cut from the data, inverse of Insert
// Prerequisites: set Index to lowest byte to cut
//
// index names:
//
//       lo  cutlo  cuthi  hi
//       [--------sz--------]
//           [---cutsz--]
//  [-----------totalsz---------]
//
// lo and hi denote the boundaries for the current piece
// cutlo and cuthi denote the boundaries that we want cut out
// these two pairs of indices are in absolute index units
//
// sz and cutsz are lengths 
// 
// MODIFICATIONS TO Start must be RELATIVE
//
func (pt *PieceTable) Cut(cutsize int64) {
	lo := int64(0)
	for i := 0; i < len(pt.Pieces); i++ {
		hi := lo + pt.Pieces[i].Size
		cutlo := pt.Index
		cuthi := pt.Index+cutsize
		sz := hi - lo
		if lo == cutlo && cuthi == hi {
			// snap both
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
				// insert-before idiom
				pt.Pieces = append(
					pt.Pieces[0:i],
					pt.Pieces[i+1:]...,
				)
			}
		} else if lo == cutlo && cuthi < hi {
			// snap lo
			pt.Pieces[i].Size = (sz - cutsize)
			pt.Pieces[i].Start += cutsize
		} else if lo < cutlo && cuthi == hi {
			// snap hi
			pt.Pieces[i].Size = (sz - cutsize)
		} else if lo < cutlo && cuthi < hi {
			// no snap - inside
			topPiece := Piece{
				Original: pt.Pieces[i].Original,
				Start:    pt.Pieces[i].Start + (cuthi-lo),
				Size:     (hi - cuthi),
			}
			// split piece sizes sum to:
			//   (hi - cuthi) + (cutlo - lo)
			// = (hi - lo) - (cuthi - cutlo)
			// = (sz - cutsize)
			pt.Pieces[i].Size = (cutlo - lo)
			// append-after idiom
			pt.Pieces = append(
				pt.Pieces[0:i+1],
				append(
					[]Piece{topPiece},
					pt.Pieces[i+1:]...,
				)...,
			)
		} else if (cutlo < lo) && (lo <= cuthi && cuthi <= hi) {
			// cuthi in range
			// cut THIS chunk
			pt.Index = lo
			pt.Cut(cuthi - lo)
			// cut REMAINING chunks that preceded this one
			pt.Index = cutlo
			pt.Cut(lo - cutlo)
			// cuts out (cutlo - cuthi) bytes when finished
			// when working back to cut out chunks,
			// will snap to the new hi
		} else {
			// Not yet in range... we ONLY continue in this case
			lo = lo + pt.Pieces[i].Size
			continue
		}
		return
	}
}

// Insert calculates pieces,
// Prerequisite: Load
// Prerequisite: set Index to insert-before location,
// where index is set past the last byte to append
//       lo  cutlo  cuthi  hi
//       [--------sz--------]
//           [---cutsz--]
//  [-----------totalsz---------]
//
func (pt *PieceTable) Insert(cutsize int64) {
	totalsz := pt.Size()
	// This is the append-to-end case
	if totalsz == pt.Index {
		pt.Pieces = append(
			pt.Pieces,
			Piece{
				Original: false,
				Start:    pt.ModsSize,
				Size:     cutsize,
			},
		)
	} else {
		// Insert for all other cases
		lo := int64(0)
		newPiece := Piece{
			Original: false,
			Start:    pt.ModsSize,
			Size:     cutsize,
		}
		for i := 0; i < len(pt.Pieces); i++ {
			hi := lo + pt.Pieces[i].Size
			cutlo := pt.Index
			// If we are on a piece boundary (lo or hi), then no split required
			if lo == cutlo {
				// insert-before position i
				pt.Pieces = append(
					pt.Pieces[0:i],
					append(
						[]Piece{newPiece},
						pt.Pieces[i:]...,
					)...,
				)
				break
			} else if lo < cutlo && cutlo < hi {
				// SplitPiece invariants:
				// sizes total original size
				// top start ahead of original start by n
				topPiece := Piece{
					Original: pt.Pieces[i].Original,
					Start:    pt.Pieces[i].Start + (cutlo - lo),
					Size:     pt.Pieces[i].Size - (cutlo - lo),
				}
				pt.Pieces[i].Size = (cutlo - lo)
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
	pt.ModsSize += cutsize
	// May be modified by caller
	pt.Index += cutsize
}

