package condition

import (
	"strconv"
	"testing"
)

func TestCondition(t *testing.T) {
	for i := 0; i < 3; i++ {
		if num, err := someFunc(i); err == "" {
			t.Log(num)
		}
	}
}

func someFunc(n int) (string, string) {
	if n > 0 {
		return strconv.Itoa(n), ""
	} else {
		return "error", "error"
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6, 8:
			t.Log("Even")
		case 1, 3, 5, 7, 9:
			t.Log("Odd")
		default:
			t.Log("What")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("What")
		}
	}
}
