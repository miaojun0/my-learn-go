package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 1
	for n < 5 {
		t.Log(n)
		n++
	}

	for i := 0; i < 5; i++ {
		t.Log(i)
	}
}
