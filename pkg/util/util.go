package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func CalculateTimeout(baseTimeout int, numberOfPackets int) time.Duration {
	if numberOfPackets < 10000 {
		return time.Duration(baseTimeout) * time.Second
	}
	return time.Duration(baseTimeout+((numberOfPackets/10000)/2)) * time.Second
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}