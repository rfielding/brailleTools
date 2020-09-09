package goatrope_test

import (
	"github.com/rfielding/teeny/goatrope"
	"testing"
)

// checkPieces tests the integrity of the PieceTable vs what we expected
func checkPieces(t *testing.T, item int, pt goatrope.PieceTable, szstarts []goatrope.Piece) {
	if len(pt.Pieces) != len(szstarts) {
		t.Logf("%d: expect %d pieces got %d: %s", item, len(szstarts), len(pt.Pieces), goatrope.ToJson(pt))
		t.Fail()
	}
	for i := 0; i < len(szstarts); i++ {
		if pt.Pieces[i].Original != szstarts[i].Original {
			t.Logf(
				"%d: file mismatch at %d: got %t expected %t",
				item,
				i,
				pt.Pieces[i].Original,
				szstarts[i].Original,
			)
			t.FailNow()
		}
		if pt.Pieces[i].Size != szstarts[i].Size {
			t.Logf(
				"%d: size mismatch at %d: got %d expected %d",
				item,
				i,
				pt.Pieces[i].Size,
				szstarts[i].Size,
			)
			t.FailNow()
		}
		if pt.Pieces[i].Start != szstarts[i].Start {
			t.Logf(
				"%d: start mismatch at %d: got %d expected %d",
				item,
				i,
				pt.Pieces[i].Start,
				szstarts[i].Start,
			)
			t.FailNow()
		}
	}
}

// TestPieceTableInserts does basic integrity checks of the insert algorithm
func TestPieceTableInserts(t *testing.T) {

	// Just test the PieceTable structure
	// without writing data yet
	pt := goatrope.PieceTable{}
	pt.Load(100)

	checkPieces(t, 1, pt, []goatrope.Piece{
		{true, 0, 100},
	})

	// Append from where we were
	pt.Insert(50)

	checkPieces(t, 2, pt, []goatrope.Piece{
		{true, 0, 100},
		{false, 0, 50},
	})

	// Append to the beginning of the file
	pt.Index = 0
	pt.Insert(25)

	checkPieces(t, 3, pt, []goatrope.Piece{
		{false, 50, 25},
		{true, 0, 100},
		{false, 0, 50},
	})

	// Append in between pieces (no split required)
	pt.Index = 25
	pt.Insert(26)

	checkPieces(t, 4, pt, []goatrope.Piece{
		{false, 50, 25},
		{false, 75, 26},
		{true, 0, 100},
		{false, 0, 50},
	})

	// Split part 1 (length 26) in half
	pt.Index = 25 + 13
	pt.Insert(10)

	checkPieces(t, 5, pt, []goatrope.Piece{
		{false, 50, 25},
		{false, 75, 13},
		{false, 101, 10},
		{false, 88, 13},
		{true, 0, 100},
		{false, 0, 50},
	})
}

// TestPieceTableDeletes checks integrity of delete algorithm
func TestPieceTableDeletes(t *testing.T) {
	pt := goatrope.PieceTable{}
	pt.Load(100)
	pt.Insert(30)
	pt.Insert(20)
	pt.Insert(12)
	pt.Insert(32)
	pt.Insert(50)

	// Initial state
	checkPieces(t, 1, pt, []goatrope.Piece{
		{true, 0, 100},
		{false, 0, 30},
		{false, 30, 20},
		{false, 50, 12},
		{false, 62, 32},
		{false, 94, 50},
	})

	// lo == cutlo < cuthi == hi
	pt.Index = 130
	pt.Cut(20)
	checkPieces(t, 2, pt, []goatrope.Piece{
		{true, 0, 100},
		{false, 0, 30},
		{false, 50, 12},
		{false, 62, 32},
		{false, 94, 50},
	})

	// lo == cutlo < cuthi < hi
	pt.Index = 130
	pt.Cut(1)
	checkPieces(t, 3, pt, []goatrope.Piece{
		{true, 0, 100},
		{false, 0, 30},
		{false, 51, 11},
		{false, 62, 32},
		{false, 94, 50},
	})

	// lo < cutlo < cuthi == hi
	pt.Index = 140
	pt.Cut(1)
	checkPieces(t, 3, pt, []goatrope.Piece{
		{true, 0, 100},
		{false, 0, 30},
		{false, 51, 10},
		{false, 62, 32},
		{false, 94, 50},
	})

	// cutlo < lo < cuthi = hi
	// snap on lo cut span
	pt.Index = 130
	pt.Cut(42)
	checkPieces(t, 3, pt, []goatrope.Piece{
		{true, 0, 100},
		{false, 0, 30},
		{false, 94, 50},
	})


	// cutlo < lo < cuthi = hi
	// adjacent spans
	pt.Index = 80
	pt.Cut(60)
	checkPieces(t, 3, pt, []goatrope.Piece{
		{true, 0, 80},
		{false, 104, 40},
	})
}
