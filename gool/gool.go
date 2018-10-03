package gool

import (
	"unsafe"
)

type Gool struct {
	Values      [8]bool
	encodedByte byte
}

func (g *Gool) Encode() byte {
	var encoded byte

	for i := 0; i < 8; i++ {
		bit := g.bool2int(g.Values[i])
		encoded = (encoded << 1) | bit
	}

	g.encodedByte = encoded

	return encoded
}

func (g Gool) Decode() [8]bool {
	var flags [8]bool
	encodedByte := g.encodedByte

	for i := range flags {
		flag := g.byte2bool(encodedByte & 1)
		flags[len(flags)-1-i] = flag

		encodedByte >>= 1
	}

	return flags
}

func (g Gool) byte2bool(b byte) bool {
	return b == 1
}

func (g Gool) bool2int(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func GetBinaryDigits(b byte) string {
	var mask byte = 1 << byte(unsafe.Sizeof(b)*8-1)
	var bitField string

	for i := 0; byte(i) < byte(unsafe.Sizeof(b)*8); i++ {
		if (b & mask) == 0 {
			bitField += "0"
		} else {
			bitField += "1"
		}

		mask >>= 1
	}

	return bitField
}
