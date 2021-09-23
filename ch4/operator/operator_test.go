package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{}

	t.Log(a == b)
	t.Log(1 &^ 0)
}
