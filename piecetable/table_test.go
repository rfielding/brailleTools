package piecetable_test

import (
	"os"
	"testing"
	"github.com/rfielding/teeny/piecetable"
)

func TestPtOpen(t *testing.T) {
	pt := piecetable.New()
	defer pt.Close()
	pt.Load("sample.txt")
	pt.Store("sample.txt.copied.txt")

	stat1, err1 := os.Stat("sample.txt")
	stat2, err2 := os.Stat("sample.txt.copied.txt")

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	if stat1.Size() != stat2.Size() {
		panic("lengths do not match!")
	}
}
