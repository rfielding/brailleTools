package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"io"
	"os"
	"time"
)

const (
	d1       = 1
	d2       = 2
	d3       = 4
	d4       = 8
	d5       = 16
	d6       = 32
	d7       = 64
	d8       = 128
	dNewLine = 256
)

type Cell int

type BrailleTerminal struct {
	Err      io.Writer
	CursorX  int
	CursorY  int
	Timer    int
	Output   [][]Cell
	Table    []Cell
	TermX    int
	TermY    int
	TermLine int
	Quit     bool
	EditMode bool
}

// This is for trivial 1-cell runes only
func (b *BrailleTerminal) setupBrailleTable() {
	b.Table[33] = d2|d3|d4|d6
	b.Table[34] = d5
	b.Table[35] = d3|d4|d5|d6
	b.Table[36] = d1|d2|d4|d6
	b.Table[37] = d1|d4|d6
	b.Table[38] = d1|d2|d3|d4|d6
	b.Table[39] = d3
	b.Table[40] = d1|d2|d3|d5|d6
	b.Table[41] = d2|d3|d4|d5|d6
	b.Table[42] = d1|d6
	b.Table[43] = d3|d4|d6
	b.Table[44] = d6
	b.Table[45] = d3|d6
	b.Table[46] = d4|d6
	b.Table[47] = d3|d4
	b.Table[58] = d1|d5|d6
	b.Table[59] = d5|d6
	b.Table[60] = d1|d2|d6
	b.Table[61] = d1|d2|d3|d4|d5|d6
	b.Table[62] = d3|d4|d5
	b.Table[63] = d1|d4|d5|d6
	b.Table[64] = d4|d7
	b.Table[91] = d2|d4|d6|d7
	b.Table[92] = d1|d2|d5|d6|d7
	b.Table[93] = d1|d2|d4|d5|d6|d7
	b.Table[94] = d4|d5|d7
	b.Table[95] = d4|d5|d6
	b.Table[96] = d4
	b.Table[123] = d2|d4|d6
	b.Table[124] = d1|d2|d5|d6
	b.Table[125] = d1|d2|d4|d5|d6
	b.Table[126] = d4|d5
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
	b.Table = make([]Cell, 256*256)
	b.setupBrailleTable()
	b.Output = make([][]Cell, b.TermY)
	for y := 0; y < b.TermY; y++ {
		b.Output[y] = make([]Cell, b.TermLine)
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
		if withCursor && b.CursorX == i {
			if b.EditMode == false {
				if b.Timer%2 == 0 {
					cursor = d7
				} else {
					cursor = d8
				}
			} else {
				if (b.Timer % 2) == 0 {
					cursor = d7 | d8
				}
			}
		}
		v := b.Output[b.CursorY][i] | Cell(cursor)
		if b.Output[b.CursorY][i] == dNewLine {
			if (b.Timer % 2) == 0 {
				v = d8
			} else {
				v = 0
			}
		}
		fmt.Printf("%c", rune((0x2800 + int(v))))
	}
}

func (b *BrailleTerminal) kb() <-chan keyboard.KeyEvent {
	evs, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	return evs
}

func (b *BrailleTerminal) Put(c rune) {
	if c == '\n' {
		if b.CursorY+1 < b.TermY {
			b.Output[b.CursorY][b.CursorX] = dNewLine
		}
		for i := b.CursorX + 1; i < b.TermX; i++ {
			b.Output[b.CursorY][i] = 0
		}
		b.CursorNewline()
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

func (b *BrailleTerminal) Handle(c keyboard.KeyEvent) {
	b.Errf("key: %d rune: %c\n", c.Key, c.Rune)
	switch c.Key {
	case keyboard.KeyArrowUp:
		b.CursorUp()
	case keyboard.KeyArrowDown:
		b.CursorDown()
	case keyboard.KeyArrowLeft:
		if b.CursorX == 0 && b.CursorY > 0 {
			b.CursorY--
			for i := 0; i < b.TermX; i++ {
				if b.Output[b.CursorY][i] == dNewLine {
					b.CursorX = i
					break
				}
			}
		} else {
			b.CursorLeft()
		}
	case keyboard.KeyArrowRight:
		if b.Output[b.CursorY][b.CursorX] == dNewLine {
			b.CursorNewline()
		} else {
			b.CursorRight()
		}
	case keyboard.KeyEsc:
		b.EditMode = !b.EditMode
	default:
		if b.EditMode {
			switch c.Key {
			case keyboard.KeyEnter:
				b.Puts("\n")
			default:
				b.Puts(fmt.Sprintf("%c", c.Rune))
			}
		} else {
			switch c.Rune {
			case 'q':
				b.Quit = true
			}
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
		case ev := <-kb:
			b.Handle(ev)
		case <-time.After(250 * time.Millisecond):
		}
		b.RenderLine(true)
		b.Timer++
	}
}

func main() {
	b := NewTerminal(19,19, 20, "BrailleTerminal.err.log")
	b.EditMode = true
	//b.Puts("In the beginning\nGod created\nthe heaven\nand the earth\n")
	b.Render()
}
