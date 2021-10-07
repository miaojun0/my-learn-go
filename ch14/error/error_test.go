package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n must be not less than 2")
var MoreThanHundredError = errors.New("n must be not more than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, MoreThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestFib(t *testing.T) {
	//t.Log(GetFibonacci(10))
	if ret, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("it is less than 2")
		} else if err == MoreThanHundredError {
			fmt.Println("it is more than 100")
		}
	} else {
		fmt.Println(ret)
	}
}
