package service

import (
	"errors"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

func ConvertString(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("convert ERROR: the string must not be empty")
	}
	isMorse := !strings.ContainsFunc(s, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})
	if isMorse {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
