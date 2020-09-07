package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"io"
	"os"
	"time"
	//	"math/rand"
)

const (
	d1 = 1
	d2 = 2
	d3 = 4
	d4 = 8
	d5 = 16
	d6 = 32
	d7 = 64
	d8 = 128
)

type BrailleTerminal struct {
	Err      io.Writer
	CursorX  int
	CursorY  int
	Timer    int
	Output   [][]int
	Table    []int
	TermX    int
	TermY    int
	TermLine int
	Quit     bool
	EditMode bool
}

func (b *BrailleTerminal) setupBrailleTable() {
	// a through j
	b.Table[0+97] = d1
	b.Table[1+97] = d1 | d2
	b.Table[2+97] = d1 | d4
	b.Table[3+97] = d1 | d4 | d5
	b.Table[4+97] = d1 | d5
	b.Table[5+97] = d1 | d2 | d4
	b.Table[6+97] = d1 | d2 | d4 | d5
	b.Table[7+97] = d1 | d2 | d5
	b.Table[8+97] = d2 | d4
	b.Table[9+97] = d2 | d4 | d5
	for i := 0; i <= 9; i++ {
		// k through t
		b.Table[i+0x6b] =
			b.Table[i+97] + d3
	}
	// u through z (irregular
	b.Table[0+0x75] = b.Table[97] | d3 | d6
	b.Table[1+0x75] = b.Table[98] | d3 | d6
	b.Table[2+0x75] = d2 | d4 | d5 | d6
	b.Table[3+0x75] = d1 | d4 | d3 | d6 | d3 | d6
	b.Table[4+0x75] = d1 | d4 | d3 | d6 | d5 | d3 | d6
	b.Table[5+0x75] = d1 | d3 | d5 | d6 | d3 | d6
	// Upper case
	for i := 0; i < 26; i++ {
		b.Table[65+i] = b.Table[97+i] + d7
	}
	// 1 to 10
	b.Table[0+0x30] = d3 | d5 | d6
	b.Table[1+0x30] = d2
	b.Table[2+0x30] = d2 | d3
	b.Table[3+0x30] = d2 | d5
	b.Table[4+0x30] = d2 | d5 | d6
	b.Table[5+0x30] = d2 | d6
	b.Table[6+0x30] = d2 | d3 | d5
	b.Table[7+0x30] = d2 | d3 | d5 | d6
	b.Table[8+0x30] = d2 | d3 | d6
	b.Table[9+0x30] = d3 | d5
}

func NewTerminal(termX, termLine, termY int, errLog string) *BrailleTerminal {
	b := &BrailleTerminal{
		TermX:    termX,
		TermY:    termY,
		TermLine: termLine,
	}
	berr, err := os.Create(errLog)
	if err != nil {
		panic(err)
	}
	b.Err = berr
	b.Table = make([]int, 256)
	b.setupBrailleTable()
	b.Output = make([][]int, b.TermY)
	for y := 0; y < b.TermY; y++ {
		b.Output[y] = make([]int, b.TermLine)
	}
	return b
}

func (b *BrailleTerminal) Errf(msg string, args ...interface{}) {
	b.Err.Write([]byte(fmt.Sprintf(msg, args...)))
}

func (b *BrailleTerminal) CursorUp() {
	if b.CursorY > 0 {
		b.CursorY--
	}
}

func (b *BrailleTerminal) CursorDown() {
	if b.CursorY+1 < b.TermY {
		b.CursorY++
	}
}

func (b *BrailleTerminal) CursorLeft() {
	if b.CursorX > 0 {
		b.CursorX--
	}
}

func (b *BrailleTerminal) CursorRight() {
	if b.CursorX+1 < b.TermX {
		b.CursorX++
	}
}

func (b *BrailleTerminal) CursorNewline() {
	if b.CursorY+1 < b.TermY {
		b.CursorY++
		b.CursorX = 0
	}
}

// Render braille on a single line of terminalLength
func (b *BrailleTerminal) RenderLine(withCursor bool) {
	fmt.Printf("\r")
	for i := 0; i < b.TermLine; i++ {
		cursor := 0
		if withCursor && b.CursorX == i && (b.Timer%2) == 0 {
			cursor = d7 | d8
		}
		fmt.Printf("%c", rune((0x2800+b.Output[b.CursorY][i])|cursor))
	}
}

func (b *BrailleTerminal) kb() chan rune {
	ret := make(chan rune, 0)
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	go func() {
		for {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			ret <- char
		}
		_ = keyboard.Close()
	}()
	return ret
}

func (b *BrailleTerminal) Put(c rune) {
	if c == '\n' {
		if b.CursorY+1 < b.TermY {
			b.CursorX = 0
			b.CursorY++
		}
	} else {
		if b.CursorX+1 < b.TermX {
			b.Output[b.CursorY][b.CursorX] = b.Table[c]
			b.CursorX++
		}
	}
}

func (b *BrailleTerminal) Puts(s string) {
	for i := 0; i < len(s); i++ {
		b.Put(rune(s[i]))
	}
}

func (b *BrailleTerminal) ClearLine() {
	b.CursorX = 0
	for i := 0; i < b.TermLine; i++ {
		b.Output[b.CursorY][i] = 0
	}
}

func (b *BrailleTerminal) Handle(c rune) {
	//b.Errf("%c",c)
	fmt.Printf("%c",c)
	if b.EditMode {
		switch c {
		case '`':
			b.EditMode = false
		default:
			b.Puts(fmt.Sprintf("%c", c))
		}
	} else {
		switch c {
		case 'i':
			b.EditMode = true
		case 'q':
			b.Quit = true
		case 'h':
			b.CursorLeft()
		case 'l':
			b.CursorRight()
		case 'k':
			b.CursorUp()
		case 'j':
			b.CursorDown()
		case ' ':
			b.CursorRight()
		}
	}
}

func (b *BrailleTerminal) Render() {
	kb := b.kb()
	// Set terminal state and stall until we are done
	fmt.Printf("\033[2J\033[?25l")
	defer func() {
		fmt.Printf("\033[?25h\n")
	}()

	b.Timer = 0
	// Handle keys if available, or tick
	for b.Quit == false {
		select {
		case c := <-kb:
			b.Handle(c)
		case <-time.After(250 * time.Millisecond):
		}
		b.RenderLine(true)
		b.Timer++
	}
}

func main() {
	b := NewTerminal(19, 19, 20, "BrailleTerminal.err.log")
	b.Puts("test")
	b.Puts("\n")
	b.Puts("Decipher")
	b.Render()
}
