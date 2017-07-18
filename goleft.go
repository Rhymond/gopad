package goleft

import (
	"bytes"
	"unicode/utf8"
)

// LeftPad pads out the lefthand-side of strings
// with given rune or space by default
// it uses caching for common use cases
// and buffer for quicker concatenation
func LeftPad(s string, l int, ch ...rune) string {

	cache := []string{"", " ", "  ", "   ", "    ", "     ", "      ", "       ", "        ", "         "}
	l = l - utf8.RuneCountInString(s)

	if l <= 0 {
		return s
	}

	// By default padding character is space
	c := ' '
	if len(ch) > 0 {
		c = ch[0]
	}

	// Use cache if possible
	if c == ' ' && l < 10 {
		return cache[l] + s
	}

	var b bytes.Buffer
	for ; l != 0; l-- {
		b.WriteRune(c)
	}

	b.WriteString(s)
	return b.String()
}
