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

func TestGoatRopeByLine(t *testing.T) {

	g := goatrope.NewGoatRope()
	defer g.Close()
	g.LoadByName("goatrope.go")

	g.Seek(17, io.SeekStart)
	g.Write([]byte("// world\n"))
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// hello\n"))
	g.Seek(0, io.SeekStart)

	line := int64(0)
	rdr := g.RuneScanner()
	for {
		lineBytes, _, err := rdr.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("%04d: %s\n", line, string(lineBytes))
		line++
	}
}
