package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a, b := 1, 1
	fmt.Println(a)
	fmt.Println(b)
	for i := 0; i < 5; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}
}

func TestExchange(t *testing.T) {

}
