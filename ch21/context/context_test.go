package context

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("正在读取 CPU 信息")
	go cpuInfo(&wg)

	wg.Wait()

	t.Log("信息监控完成")
}

func cpuInfo(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println("CPU 信息读取完毕")
	wg.Done()
}
