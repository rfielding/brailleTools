package piecetable_test

import (
	"fmt"
	"os"
	"io"
	"testing"
	"github.com/rfielding/teeny/piecetable"
)

func TestPtOpen(t *testing.T) {
	pt := piecetable.New()
	defer pt.Close()

	var err error
	// Load it up
	err = pt.Load("sample.txt")
	if err != nil {
		panic(err)
	}
	// Seek into the middle somewhere
	_, err = pt.Seek(35, io.SeekStart)
	if err != nil {
		panic(err)
	}
	// Seek to known original text
	buffer := make([]byte, 54)
	_, err = pt.Read(buffer)
	if err != nil {
		panic(err)
	}
	expected := "In the beginning God created the heaven and the earth."
	if expected != string(buffer) {
		panic(fmt.Sprintf("unexpected string:\n%s\nvs\n%s", expected, string(buffer)))
	}

	// Write the whole thing out
	_, err = pt.Store("sample.txt.copied.txt")
	if err != nil {
		panic(err)
	}

	// Verify that they are same length
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
