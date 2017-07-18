package gopad

import (
	"bytes"
	"unicode/utf8"
)

var cache = []string{"", " ", "  ", "   ", "    ", "     ", "      ", "       ", "        ", "         "}

// pad pads out a string
// it uses given rune slice with
// first value or space by default for filling a space
// it has caching for common use cases
// and buffer for quicker concatenation
func pad(s string, l int, ch []rune, di bool) string {
	l = l - utf8.RuneCountInString(s)

	if l <= 0 {
		return s
	}

	c := ' '
	if len(ch) > 0 {
		c = ch[0]
	}

	// Use cache if possible
	if c == ' ' && l < 10 {
		if di {
			return s + cache[l]
		}

		return cache[l] + s
	}

	var b bytes.Buffer

	if di {
		b.WriteString(s)
	}

	for ; l != 0; l-- {
		b.WriteRune(c)
	}

	if !di {
		b.WriteString(s)
	}

	return b.String()
}

// Left pads out the lefthand-side of strings
func Left(s string, l int, ch ...rune) string {
	return pad(s, l, ch, false)
}

// Right pads out the righthand-side of strings
func Right(s string, l int, ch ...rune) string {
	return pad(s, l, ch, true)
}
