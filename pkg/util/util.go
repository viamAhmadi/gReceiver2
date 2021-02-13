package util

import (
	"strconv"
	"strings"
)

// RemoveAdditionalCharacters
func RemoveAdditionalCharacters(b []byte) string {
	return strings.TrimSpace(strings.ReplaceAll(string(b), "A", " "))
}

// ConvertDesToBytes converts destination to bytes array
func ConvertDesToBytes(d string) []byte {
	b := []byte(d)
	if len(b) < 27 {
		l := 27 - len(b)
		for i := 1; i <= l; i++ {
			b = append(b, 65)
		}
	}
	return b
}

// ConvertIntToBytes
func ConvertIntToBytes(id int) []byte {
	b := []byte(strconv.Itoa(id))
	if len(b) < 5 {
		l := 5 - len(b)
		for i := 1; i <= l; i++ {
			b = append(b, 65)
		}
	}
	return b
}

