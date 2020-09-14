package goatrope_test

import (
	"fmt"
	"github.com/rfielding/teeny/goatrope"
	"io"
	"os"
	"testing"
)

func TestGoatRope(t *testing.T) {

	g := goatrope.NewGoatRope()
	defer g.Close()
	g.LoadByName("goatrope.go")
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// world\n"))
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// hello\n"))
	g.Seek(0, io.SeekStart)
	io.Copy(os.Stdout, g)
}

func xTestGoatRopeByLine(t *testing.T) {

	g := goatrope.NewGoatRope()
	defer g.Close()
	g.LoadByName("goatrope.go")

	/*
		g, err := os.Open("goatrope.go")
		if err != nil {
			panic(err)
		}
		defer g.Close()
	*/

	g.Seek(17, io.SeekStart)
	g.Write([]byte("// world\n"))
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// hello\n"))
	g.Seek(0, io.SeekStart)

	line := int64(0)
	g.Seek(0, io.SeekStart)
	buffer := make([]byte, 40)
	for {
		start, stop, err := goatrope.Lines(g, buffer, 0, 1)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("%04d %s", line, goatrope.Render(g, start, stop))

		line++
	}
}
