package gool

import "testing"

func TestEncode(t *testing.T) {
	flags := [8]bool{true, false, false, true, true, true, false, true}

	g := Gool{Values: flags}

	resultByte := g.Encode()

	if resultByte != 157 { // 157 = 10011101
		t.Error("Expected 157 (10011101), got ", resultByte)
	}
}

func TestDecode(t *testing.T) {
	flags := [8]bool{true, false, false, true, true, true, false, true}

	g := Gool{Values: flags}
	g.Encode()

	resultSlice := g.Decode()

	if len(flags) != len(resultSlice) {
		t.Error("Not the same size! Expected ", flags, ", got ", resultSlice)
	}

	for i := range flags {
		if flags[i] != resultSlice[i] {
			t.Error("Expected ", flags, ", got ", resultSlice)
		}
	}
}

func TestGetBinaryDigits(t *testing.T) {
	var testByte byte = 3
	resultString := GetBinaryDigits(testByte)
	expectedValue := "00000011"

	if resultString != expectedValue {
		t.Error("Expected ", expectedValue, ", got ", resultString)
	}
}
