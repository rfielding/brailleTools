package goatrope

import (
	"encoding/json"
	"io"
	"os"
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
type GoatRope interface {
	File
}

// MemoryFile always expects
// to be seeked to the end before
// Write
type MemoryFile struct {
}
