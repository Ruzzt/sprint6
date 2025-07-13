package service

import (
	"errors"
	"strings"
)

var textToMorse = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..",
	'E': ".", 'F': "..-.", 'G': "--.", 'H': "....",
	'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
	'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",
	'0': "-----", '1': ".----", '2': "..---", '3': "...--",
	'4': "....-", '5': ".....", '6': "-....", '7': "--...",
	'8': "---..", '9': "----.",
	' ': "/",
}

var morseToText map[string]rune

func init() {
	morseToText = make(map[string]rune)
	for k, v := range textToMorse {
		morseToText[v] = k
	}
}

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

		words := strings.Split(str, " / ")
		var result []string
		for _, word := range words {
			letters := strings.Split(word, " ")
			var decodedWord strings.Builder
			for _, l := range letters {
				if l == "" {
					continue
				}
				if ch, ok := morseToText[l]; ok {
					decodedWord.WriteRune(ch)
				} else {
					return "", errors.New("некорректный код Морзе: " + l)
				}
			}
			result = append(result, decodedWord.String())
		}
		return strings.Join(result, " "), nil
	} else {

		var morse []string
		upper := strings.ToUpper(str)
		for _, c := range upper {
			if code, ok := textToMorse[c]; ok {
				morse = append(morse, code)
			} else {
				return "", errors.New("недопустимый символ в тексте: " + string(c))
			}
		}
		return strings.Join(morse, " "), nil
	}
}
