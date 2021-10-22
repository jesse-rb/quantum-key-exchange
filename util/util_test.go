package util

import (
	"fmt"
	"testing"
)

func TestStrBinstrConversion(t *testing.T) {
	s := "hello"
	b := StrToBinstr(s)
	s1 := BinstrToStr(b)
	fmt.Println(s + " -> " + b)
	fmt.Println(b + " -> " + s1)
	
	if (s == b) {
		t.Fatalf("StrToBinstr(string) -> expected: s=%v != b=%v, got: s=%v == b=%v", s, b, s, b)
	}
	if (s != s1) {
		t.Fatalf("BinstrToStr(binstr) -> expected: s=%v == s1=%v, got: s=%v != s1=%v", s, b, s, b)
	}
}