package converter

import (
	"unsafe"
)

func GetBinaryDigits(value uint) string {
	var mask uint = 1 << uint(unsafe.Sizeof(value) * 8 - 1)
	var bitField string

	for i := 0; i < int(unsafe.Sizeof(value) * 8); i++ {
		if (value & mask) == 0 {
			bitField += "0"
		} else {
			bitField += "1"
		}

		mask >>= 1
	}

	return bitField
}