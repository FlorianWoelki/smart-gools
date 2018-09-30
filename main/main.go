package main

import (
	"fmt"
	"smart-gools/gool"
)

func main() {
	flags := []bool{true, false, false, true, true, true, false, true}

	g := gool.Gool{Values: flags}
	fmt.Println(g.GetBinaryDigits())
}
