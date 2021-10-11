package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 仅需任意任务完成

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	// 打印当前有多少协程
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	// 若不使用 buffer chan，会导致其他协程阻塞
	t.Log("After:", runtime.NumGoroutine())
}
