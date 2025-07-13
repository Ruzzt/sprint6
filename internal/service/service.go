package service

import (
	"errors"
	"strings"
)

var textToMorse = map[rune]string{
	// Русские буквы
	'А': ".-", 'Б': "-...", 'В': ".--", 'Г': "--.", 'Д': "-..",
	'Е': ".", 'Ж': "...-", 'З': "--..", 'И': "..", 'Й': ".---",
	'К': "-.-", 'Л': ".-..", 'М': "--", 'Н': "-.", 'О': "---",
	'П': ".--.", 'Р': ".-.", 'С': "...", 'Т': "-", 'У': "..-",
	'Ф': "..-.", 'Х': "....", 'Ц': "-.-.", 'Ч': "---.", 'Ш': "----",
	'Щ': "--.-", 'Ъ': "--.--", 'Ы': "-.--", 'Ь': "-..-", 'Э': "..-..",
	'Ю': "..--", 'Я': ".-.-",

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
	morseToText = map[string]rune{
		".-":    'А',
		"-...":  'Б',
		".--":   'В',
		"--.":   'Г',
		"-..":   'Д',
		".":     'Е',
		"...-":  'Ж',
		"--..":  'З',
		"..":    'И',
		".---":  'Й',
		"-.-":   'К',
		".-..":  'Л',
		"--":    'М',
		"-.":    'Н',
		"---":   'О',
		".--.":  'П',
		".-.":   'Р',
		"...":   'С',
		"-":     'Т',
		"..-":   'У',
		"..-.":  'Ф',
		"....":  'Х',
		"-.-.":  'Ц',
		"---.":  'Ч',
		"----":  'Ш',
		"--.-":  'Щ',
		"--.--": 'Ъ',
		"-.--":  'Ы',
		"-..-":  'Ь',
		"..-..": 'Э',
		"..--":  'Ю',
		".-.-":  'Я',
		"-----": '0',
		".----": '1',
		"..---": '2',
		"...--": '3',
		"....-": '4',
		".....": '5',
		"-....": '6',
		"--...": '7',
		"---..": '8',
		"----.": '9',
		"/":     ' ',
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
		// Морзе в текст
		words := strings.Split(str, " / ")
		var result []string
		for _, word := range words {
			letters := strings.Split(word, " ")
			var decodedWord strings.Builder
			for _, l := range letters {
				if l == "" {
					continue
				}
				ch, ok := morseToText[l]
				if !ok {
					return "", errors.New("некорректный код Морзе: " + l)
				}
				decodedWord.WriteRune(ch)
			}
			result = append(result, decodedWord.String())
		}
		return strings.Join(result, " "), nil
	} else {
		// Текст в Морзе
		var morse []string
		upper := strings.ToUpper(str)
		for _, c := range upper {
			code, ok := textToMorse[c]
			if !ok {
				return "", errors.New("недопустимый символ в тексте: " + string(c))
			}
			morse = append(morse, code)
		}
		return strings.Join(morse, " "), nil
	}
}
