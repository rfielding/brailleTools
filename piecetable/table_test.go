package piecetable_test

import (
	"fmt"
	"github.com/rfielding/teeny/piecetable"
	"io"
	"os"
	"testing"
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
		panic(
			fmt.Sprintf(
				"lengths do not match %d vs %d", 
				stat1.Size(), 
				stat2.Size(),
			),
		)
	}
}

func xTestPtAppend(t *testing.T) {
	pt := piecetable.New()
	defer pt.Close()

	var err error
	// Load it up
	err = pt.Load("sample.txt")
	if err != nil {
		panic(err)
	}
	sz := pt.Size()
	t.Logf("seek to the end at size: %d", sz)
	_, err = pt.Seek(sz, io.SeekStart)
	if err != nil {
		panic(err)
	}
	_, err = pt.Write([]byte("123456"))
	if err != nil {
		panic(err)
	}
	// Write the whole thing out
	_, err = pt.Store("sample.txt.appended.txt")
	if err != nil {
		panic(err)
	}

	stat1, err1 := os.Stat("sample.txt")
	stat2, err2 := os.Stat("sample.txt.appended.txt")
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	if stat1.Size()+6 != stat2.Size() {
		panic(
			fmt.Sprintf(
				"lengths do not match: %d vs %d",
				stat1.Size(),
				stat2.Size(),
			),
		)
	}
}

