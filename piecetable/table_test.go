package piecetable_test

import (
	"testing"
	"github.com/rfielding/teeny/piecetable"
)

func TestPtOpen(t *testing.T) {
	pt := piecetable.New()
	defer pt.Close()
	pt.Load("sample.txt")
}
