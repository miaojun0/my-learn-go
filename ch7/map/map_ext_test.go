package _map

import (
	"fmt"
	"testing"
)

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](3), m[3](4))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	setExisting(mySet, 3)
	t.Log(len(mySet)) // 1

	mySet[3] = true
	t.Log(len(mySet)) // 2

	delete(mySet, 1)

	setExisting(mySet, 3)
	t.Log(len(mySet)) // 1
}

func setExisting(set map[int]bool, n int) {
	if set[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
}
