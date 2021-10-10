package cancel_by_close

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//任务是否已被取消
//实现原理：
//检查是否从 channel 收到一个消息，如果收到一个消息，我们就返回 true，代表任务已经被取消了
//当没有收到消息，channel 会被阻塞，多路选择机制就会走到 default 分支上去。
func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//执行任务取消
//因为 close() 是一个广播机制，所以所有的协程都会收到消息
func execCancel(cancelChan chan struct{}) {
	// close(cancelChan)会使所有处于处于阻塞等待状态的消息接收者（<-cancelChan)收到消息
	close(cancelChan)
}

//利用 CSP, 多路选择机制和 channel 的关闭与广播实现任务取消功能
func TestCancel(t *testing.T) {
	var wg sync.WaitGroup
	cancelChan := make(chan struct{}, 0)

	//启动 5 个协程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, cancelChan chan struct{}, wg *sync.WaitGroup) {
			//做一个 while(true) 的循环，一直检查任务是否有被取消
			for {
				if isCanceled(cancelChan) {
					fmt.Println(i, "is Canceled")
					wg.Done()
					break
				} else {
					//其它正常业务逻辑
					time.Sleep(time.Millisecond * 5)
				}
			}
		}(i, cancelChan, &wg)
	}
	//执行任务取消
	execCancel(cancelChan)
	wg.Wait()
}
