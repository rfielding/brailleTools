package goatrope

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
	"unicode/utf8"
	//	"fmt"
)

func ToJson(v interface{}) string {
	s, _ := json.Marshal(v)
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

// Up and down are recognized by carriage returns
// These are convenience methods for seek
type Navigator interface {
	Lines(start int64, count int64) (int64, int64, error)
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

type RuneScanner struct {
	owner *GoatRope
	rdr   *bufio.Reader
	runeLast rune
	runeSzLast int
	unRead bool
}

var _ io.RuneScanner = NewGoatRope().RuneScanner()

func (g *GoatRope) RuneScanner() *RuneScanner {
	return &RuneScanner{
		owner: g,
		rdr: bufio.NewReader(g),
	}
}


func (r *RuneScanner) ReadLine() ([]byte, bool, error) {
	return r.rdr.ReadLine()
}

func (r *RuneScanner) UnreadRune() error {
	r.unRead = true
	return nil
}

func (r *RuneScanner) ReadRune() (rune, int, error) {
	if r.unRead {
		r.unRead = false
		return r.runeLast, r.runeSzLast, nil
	}
	for peekBytes := 4; peekBytes > 0; peekBytes-- {
		b, err := r.rdr.Peek(peekBytes)
		if err == nil {
			rn, sz := utf8.DecodeRune(b)
			if rn == utf8.RuneError {
				return rn, sz, fmt.Errorf("Rune error")
			}
			// success
			r.runeLast = rn
			r.runeSzLast = sz
			return rn, sz, nil
		}
	}
	return -1, 0, io.EOF
}

func Render(g File, start int64, stop int64) []byte {
	buffer := make([]byte, stop-start)
	g.Seek(start, io.SeekStart)
	_, _ = io.ReadFull(g, buffer)
	return buffer
}

func Lines(g File, buffer []byte, start int64, count int64) (int64, int64, error) {
	idx, err := g.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, 0, err
	}
	lo := idx
	hi := lo
	read := int64(0)
	line := int64(0)
	for {
		rd, err := g.Read(buffer)
		if rd == 0 || err == io.EOF {
			return 0, 0, io.EOF
		}
		if err != nil {
			return 0, 0, err
		}
		for i := 0; i < rd; i++ {
			if buffer[i] == '\n' {
				line++
				if line == start {
					lo = idx + read + int64(i) + 1
				}
				if line == start+count {
					hi = idx + read + int64(i) + 1
					g.Seek(hi, io.SeekStart)
					return lo, hi, nil
				}
			}
		}
		read += int64(rd)
	}
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
	if whence == io.SeekStart {
		g.PieceTable.Index = to
	} else if whence == io.SeekCurrent {
		g.PieceTable.Index += to
	} else if whence == io.SeekEnd {
		g.PieceTable.Index = g.PieceTable.Size() - 1 - to
	}
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

func minint64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func (g *GoatRope) Read(data []byte) (int, error) {
	read := 0 // 0 .. len(data)
	lo := int64(0)
	cutlo := g.PieceTable.Index
	cuthi := cutlo + int64(len(data))
	for i := 0; i < len(g.PieceTable.Pieces); i++ {
		hi := lo + g.PieceTable.Pieces[i].Size
		if hi <= cutlo {
			//next
		} else if cuthi <= lo {
			break
		} else {
			// Select the file
			isOriginal := g.PieceTable.Pieces[i].Original
			f := g.theFile(isOriginal)
			for {
				n := cutlo - lo
				allowed := len(data) - read
				available := g.PieceTable.Pieces[i].Size - n
				max := minint64(int64(allowed), available)
				if max == 0 {
					break
				}
				seekTo := g.PieceTable.Pieces[i].Start + n
				// Seek to create the buffer
				f.Seek(
					seekTo,
					io.SeekStart,
				)
				rd, err := io.ReadFull(
					f,
					data[read:read+int(max)],
				)
				cutlo += int64(rd)
				read += rd
				g.PieceTable.Index += int64(rd)
				if err != nil && err != io.EOF {
					return read, err
				}
			}
		}
		lo += g.PieceTable.Pieces[i].Size
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
