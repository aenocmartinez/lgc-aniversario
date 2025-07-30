package util

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func ConvertStringToID(idStr string) (int64, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("ID inválido")
	}
	return id, nil
}

func ConvertStringToInt(valueStr string) (int, error) {
	i, err := strconv.Atoi(valueStr)
	if err != nil || i <= 0 {
		return 0, errors.New("valor entero inválido")
	}
	return i, nil
}

func ToCapitalCase(input string) string {
	words := strings.Fields(strings.ToLower(input))
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}
