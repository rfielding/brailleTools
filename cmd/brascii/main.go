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
	0b00001000, 0b00000001, 0b00000011, 0b00001001, 0b00011001, 0b00010001, 0b00001011, 0b00011011, 0b00010011, 0b01001010, 0b01011010, 0b01000101, 0b01000111, 0b01001101, 0b01011101, 0b01010101,
	0b00001111, 0b00011111, 0b00010111, 0b00001110, 0b00011110, 0b00100101, 0b00100111, 0b01111010, 0b00101101, 0b01111101, 0b01110101, 0b00101010, 0b00110011, 0b00111011, 0b00011000, 0b00111000,
}

// When looking at 0x20 through 0x5F as 6-dot, mask off dot 7 first,
// as there are only upper-case letters in braille ascii
var braillePerm = []int{
	0b01000000, 0b01101110, 0b01010000, 0b01111100, 0b01101011, 0b01101001, 0b01101111, 0b01000100, 0b01110111, 0b01111110, 0b01100001, 0b01101100, 0b01100000, 0b01100100, 0b01101000, 0b01001100,
	0b01110100, 0b01000010, 0b01000110, 0b01010010, 0b01110010, 0b01100010, 0b01010110, 0b01110110, 0b01100110, 0b01010100, 0b01110001, 0b01110000, 0b01100011, 0b01111111, 0b01011100, 0b01111001,
	// This range is the standard ascii braille, except dot7 is flipped for letters
	//       0           1           2           3           4           5           6            7           8           9           A           B           C           D           E           F
	0b00000000, 0b00101110, 0b00010000, 0b00111100, 0b00101011, 0b00101001, 0b00101111, 0b00000100, 0b00110111, 0b00111110, 0b00100001, 0b00101100, 0b00100000, 0b00100100, 0b00101000, 0b00001100,
	0b00110100, 0b00000010, 0b00000110, 0b00010010, 0b00110010, 0b00100010, 0b00010110, 0b00110110, 0b00100110, 0b00010100, 0b00110001, 0b00110000, 0b00100011, 0b00111111, 0b00011100, 0b00111001,
	0b00001000, 0b01000001, 0b01000011, 0b01001001, 0b01011001, 0b01010001, 0b01001011, 0b01011011, 0b01010011, 0b01001010, 0b01011010, 0b01000101, 0b01000111, 0b01001101, 0b01011101, 0b01010101,
	0b01001111, 0b01011111, 0b01010111, 0b01001110, 0b01011110, 0b01100101, 0b01100111, 0b01111010, 0b01101101, 0b01111101, 0b01110101, 0b00101010, 0b00110011, 0b00111011, 0b00011000, 0b00111000,
	// end of standard range
	0b01001000, 0b00000001, 0b00000011, 0b00001001, 0b00011001, 0b00010001, 0b00001011, 0b00011011, 0b00010011, 0b00001010, 0b00011010, 0b00000101, 0b00000111, 0b00001101, 0b00011101, 0b00010101,
	0b00001111, 0b00011111, 0b00010111, 0b00001110, 0b00011110, 0b00100101, 0b00100111, 0b00111010, 0b00101101, 0b00111101, 0b00110101, 0b01101010, 0b01110011, 0b01111011, 0b01011000, 0b01111000,
	/// high bits
	0b11000000, 0b11101110, 0b11010000, 0b11111100, 0b11101011, 0b11101001, 0b11101111, 0b11000100, 0b11110111, 0b11111110, 0b11100001, 0b11101100, 0b11100000, 0b11100100, 0b11101000, 0b11001100,
	0b11110100, 0b11000010, 0b11000110, 0b11010010, 0b11110010, 0b11100010, 0b11010110, 0b11110110, 0b11100110, 0b11010100, 0b11110001, 0b11110000, 0b11100011, 0b11111111, 0b11011100, 0b11111001,
	// This range is the standard ascii braille, except dot7 is flipped for letters
	//       0           1           2           3           4           5           6            7           8           9           A           B           C           D           E           F
	0b10000000, 0b10101110, 0b10010000, 0b10111100, 0b10101011, 0b10101001, 0b10101111, 0b10000100, 0b10110111, 0b10111110, 0b10100001, 0b10101100, 0b10100000, 0b10100100, 0b10101000, 0b10001100,
	0b10110100, 0b10000010, 0b10000110, 0b10010010, 0b10110010, 0b10100010, 0b10010110, 0b10110110, 0b10100110, 0b10010100, 0b10110001, 0b10110000, 0b10100011, 0b10111111, 0b10011100, 0b10111001,
	0b10001000, 0b11000001, 0b11000011, 0b11001001, 0b11011001, 0b11010001, 0b11001011, 0b11011011, 0b11010011, 0b11001010, 0b11011010, 0b11000101, 0b11000111, 0b11001101, 0b11011101, 0b11010101,
	0b11001111, 0b11011111, 0b11010111, 0b11001110, 0b11011110, 0b11100101, 0b11100111, 0b11111010, 0b11101101, 0b11111101, 0b11110101, 0b10101010, 0b10110011, 0b10111011, 0b10011000, 0b10111000,
	// end of standard range
	0b11001000, 0b10000001, 0b10000011, 0b10001001, 0b10011001, 0b10010001, 0b10001011, 0b10011011, 0b10010011, 0b10001010, 0b10011010, 0b10000101, 0b10000111, 0b10001101, 0b10011101, 0b10010101,
	0b10001111, 0b10011111, 0b10010111, 0b10001110, 0b10011110, 0b10100101, 0b10100111, 0b10111010, 0b10101101, 0b10111101, 0b10110101, 0b11101010, 0b11110011, 0b11111011, 0b11011000, 0b11111000,
	/// high bits
}

func brailleInit() {
	// Copy in the standard braille ascii patern
	for i := 0; i < 64; i++ {
		asciiPerm[0x20+i] = brailleAsciiPattern[i]
	}
	// Flip the case of the alphabet
	for i := 0x41; i <= 0x5A; i++ {
		asciiPerm[i] = asciiPerm[i] ^ 0x40
	}
	// Copy lower half of standard to cover control codes
	for i := 0; i < 32; i++ {
		asciiPerm[i] = (asciiPerm[i+0x20]) ^ 0x40
	}
	// Copy upper half of standard to cover upper case
	for i := 0; i < 32; i++ {
		asciiPerm[0x60+i] = asciiPerm[0x40+i] ^ 0x40
	}
	// Duplicated it all in high bits
	for i := 0; i < 128; i++ {
		asciiPerm[0x80+i] = asciiPerm[i] ^ 0x80
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
				if b[0] == '\n' {
					fmt.Printf("\n")
					continue
				}
				if b[0] == '\r' && *keepCR == false {
					continue
				}
			}
			fmt.Printf("%s", string(braillePerm[v]+0x2800))
		}
	}
}
