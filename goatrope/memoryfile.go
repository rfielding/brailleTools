package goatrope

import (
	"io"
	"os"
	"time"
)

type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modtime time.Time
	isdir   bool
	sys     interface{}
}

func (fi *FileInfo) Name() string {
	return fi.name
}

func (fi *FileInfo) Size() int64 {
	return fi.size
}

func (fi *FileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi *FileInfo) ModTime() time.Time {
	return fi.modtime
}

func (fi *FileInfo) IsDir() bool {
	return fi.isdir
}

func (fi *FileInfo) Sys() interface{} {
	return fi.sys
}

// MemoryFile is an in-memory file
// - append ONLY for writes
// - seek to read position to read out of it
type MemoryFile struct {
	Data  []byte
	Index int
}

var _ File = &MemoryFile{}

func (m *MemoryFile) Close() error {
	return nil
}

func (m *MemoryFile) Stat() (os.FileInfo, error) {
	return &FileInfo{
		name:    "mods",
		size:    int64(len(m.Data)),
		mode:    0,
		modtime: time.Now(),
		isdir:   false,
		sys:     nil,
	}, nil
}

func (m *MemoryFile) Seek(to int64, whence int) (int64, error) {
	// ASSUME: io.SeekStart
	m.Index = int(to)
	return int64(m.Index), nil
}

// This only APPENDS to the end!
// It honors seeking for READ purposes only
func (m *MemoryFile) Write(data []byte) (int, error) {
	m.Data = append(m.Data, data...)
	return len(data), nil
}

func (m *MemoryFile) Read(data []byte) (int, error) {
	dlen := len(data)
	mlen := len(m.Data)
	thelen := dlen
	if mlen < thelen {
		thelen = mlen - m.Index
	}
	if thelen <= 0 {
		return 0, io.EOF
	}
	for i := 0; i < thelen; i++ {
		data[i] = m.Data[m.Index+i]
	}
	m.Index += thelen
	return thelen, nil
}
