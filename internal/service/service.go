package service

import (
	"errors"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
	"unicode"
)

func ConvertString(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("convert ERROR: the string must not be empty")
	}
	isMorse := true
	slice := strings.Split(s, " ")
	for _, word := range slice {
		for _, char := range word {
			if char == '-' || char == '.' {
				continue
			} else {
				isMorse = false
				if unicode.IsLetter(char) {
					char = unicode.ToUpper(char)
				}
				_, ok := morse.DefaultMorse[char]
				if !ok {
					return "", errors.New(`convert ERROR: the string may contain only Cyrillic letters, digits, or the following symbols: '.', ',', ':', '?', '\', '-', '/', '(', ')', '"'`)
				}
			}
		}
	}
	if isMorse {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
