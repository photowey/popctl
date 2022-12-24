package alphabet

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	EmptyString = ""
)

func CamelCase(src string) string {
	if src == EmptyString {
		return EmptyString
	}

	return strings.ToLower(src[:1]) + src[1:]
}

func Snake2Pascal(snakeCase string) string {
	snakeCase = strings.Replace(snakeCase, "_", " ", -1)
	snakeCase = cases.Title(language.Dutch).String(snakeCase)
	return strings.Replace(snakeCase, " ", "", -1)
}

func Snake2Camel(snakeCase string) string {
	return CamelCase(Snake2Pascal(snakeCase))
}

func PascalCase(src string) string {
	if src == EmptyString {
		return EmptyString
	}

	return strings.ToUpper(strings.ToLower(src[:1])) + src[1:]
}

func SnakeCase(src string) string {
	if src == EmptyString {
		return src
	}
	srcLen := len(src)
	result := make([]byte, 0, srcLen*2)
	caseSymbol := false
	for i := 0; i < srcLen; i++ {
		char := src[i]
		if i > 0 && char >= 'A' && char <= 'Z' && caseSymbol { // _xxx || yyy__zzz
			result = append(result, '_')
		}
		caseSymbol = char != '_'

		result = append(result, char)
	}

	snakeCase := strings.ToLower(string(result))

	return snakeCase
}
