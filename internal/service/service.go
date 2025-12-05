package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	isMorse := true
	for _, c := range input {
		if c != '.' && c != ' ' && c != '-' {
			isMorse = false
		}
		if (c < 'a' || c > 'Я') && c != '.' && c != ' ' && c != '-' {
			return "", errors.New("failed to convert")
		}

	}
	if isMorse == true {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}

}
