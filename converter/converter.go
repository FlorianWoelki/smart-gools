package converter

import (
	"unsafe"
)

func GetBinaryDigits(value byte) string {
	var mask byte = 1 << byte(unsafe.Sizeof(value)*8-1)
	var bitField string

	for i := 0; byte(i) < byte(unsafe.Sizeof(value)*8); i++ {
		if (value & mask) == 0 {
			bitField += "0"
		} else {
			bitField += "1"
		}

		mask >>= 1
	}

	return bitField
}
