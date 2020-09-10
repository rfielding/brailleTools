package goatrope

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

func ToJson(v interface{}) string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}

/*
 A GoatRope is my variant of a PieceTable for bytes

 This is the hard, and most general component
 of making a text editor.  It is kind of like
 extending the normal Unix file interface to
 handle inserts and deletions efficiently.

 It is like a standard *io.File that
 allows for efficient inserts and deletes.

 It is also possible to use the Write method
 to delete from it, by writing in escape
 sequences that perform deletes and backspaces
 on the file rather than writing them IN
 to the file.
*/

// File puts down an interface for what passes as a File for us
type File interface {
	io.Reader
	io.Seeker
	io.Closer
	io.Writer
	Stat() (os.FileInfo, error)
}

// Ensure that indeed a *os.File does implement File
var _ File = &os.File{}

// GoatRope extends File with extra behaviors
type GoatRope struct {
	PieceTable *PieceTable
	FileInfo   *GoatRopeFileInfo
	Original   File
	Mods       File
}

var _ File = &GoatRope{}

type GoatRopeFileInfo struct {
	PieceTable *PieceTable
	name       string
	size       int64
	mode       os.FileMode
	modtime    time.Time
	isdir      bool
	sys        interface{}
}

func (fi *GoatRopeFileInfo) Name() string {
	return fi.name
}

func (fi *GoatRopeFileInfo) Size() int64 {
	return fi.PieceTable.Size()
}

func (fi *GoatRopeFileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi *GoatRopeFileInfo) ModTime() time.Time {
	return fi.modtime
}

func (fi *GoatRopeFileInfo) IsDir() bool {
	return fi.isdir
}

func (fi *GoatRopeFileInfo) Sys() interface{} {
	return fi.sys
}

// Allocate a new and unloaded GoatRope
func NewGoatRope() *GoatRope {
	g := &GoatRope{}
	g.PieceTable = &PieceTable{}
	g.FileInfo = &GoatRopeFileInfo{PieceTable: g.PieceTable}
	g.Mods = &MemoryFile{}
	return g
}

func (g *GoatRope) LoadByName(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	g.Original = f
	s, err := os.Stat(name)
	if err != nil {
		return err
	}
	g.Mods = &MemoryFile{}
	g.PieceTable.Load(s.Size())
	return nil
}

func (g *GoatRope) Seek(to int64, whence int) (int64, error) {
	g.PieceTable.Index = to
	return to, nil
}

func (g *GoatRope) Write(b []byte) (int, error) {
	// This is an append-only device, so
	// seeking done for reads are ignored here.
	written, err := g.Mods.Write(b)
	if err != nil {
		return written, err
	}
	g.Mods.Seek(g.PieceTable.Size(), io.SeekStart)
	g.PieceTable.Insert(int64(written))
	return written, nil
}

func (g *GoatRope) Close() error {
	err := g.Original.Close()
	err2 := g.Mods.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func (g *GoatRope) theFile(isOriginal bool) File {
	if isOriginal {
		return g.Original
	}
	return g.Mods
}

func (g *GoatRope) Read(data []byte) (int, error) {
	read := 0
	// try to completely fill the data array
	lo := int64(0)
	// cutlo can move up as we read data
	for i := 0; i < len(g.PieceTable.Pieces); i++ {
		cutlo := g.PieceTable.Index
		cuthi := cutlo + int64(len(data))
		hi := lo + g.PieceTable.Pieces[i].Size
		sz := hi - lo
		if hi <= cutlo {
			// wait for matching part
		} else if cuthi <= lo {
			break
		} else {
			// There is some kind of reading work
			// to do
			f := g.theFile(
				g.PieceTable.Pieces[i].Original,
			)
			f.Seek(
				g.PieceTable.Pieces[i].Start+
					(cutlo-lo),
				io.SeekStart,
			)
			max := (hi - cutlo)
			if cuthi < hi {
				max = (cuthi - cutlo)
			}
			if max == 0 {
				return read, io.EOF
			}
			rd, err := io.ReadFull(
				f,
				data[read:read+int(max)],
			)
			read += rd
			g.PieceTable.Index += int64(rd)
			if err != nil {
				return read, err
			}
		}
		lo += sz
	}
	if read == 0 {
		return 0, io.EOF
	}
	return read, nil
}

func (g *GoatRope) Stat() (os.FileInfo, error) {
	return g.FileInfo, nil
}

func (g *GoatRope) Cut(n int64) {
	g.PieceTable.Cut(n)
}
