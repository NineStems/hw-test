package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(in string) (string, error) {
	runes := convertFromStr(in)
	if len(runes) == 0 {
		return "", nil
	}
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	var (
		current, next rune
		res           strings.Builder
		part          string
		err           error
	)
	l := len(runes)
	for i := 0; i < l-1; i++ {
		part = ""
		current = runes[i]
		next = runes[i+1]

		multiSelectSymbol := !unicode.IsDigit(current) && unicode.IsDigit(next)
		justAddSelectSymbol := unicode.IsLetter(current) && !unicode.IsDigit(next)
		andEndRowAfterPreEnd := i+1 == l-1 && unicode.IsLetter(next)
		isInvalidString := unicode.IsDigit(current) && unicode.IsDigit(next)
		endRow := i+1 == l-1 && unicode.IsLetter(next)

		switch {
		case multiSelectSymbol:
			part, err = returnMultiValue(current, next)
		case justAddSelectSymbol:
			part = string(current)
			if andEndRowAfterPreEnd {
				part += string(next)
			}
		case isInvalidString:
			err = ErrInvalidString
		case endRow:
			part = string(next)
		}
		if err != nil {
			return "", err
		}

		if part != "" {
			res.WriteString(part)
		}
	}

	return res.String(), nil
}

func convertFromStr(in string) []rune {
	res := make([]rune, 0, len(in)/2) // Умышленно создаем массив рун такой же длинны, как и строку
	for _, symbol := range in {
		res = append(res, symbol)
	}
	return res
}

func returnMultiValue(r, c rune) (string, error) {
	count, err := strconv.Atoi(string(c))
	if err != nil {
		return "", err
	}
	if count == 0 {
		return "", nil
	}
	res := strings.Repeat(string(r), count)
	return res, nil
}
