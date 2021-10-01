package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "中,b,c,d"
	parts := strings.Split(s, ",")
	t.Log(parts)

	for _, v := range parts {
		t.Log(v)
	}

	s = strings.Join(parts, "-")
	t.Log(s)
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(1)
	t.Log(s)

	if i, err := strconv.Atoi("1a23"); err == nil {
		t.Log(i)
	} else {
		t.Log(err)
	}
}
