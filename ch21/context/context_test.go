package context

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var stop chan bool = make(chan bool)

func cpuInfo(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("退出 CPU 监控")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("CPU 信息读取完毕")
		}

	}

}

func TestContext(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	go cpuInfo(&wg)
	time.Sleep(time.Second * 6)
	stop <- true
	wg.Wait()

	t.Log("信息监控完成")
}
