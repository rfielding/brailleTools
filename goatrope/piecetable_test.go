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
				"%d: file mismatch at %d: got %t expected %t, vs %s",
				item,
				i,
				pt.Pieces[i].Original,
				szstarts[i].Original,
				goatrope.ToJson(pt),
			)
			t.FailNow()
		}
		if pt.Pieces[i].Size != szstarts[i].Size {
			t.Logf(
				"%d: size mismatch at %d: got %d expected %d vs %s",
				item,
				i,
				pt.Pieces[i].Size,
				szstarts[i].Size,
				goatrope.ToJson(pt),
			)
			t.FailNow()
		}
		if pt.Pieces[i].Start != szstarts[i].Start {
			t.Logf(
				"%d: start mismatch at %d: got %d expected %d vs %s",
				item,
				i,
				pt.Pieces[i].Start,
				szstarts[i].Start,
				goatrope.ToJson(pt),
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
		{false, 50, 51},
		{true, 0, 100},
		{false, 0, 50},
	})

	// Split part 1 (length 26) in half
	pt.Index = 30
	pt.Insert(10)

	checkPieces(t, 5, pt, []goatrope.Piece{
		{false, 50, 30},
		{false, 101, 10},
		{false, 80, 21},
		{true, 0, 100},
		{false, 0, 50},
	})

	pt.Index = 40
	pt.Insert(21)

	checkPieces(t, 6, pt, []goatrope.Piece{
		{false, 50, 30},
		{false, 101, 31},
		{false, 80, 21},
		{true, 0, 100},
		{false, 0, 50},
	})

	pt.Index = 30
	pt.Insert(21)
	checkPieces(t, 7, pt, []goatrope.Piece{
		{false, 50, 30},
		{false, 132, 21},
		{false, 101, 31},
		{false, 80, 21},
		{true, 0, 100},
		{false, 0, 50},
	})
}

// TestPieceTableDeletes checks integrity of delete algorithm
func xTestPieceTableDeletes(t *testing.T) {
	pt := goatrope.PieceTable{}
	pt.Load(100)
	pt.Index = 30
	pt.Insert(30)
	pt.Index = 102
	pt.Insert(20)
	pt.Index = 50
	pt.Insert(12)
	pt.Index = 0
	pt.Insert(32)
	pt.Index = pt.Size()
	pt.Insert(50)
	checkPieces(t, 1, pt, []goatrope.Piece{
		{false, 62, 32},
		{true, 0, 30},
		{false, 0, 20},
		{false, 50, 12},
		{false, 20, 10},
		{true, 30, 42},
		{false, 30, 20},
		{true, 72, 28},
		{false, 94, 50},
	})

	pt.Index = 0
	pt.Cut(32)
	checkPieces(t, 2, pt, []goatrope.Piece{
		{true, 0, 30},
		{false, 0, 20},
		{false, 50, 12},
		{false, 20, 10},
		{true, 30, 42},
		{false, 30, 20},
		{true, 72, 28},
		{false, 94, 50},
	})

	pt.Index = 30
	pt.Cut(20)
	checkPieces(t, 3, pt, []goatrope.Piece{
		{true, 0, 30},
		{false, 50, 12},
		{false, 20, 10},
		{true, 30, 42},
		{false, 30, 20},
		{true, 72, 28},
		{false, 94, 50},
	})

	pt.Index = 20
	pt.Cut(20)
	checkPieces(t, 4, pt, []goatrope.Piece{
		{true, 0, 20},
		{false, 60, 2},
		{false, 20, 10},
		{true, 30, 42},
		{false, 30, 20},
		{true, 72, 28},
		{false, 94, 50},
	})

	pt.Index = 20
	pt.Cut(12)
	checkPieces(t, 5, pt, []goatrope.Piece{
		{true, 0, 20},
		{true, 30, 42},
		{false, 30, 20},
		{true, 72, 28},
		{false, 94, 50},
	})

	pt.Index = 30
	pt.Cut(1)
	checkPieces(t, 6, pt, []goatrope.Piece{
		{true, 0, 20},
		{true, 30, 10},
		{true, 41, 31},
		{false, 30, 20},
		{true, 72, 28},
		{false, 94, 50},
	})

	pt.Index = 109
	pt.Cut(50)
	checkPieces(t, 7, pt, []goatrope.Piece{
		{true, 0, 20},
		{true, 30, 10},
		{true, 41, 31},
		{false, 30, 20},
		{true, 72, 28},
	})

	pt.Index = 20
	pt.Cut(89)
	checkPieces(t, 8, pt, []goatrope.Piece{
		{true, 0, 20},
	})

	pt.Index = 0
	pt.Cut(20)
	checkPieces(t, 9, pt, []goatrope.Piece{})

	pt.Insert(5)
	checkPieces(t, 10, pt, []goatrope.Piece{
		{false, 144, 5},
	})

}
