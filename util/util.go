package util

import (
	"fmt"
	"strconv"
)

//
// Convert string to binary string
//
func StrToBinstr(str string) string {
	var binstr string
	for i := 0; i < len(str); i++{
		binstr = fmt.Sprintf("%s%.8b", binstr, str[i]) // Append each byte as 8 bit binary string
	}
	return binstr
}

//
// Convert binary string to string
//
func BinstrToStr(binstr string) string {
	var str string
	for i := 0; i < len(binstr); i+=8{
		b, _ := strconv.ParseInt(binstr[i:i+8], 2, 8) // Get groups of 8 characters (8 bits) as single byte
		str = fmt.Sprintf("%s%c", str, b) // Append single byte as char
	}
	return str
}