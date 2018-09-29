package main

import (
	"fmt"
	"smart-gools/converter"
)

func main() {
	flags := [8]bool{true, false, false, true, true, true, false, true}

	var encoded byte

	for i := 0; i < 8; i++ {
		bit := bool2int(flags[i])
		encoded = (encoded << 1) | bit
	}

	fmt.Println(converter.GetBinaryDigits(encoded))
}

func bool2int(b bool) byte {
	if b {
		return 1
	}
	return 0
}
