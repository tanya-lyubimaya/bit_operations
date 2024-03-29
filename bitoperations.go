package bitoperations

import (
	"errors"
	"strings"
)

const (
	UPPER = 1 << iota
	LOWER
	CAP
	REV
)

func Upper(str string) (string, error) {
	return processString(str, UPPER)
}

func Lower(str string) (string, error) {
	return processString(str, LOWER)
}

func Capitalize(str string) (string, error) {
	return processString(str, CAP)
}

func Reverse(str string) (string, error) {
	return processString(str, REV)
}

func UpperAndReverse(str string) (string, error) {
	return processString(str, UPPER|REV)
}

func LowerAndReverse(str string) (string, error) {
	return processString(str, LOWER|REV)
}

func CapitalizeAndReverse(str string) (string, error) {
	return processString(str, CAP|REV)
}

func SameSign(a, b int) bool {
	// sameSign checks whether numbers a and b have the same sign
	return a^b >= 0
}

func SetNthBit(num, pos int) int {
	// setNthBit sets bit on position pos in num
	return num | (1 << (pos - 1))
}

func UnsetNthBit(num, pos int) int {
	// unsetNthBit unsets bit on position pos in num
	return num &^ (1 << (pos - 1))
}

func GetNthBit(num, pos int) int {
	// getNthBit returns bit on position pos in num
	return num & (1 << (pos - 1))
}

func processString(str string, conf byte) (string, error) {
	// processString changes string str according to configuration conf:
	// UPPER - make str uppercase
	// LOWER - make str lowercase
	// CAP - capitalize first letter of each word in str
	// REV - reverse the letters order in str
	// conf can set multiple parameters as in conf = REV|CAP
	if len(str) == 0 {
		return str, errors.New("Empty string")
	}
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	title := func(input string) string {
		words := strings.Fields(input)
		smallwords := " a an on the to "

		for index, word := range words {
			if strings.Contains(smallwords, " "+word+" ") {
				words[index] = word
			} else {
				words[index] = strings.Title(word)
			}
		}
		return strings.Join(words, " ")
	}

	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = title(strings.ToLower(str))
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str, nil
}
