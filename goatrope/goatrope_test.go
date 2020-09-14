package goatrope_test

import (
	"fmt"
	"github.com/rfielding/teeny/goatrope"
	"io"
	"os"
	"testing"
)

func xTestGoatRope(t *testing.T) {

	g := goatrope.NewGoatRope()
	defer g.Close()
	g.LoadByName("goatrope.go")
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// world\n"))
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// hello\n"))
	g.Seek(26, io.SeekStart)
	g.Write([]byte("// cruel\n"))
	g.Seek(0, io.SeekStart)
	io.Copy(os.Stdout, g)
}

func TestGoatRopeByLine(t *testing.T) {

	g := goatrope.NewGoatRope()
	defer g.Close()
	g.LoadByName("goatrope.go")

	g.Seek(17, io.SeekStart)
	g.Write([]byte("// world\n"))
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// hello\n"))
	g.Seek(0, io.SeekStart)

	at := int64(0)
	line := int64(20)
	_, idx := g.SeekToLine(line)
	linesToShow := 5
	fmt.Printf("start from %d, with insert at 273\n", idx)

	// Make an edit (insert-after)
	g.Seek(273, io.SeekStart)
	g.Write([]byte(" \u001b[31mreally\u001b[0m"))

	// Render the changed screen
	g.Seek(idx, io.SeekStart)
	rdr := g.RuneScanner()
	for linesToShow > 0 {
		lineBytes, _, err := rdr.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("\u001b[33m%04d\u001b0m \u001b[34m(%04d):\u001b[0m %s\n", line, at+idx, string(lineBytes))
		line++
		linesToShow--
		at += int64(len(lineBytes)) + 1
	}
}
