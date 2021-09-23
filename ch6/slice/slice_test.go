package slice

import (
	"testing"
)

func TestSlice(t *testing.T) {
	var a []int
	t.Log(len(a), cap(a))

	a = append(a, 1)
	t.Log(len(a), cap(a))

	a = append(a, 1)
	t.Log(len(a), cap(a))

	a = append(a, 1)
	t.Log(len(a), cap(a))
}
func TestSliceGrowing(t *testing.T) {
	var arr []int
	for i := 0; i < 17; i++ {
		arr = append(arr, i)
		t.Logf("len: %d, cap: %d, id: %p", len(arr), cap(arr), &arr)
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2)
}
