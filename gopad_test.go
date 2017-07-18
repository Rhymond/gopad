package gopad

import "testing"

func TestLeft(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{Left("foobar", 6), "foobar"},
		{Left("foobar", 5), "foobar"},
		{Left("foobar", -1), "foobar"},
		{Left("foobar", 6, '1'), "foobar"},
		{Left("foobar", 5, '1'), "foobar"},
		{Left("foobar", -1, '1'), "foobar"},
		{Left("foobar", 8, ' '), "  foobar"},
		{Left("foobar", 8, '0'), "00foobar"},
		{Left("0", 3, '1'), "110"},
		{Left("true", 7), "   true"},
		{Left("", 2), "  "},
		{Left("foo", 5), "  foo"},
		{Left("foobar", 6), "foobar"},
		{Left("1", 2, '0'), "01"},
		{Left("好吗", 3, '你'), "你好吗"},
		{Left("好吗", 5, '你'), "你你你好吗"},
		{Left("1", 19, '0'), "0000000000000000001"},
		{Left("foo", 2), "foo"},
		{Left("foo", 3), "foo"},
		{Left("foo", 4), " foo"},
		{Left("foo", 5), "  foo"},
		{Left("foo", 12), "         foo"},
		{Left("foo", 13), "          foo"},
		{Left("foo", 2, ' '), "foo"},
		{Left("foo", 3, ' '), "foo"},
		{Left("foo", 4, ' '), " foo"},
		{Left("foo", 5, ' '), "  foo"},
		{Left("foo", 12, ' '), "         foo"},
		{Left("foo", 13, ' '), "          foo"},
		{Left("1", 2, '0'), "01"},
		{Left("1", 2, '-'), "-1"},
		{Left("foo", 4, '*'), "*foo"},
		{Left("foo", 5, '*'), "**foo"},
		{Left("foo", 6, '*'), "***foo"},
		{Left("foo", 7, '*'), "****foo"},
		{Left("foo", 103, '*'), "****************************************************************************************************foo"},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s, actual %s", v.expected, v.s)
		}
	}
}

func TestRight(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{Right("foobar", 6), "foobar"},
		{Right("foobar", 5), "foobar"},
		{Right("foobar", -1), "foobar"},
		{Right("foobar", 6, '1'), "foobar"},
		{Right("foobar", 5, '1'), "foobar"},
		{Right("foobar", -1, '1'), "foobar"},
		{Right("foobar", 8, ' '), "foobar  "},
		{Right("foobar", 8, '0'), "foobar00"},
		{Right("0", 3, '1'), "011"},
		{Right("true", 7), "true   "},
		{Right("", 2), "  "},
		{Right("foo", 5), "foo  "},
		{Right("foobar", 6), "foobar"},
		{Right("1", 2, '0'), "10"},
		{Right("好吗", 3, '你'), "好吗你"},
		{Right("好吗", 5, '你'), "好吗你你你"},
		{Right("1", 19, '0'), "1000000000000000000"},
		{Right("foo", 2), "foo"},
		{Right("foo", 3), "foo"},
		{Right("foo", 4), "foo "},
		{Right("foo", 5), "foo  "},
		{Right("foo", 12), "foo         "},
		{Right("foo", 13), "foo          "},
		{Right("foo", 2, ' '), "foo"},
		{Right("foo", 3, ' '), "foo"},
		{Right("foo", 4, ' '), "foo "},
		{Right("foo", 5, ' '), "foo  "},
		{Right("foo", 12, ' '), "foo         "},
		{Right("foo", 13, ' '), "foo          "},
		{Right("1", 2, '0'), "10"},
		{Right("1", 2, '-'), "1-"},
		{Right("foo", 4, '*'), "foo*"},
		{Right("foo", 5, '*'), "foo**"},
		{Right("foo", 6, '*'), "foo***"},
		{Right("foo", 7, '*'), "foo****"},
		{Right("foo", 103, '*'), "foo****************************************************************************************************"},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s, actual %s", v.expected, v.s)
		}
	}
}

func BenchmarkPad(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pad("foobar", 6, []rune{' '}, false)
		pad("foobar", 5, []rune{' '}, true)
		pad("foobar", -1, []rune{' '}, false)
		pad("foobar", 6, []rune{'1'}, true)
		pad("foobar", 5, []rune{'1'}, false)
		pad("foobar", -1, []rune{' '}, true)
		pad("foobar", 8, []rune{' '}, false)
		pad("foobar", 8, []rune{'0'}, true)
		pad("0", 3, []rune{'1'}, false)
		pad("true", 7, []rune{' '}, true)
		pad("", 2, []rune{' '}, false)
		pad("foo", 5, []rune{' '}, true)
		pad("foobar", 6, []rune{' '}, false)
		pad("1", 2, []rune{'0'}, true)
		// pad("好吗", 3, []rune{'你'}, false)
		// pad("好吗", 5, []rune{'你'}, true)
		pad("1", 19, []rune{'0'}, false)
		pad("foo", 2, []rune{'0'}, true)
		pad("foo", 3, []rune{' '}, false)
		pad("foo", 4, []rune{' '}, true)
		pad("foo", 5, []rune{' '}, false)
		pad("foo", 12, []rune{' '}, true)
		pad("foo", 13, []rune{' '}, false)
		pad("foo", 2, []rune{' '}, true)
		pad("foo", 3, []rune{' '}, false)
		pad("foo", 4, []rune{' '}, true)
		pad("foo", 5, []rune{' '}, false)
		pad("foo", 12, []rune{' '}, true)
		pad("foo", 13, []rune{' '}, false)
		pad("1", 2, []rune{'0'}, true)
		pad("1", 2, []rune{'-'}, false)
		pad("foo", 4, []rune{'*'}, true)
		pad("foo", 5, []rune{'*'}, false)
		pad("foo", 6, []rune{'*'}, true)
		pad("foo", 7, []rune{'*'}, false)
		pad("foo", 103, []rune{'*'}, true)
	}
}
