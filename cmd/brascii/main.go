package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 1) start with braille ascii for 0x20 to 0x5F
// 2) flip dot 7 for all letters in 0x40 to 0x5F
// 3) copy 0x20..0x3F to 0x00, and set dot 7 on 0x00..0x1F
// 4) copy 0x40..0x5F to 0x60, and set dot 7 on 0x60..0x7F
// 5) flip 7 dot for all letters in 0x60..0x7F
// 6) copy all of 0x00..0x7F to 0x80, and set dot 8 on 0x80..0xFF
//
// Note that in this table, you can read the braille directly:
// using dot number:
//            0b01011001
//              87654321  .. the braille dots... imagine re-laid out:
//
//               14
//               25
//               36
//               78
//
// by looking at it sideways.  after some practice, you can read dots from binary.
//
var brailleAsciiPattern = []int{
	0b00000000, 0b00101110, 0b00010000, 0b00111100, 0b00101011, 0b00101001, 0b00101111, 0b00000100, 0b00110111, 0b00111110, 0b00100001, 0b00101100, 0b00100000, 0b00100100, 0b00101000, 0b00001100,
	0b00110100, 0b00000010, 0b00000110, 0b00010010, 0b00110010, 0b00100010, 0b00010110, 0b00110110, 0b00100110, 0b00010100, 0b00110001, 0b00110000, 0b00100011, 0b00111111, 0b00011100, 0b00111001,
	0b00001000, 0b00000001, 0b00000011, 0b00001001, 0b00011001, 0b00010001, 0b00001011, 0b00011011, 0b00010011, 0b00001010, 0b00011010, 0b00000101, 0b00000111, 0b00001101, 0b00011101, 0b00010101,
	0b00001111, 0b00011111, 0b00010111, 0b00001110, 0b00011110, 0b00100101, 0b00100111, 0b00111010, 0b00101101, 0b00111101, 0b00110101, 0b00101010, 0b00110011, 0b00111011, 0b00011000, 0b00111000,
}

// When looking at 0x20 through 0x5F as 6-dot, mask off dot 7 first,
// as there are only upper-case letters in braille ascii
var braillePerm = make([]int, 256)

func brailleInit() {
	// Copy in the standard braille ascii patern
	for i := 0; i < 64; i++ {
		braillePerm[0x20+i] = brailleAsciiPattern[i]
	}
	// Flip the case of the alphabet
	for i := 0x41; i <= 0x5A; i++ {
		braillePerm[i] = braillePerm[i] ^ 0x40
	}
	// Copy lower half of standard to cover control codes
	for i := 0; i < 32; i++ {
		braillePerm[i] = (braillePerm[i+0x20]) ^ 0x40
	}
	// Copy upper half of standard to cover upper case
	for i := 0; i < 32; i++ {
		braillePerm[0x60+i] = braillePerm[0x40+i] ^ 0x40
	}
	// Duplicated it all in high bits
	for i := 0; i < 128; i++ {
		braillePerm[0x80+i] = braillePerm[i] ^ 0x80
	}
	// Reverse mapping
	for i := 0; i < 256; i++ {
		asciiPerm[braillePerm[i]] = i
		present[i]++
	}
	// Panic if codes are missing or duplicated
	for i := 0; i < 256; i++ {
		if present[i] != 1 {
			panic(fmt.Sprintf("inconsistency at %d", i))
		}
	}

}

var asciiPerm = make([]int, 256)
var present = make([]int, 256)

// Byte by byte translation to braille
func main() {
	sixDot := flag.Bool("sixdot", false, "decode as 6-dot ascii braille")
	decode := flag.Bool("decode", false, "decode braile to ascii binary")
	isBinary := flag.Bool("binary", false, "literal binary translation, even of CR/LF")
	keepCR := flag.Bool("keep-cr", false, "keep literal CR so that binary back translate is still unambiguous")
	help := flag.Bool("help", false, "show help")
	flag.Parse()
	if *help == true {
		flag.Usage()
		os.Exit(0)
	}
	brailleInit()
	// Setup the reverse table to convert braille to ascii

	if *decode == true {
		r := bufio.NewReader(os.Stdin)
		for {
			c, l, err := r.ReadRune()
			if l == 0 {
				break
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			if 0x2800 <= c && c < 0x28FF {
				fmt.Printf("%s", string(asciiPerm[c-0x2800]))
			} else {
				fmt.Printf("%s", string(c))
			}
		}
	} else {
		b := make([]byte, 1)
		for {
			count, _ := os.Stdin.Read(b)
			// Assume that we only get 0 bytes when no data
			if count == 0 {
				break
			}

			if count == 0 {
				continue
			}
			v := int(b[0])
			// Encode bytes to braille
			if *isBinary == false {
				if v == '\n' {
					fmt.Printf("\n")
					continue
				}
				if v == '\r' && *keepCR == false {
					continue
				}
			}
			var c int
			if *sixDot {
				c = (braillePerm[v] & 0x3F)+0x2800
			} else {
				c = braillePerm[v]+0x2800
			}
			fmt.Printf("%s", string(c))
		}
	}
}
