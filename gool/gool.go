package gool

import "unsafe"

type Gool struct {
	Values []bool
}

func (g Gool) Encode() byte {
	var encoded byte

	for i := 0; i < 8; i++ {
		bit := g.bool2int(g.Values[i])
		encoded = (encoded << 1) | bit
	}

	return encoded
}

func (g Gool) bool2int(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func (g Gool) GetBinaryDigits() string {
	values := g.Encode()
	var mask byte = 1 << byte(unsafe.Sizeof(values)*8-1)
	var bitField string

	for i := 0; byte(i) < byte(unsafe.Sizeof(values)*8); i++ {
		if (values & mask) == 0 {
			bitField += "0"
		} else {
			bitField += "1"
		}

		mask >>= 1
	}

	return bitField
}
