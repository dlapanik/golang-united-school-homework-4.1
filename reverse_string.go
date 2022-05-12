package reverse_string

import (
	"strings"
)

func ReverseString(input string) (output string) {
	var b strings.Builder
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		b.WriteRune(runes[i])
	}

	return b.String()
}
