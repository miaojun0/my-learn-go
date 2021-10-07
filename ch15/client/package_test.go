package client

import (
	"fmt"
	"mj-learn-go/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {

	if ret, err := series.GetFibonacci(8); err != nil {
		if err == series.LessThanTwoError {
			fmt.Println("it is less than 2")
		} else if err == series.MoreThanHundredError {
			fmt.Println("it is more than 100")
		}
	} else {
		fmt.Println(ret)
	}

}
