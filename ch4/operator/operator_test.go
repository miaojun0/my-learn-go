package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{}

	c := [3]int{2, 1, 3}

	t.Log(a == b, a == c)
}

const (
	Writable = 1 << iota
	Readable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Writable
	t.Log(a&Writable == Writable)
}
