package array

import "testing"

func TestArrayInit(t *testing.T) {
	var a [3]int
	t.Log(a)

	b := new([3]int)
	t.Log(*b)

	var c = [4]int{1, 2, 3, 4}

	var d = [...]int{1, 2, 3}

	t.Log(c, d)
}

func TestArrayTravel(t *testing.T) {
	e := [...]string{"1", "2", "3"}
	for i := 0; i < len(e); i++ {
		t.Log(e[i])
	}

	for i, v := range e {
		t.Log(i, v)
	}
}

func TestArraySelection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 78, 88, 9, 2, 324, 23, 4, 234, 23, 234}

	a := arr[:]
	t.Log(a)

	b := arr[:3]
	t.Log(b)
}
