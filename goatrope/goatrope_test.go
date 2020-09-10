package goatrope_test

import (
	"github.com/rfielding/teeny/goatrope"
	"testing"
	"os"
	"io"
)

func TestGoatRope(t *testing.T) {
	g := goatrope.NewGoatRope()
	g.LoadByName("goatrope.go")
	g.Seek(17, io.SeekStart)
	g.Write([]byte("// hello\n"))
	g.Seek(0, io.SeekStart)
	io.Copy(os.Stdout, g)
	defer g.Close()
}
