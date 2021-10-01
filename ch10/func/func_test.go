package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(123), rand.Intn(123)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	ts := timeSpent(slowFun)
	t.Log(ts(10))

	t.Log(Sum(1, 2, 2, 3, 4, 5))
}

func timeSpent(inner func(int) int) func(int) int {
	return func(i int) int {
		start := time.Now()
		ret := inner(i)
		fmt.Println(time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(10)
	return op
}

func Sum(op ...int) int {
	sum := 0
	for _, i := range op {
		sum += i
	}
	return sum
}

func Clear() {
	fmt.Println("clean up resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("error")
}
