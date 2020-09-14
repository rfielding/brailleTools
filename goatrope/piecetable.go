package goatrope

import (
	"fmt"
)

// Piece describes how to include data in the stream
type Piece struct {
	Original bool
	Start    int64
	Size     int64
}

func (p *Piece) ToString() string {
	return fmt.Sprintf("(%t, %d, %d)", p.Original, p.Start, p.Size)
}

// PieceTable structure of a GoatRope
// is the bit that needs manipulation
// and testing, but not actually any of the
// file bytes
type PieceTable struct {
	Pieces   []Piece
	Index    int64
	ModsSize int64
}

// Size is the total size of pieces
func (pt *PieceTable) Size() int64 {
	totalsz := int64(0)
	for i := 0; i < len(pt.Pieces); i++ {
		totalsz += pt.Pieces[i].Size
	}
	return totalsz
}

// Load original file to represent it in PieceTable
// Prequisite: a default PieceTable value
func (pt *PieceTable) Load(cutsz int64) {
	pt.Pieces = []Piece{{true, 0, cutsz}}
	pt.Index = cutsz
}

func (pt *PieceTable) idxFind(iStart int) (lo int64, hi int64, cutlo int64, i int) {
	i = iStart
	// Append amongst existing
	cutlo = pt.Index

	// find the current index
	for {
		if i == len(pt.Pieces) {
			panic("reached beyond the end")
		}

		lo = hi
		hi += pt.Pieces[i].Size

		if lo <= cutlo && cutlo <= hi {
			break
		}

		i++
	}
	return lo, hi, cutlo, i
}

func (pt *PieceTable) Insert(cutsz int64) {
	// The new piece will look like this,
	// unless we extend the previous piece
	newPiece := Piece{false, pt.ModsSize, cutsz}
	defer func() {
		pt.ModsSize += cutsz
		pt.Index += cutsz
	}()

	// Appending empty is a trivial case
	if len(pt.Pieces) == 0 {
		pt.Pieces = []Piece{newPiece}
		return
	}

	lo, hi, cutlo, i := pt.idxFind(0)

	if cutlo == hi {
		tail := pt.Pieces[i].Start + pt.Pieces[i].Size
		if !pt.Pieces[i].Original && newPiece.Start == tail {
			pt.Pieces[i].Size += cutsz
			return
		}
		pt.Pieces = append(
			pt.Pieces[0:i+1],
			append(
				[]Piece{newPiece},
				pt.Pieces[i+1:]...,
			)...,
		)
		return
	}

	if lo == cutlo {
		pt.Pieces = append(
			pt.Pieces[0:i],
			append(
				[]Piece{newPiece},
				pt.Pieces[i:]...,
			)...,
		)
		return
	}

	// lo < cutlo.  after i, inser new and split
	n := cutlo - lo
	pt.Pieces = append(
		pt.Pieces[0:i+1],
		append(
			[]Piece{newPiece, pt.Pieces[i]},
			pt.Pieces[i+1:]...,
		)...,
	)
	pt.Pieces[i+2].Start += n
	pt.Pieces[i+2].Size -= n
	pt.Pieces[i].Size = n
}

func (pt *PieceTable) Cut(cutsz int64) {
	// Handle do-nothing situations
	if len(pt.Pieces) == 0 {
		return
	}
	if cutsz == 0 {
		return
	}
	fileSz := pt.Size()
	if pt.Index >= fileSz {
		return
	}
	// Look for matching chunk
	cutlo := pt.Index
	cuthi := cutlo + cutsz
	hi := int64(0)
	for i := 0; i < len(pt.Pieces); i++ {
		lo := hi
		hi += pt.Pieces[i].Size
		inRangeLo := (lo <= cuthi && cuthi < hi)
		inRangeHi := (lo <= cuthi && cuthi == hi)
		inRange := inRangeLo || inRangeHi

		// recurse backwards on last chunk in range
		if inRange && cutlo < lo {
			saved := cutlo
			pt.Index = lo
			pt.Cut(cuthi - lo)
			pt.Index = cutlo
			pt.Cut(lo - cutlo)
			pt.Index = saved
			return
		}

		if inRangeLo {
			if lo < cutlo {
				pt.Pieces = append(
					pt.Pieces[0:i],
					append(
						[]Piece{pt.Pieces[i]},
						pt.Pieces[i:]...,
					)...,
				)
				sz := hi - lo
				n := cutlo - lo
				pt.Pieces[i].Size = n
				pt.Pieces[i+1].Size = sz - n - cutsz
				pt.Pieces[i+1].Start += n + cutsz
			} else if lo == cutlo  {
				pt.Pieces[i].Size -= cutsz
				pt.Pieces[i].Start += cutsz
			}
			return
		}

		if inRangeHi {
			if lo < cutlo {
				pt.Pieces[i].Size -= cutsz
			} else if lo == cutlo {
				pt.Pieces = append(
					pt.Pieces[0:i],
					pt.Pieces[i+1:]...,
				)
			}
			return
		}

	}
}
