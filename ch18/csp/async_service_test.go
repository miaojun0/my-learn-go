package csp

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func OtherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task done")
}

func TestService(t *testing.T) {
	fmt.Println(Service())
	OtherTask()
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := Service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retch := AsyncService()
	OtherTask()
	fmt.Println(<-retch)
}
