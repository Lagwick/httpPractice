package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(input string) bool {
	for _, ch := range input {
		if ch != '.' && ch != '-' && ch != ' ' && ch != '/' {
			return false
		}
	}
	return true
}

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("входная строка пуста")
	}

	if isMorse(input) {
		result := morse.ToText(input)
		if strings.TrimSpace(result) == "" {
			return "", errors.New("не удалось распознать код Морзе")
		}
		return result, nil
	}

	result := morse.ToMorse(input)
	if strings.TrimSpace(result) == "" {
		return "", errors.New("не удалось преобразовать текст в код Морзе")
	}
	return result, nil
}
