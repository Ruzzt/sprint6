package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input []byte) (string, error) {
	str := strings.TrimSpace(string(input))
	if str == "" {
		return "", errors.New("пустая строка")
	}

	isMorse := true
	for _, c := range str {
		if c != '.' && c != '-' && c != ' ' && c != '/' {
			isMorse = false
			break
		}
	}

	if isMorse {
		// Морзе в текст
		return morse.ToText(str), nil
	} else {
		// Текст в Морзе
		return morse.ToMorse(str), nil
	}
}
