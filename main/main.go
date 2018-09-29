package main

import (
	"fmt"
	"smart-gools/converter"
)

func main() {
	var value uint = 240

	fmt.Println(converter.GetBinaryDigits(value))
	fmt.Println(converter.GetBinaryDigits(value >> 4))
}