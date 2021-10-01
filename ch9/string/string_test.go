package string

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s, len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)

	s = "中"
	t.Log(len(s))

	c := []byte(s)
	t.Log(c, len(c))

	d := []rune(s)
	t.Log(d, len(d))

	t.Logf("中 unicode %x", d[0])
	t.Logf("中 utf8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人命共和国"
	for _, c := range s {
		t.Logf("%[1]c, %[1]d", c)
	}
}
