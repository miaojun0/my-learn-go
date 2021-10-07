package customer_type

import (
	"fmt"
	"time"
)

type IntConv func(int) int

func timeSpent(inner IntConv) IntConv {
	return func(i int) int {
		start := time.Now()
		ret := inner(i)
		fmt.Println(time.Since(start).Seconds())
		return ret
	}
}
