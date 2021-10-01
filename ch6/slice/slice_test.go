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
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "summer"
	t.Log(year, Q2, summer)
}

func TestSliceCompare(t *testing.T) {
	c1 := []int{1, 2, 3, 4}
	c2 := []int{1, 2, 3, 4}
	//t.Log(c1 == c2)
	//invalid operation: c1 == c2 (slice can only be compared to nil)
	t.Log(c1, c2)
}
