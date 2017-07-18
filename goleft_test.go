package goleft

import "testing"

func TestLeftpad(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{LeftPad("foobar", 6), "foobar"},
		{LeftPad("foobar", 5), "foobar"},
		{LeftPad("foobar", -1), "foobar"},
		{LeftPad("foobar", 6, '1'), "foobar"},
		{LeftPad("foobar", 5, '1'), "foobar"},
		{LeftPad("foobar", -1, '1'), "foobar"},
		{LeftPad("foobar", 8, ' '), "  foobar"},
		{LeftPad("foobar", 8, '0'), "00foobar"},
		{LeftPad("0", 3, '1'), "110"},
		{LeftPad("true", 7), "   true"},
		{LeftPad("", 2), "  "},
		{LeftPad("foo", 5), "  foo"},
		{LeftPad("foobar", 6), "foobar"},
		{LeftPad("1", 2, '0'), "01"},
		{LeftPad("好吗", 3, '你'), "你好吗"},
		{LeftPad("好吗", 5, '你'), "你你你好吗"},
		{LeftPad("1", 19, '0'), "0000000000000000001"},
		{LeftPad("foo", 2), "foo"},
		{LeftPad("foo", 3), "foo"},
		{LeftPad("foo", 4), " foo"},
		{LeftPad("foo", 5), "  foo"},
		{LeftPad("foo", 12), "         foo"},
		{LeftPad("foo", 13), "          foo"},
		{LeftPad("foo", 2, ' '), "foo"},
		{LeftPad("foo", 3, ' '), "foo"},
		{LeftPad("foo", 4, ' '), " foo"},
		{LeftPad("foo", 5, ' '), "  foo"},
		{LeftPad("foo", 12, ' '), "         foo"},
		{LeftPad("foo", 13, ' '), "          foo"},
		{LeftPad("1", 2, '0'), "01"},
		{LeftPad("1", 2, '-'), "-1"},
		{LeftPad("foo", 4, '*'), "*foo"},
		{LeftPad("foo", 5, '*'), "**foo"},
		{LeftPad("foo", 6, '*'), "***foo"},
		{LeftPad("foo", 7, '*'), "****foo"},
		{LeftPad("foo", 103, '*'), "****************************************************************************************************foo"},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s, actual %s", v.expected, v.s)
		}
	}
}

func BenchmarkLeftpad(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LeftPad("foobar", 6)
		LeftPad("foobar", 5)
		LeftPad("foobar", -1)
		LeftPad("foobar", 6, '1')
		LeftPad("foobar", 5, '1')
		LeftPad("foobar", -1, '1')
		LeftPad("foobar", 8, ' ')
		LeftPad("foobar", 8, '0')
		LeftPad("0", 3, '1')
		LeftPad("true", 7)
		LeftPad("", 2)
		LeftPad("foo", 5)
		LeftPad("foobar", 6)
		LeftPad("1", 2, '0')
		LeftPad("好吗", 3, '你')
		LeftPad("好吗", 5, '你')
		LeftPad("1", 19, '0')
		LeftPad("foo", 2)
		LeftPad("foo", 3)
		LeftPad("foo", 4)
		LeftPad("foo", 5)
		LeftPad("foo", 12)
		LeftPad("foo", 13)
		LeftPad("foo", 2, ' ')
		LeftPad("foo", 3, ' ')
		LeftPad("foo", 4, ' ')
		LeftPad("foo", 5, ' ')
		LeftPad("foo", 12, ' ')
		LeftPad("foo", 13, ' ')
		LeftPad("1", 2, '0')
		LeftPad("1", 2, '-')
		LeftPad("foo", 4, '*')
		LeftPad("foo", 5, '*')
		LeftPad("foo", 6, '*')
		LeftPad("foo", 7, '*')
		LeftPad("foo", 103, '*')
	}
}
