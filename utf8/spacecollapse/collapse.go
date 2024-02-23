//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	var sb strings.Builder
	sb.Grow(len(input))

	prevSpace := false
	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		input = input[size:]

		isSpace := unicode.IsSpace(r)
		if !isSpace {
			sb.WriteRune(r)
		} else if !prevSpace {
			sb.WriteRune(' ')
		}
		prevSpace = isSpace

	}

	return sb.String()
}
