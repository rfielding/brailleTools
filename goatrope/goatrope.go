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
	FileInfo *GoatRopeFileInfo
}
var _ File = &GoatRope{}


type GoatRopeFileInfo struct {
	PieceTable *PieceTable
        name    string
        size    int64
        mode    os.FileMode
        modtime time.Time
        isdir   bool
        sys     interface{}
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
	g.PieceTable.Mods = &MemoryFile{}
	return g
}

// Load Original from the filesystem
func (g *GoatRope) LoadByName(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	g.PieceTable.Original = f
	s, err := os.Stat(name)
	if err != nil {
		return err
	}
	g.PieceTable.Mods = &MemoryFile{}
	g.PieceTable.Load(s.Size())
	return nil
}

func (g *GoatRope) Seek(to int64, whence int) (int64, error) {
	g.PieceTable.Index = to
	return to, nil
}

// Seek to where you want to write first!
func (g *GoatRope) Write(b []byte) (int, error) {
	g.PieceTable.Mods.Write(b)
	g.PieceTable.Insert(int64(len(b)))
	return len(b), nil
}

func (g *GoatRope) Close() error {
	return g.PieceTable.Original.Close()
}

func (g *GoatRope) Read(data []byte) (int, error) {
	panic("error")
}

func (g *GoatRope) Stat() (os.FileInfo, error) {
	return g.FileInfo,nil
}

func (g *GoatRope) Cut(n int64) {
	g.PieceTable.Cut(n)
}
