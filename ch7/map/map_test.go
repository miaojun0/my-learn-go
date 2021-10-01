package _map

import (
	"testing"
)

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	t.Log(m1, len(m1))

	m2 := map[int]int{}
	t.Log(m2, len(m2))

	var m3 map[int]int
	t.Log(m3, len(m3))

	m4 := make(map[int]int, 10)
	t.Log(m4, len(m4))
}

func TestNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[2])

	if v, ok := m1[3]; ok {
		t.Log("exist", v)
	} else {
		t.Log("not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := make(map[int]int, 8)
	for i := 0; i < 10; i++ {
		m1[i] = i
		t.Logf("address: %p, len: %d", &m1, len(m1))
	}

	for k, v := range m1 {
		t.Log(k, v)
	}

}
