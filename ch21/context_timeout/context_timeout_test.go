package context_timeout

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func cpuInfo(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("CPU 监控退出")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("CPU 信息读取完毕")
		}
	}
}

func TestContextCancel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	go cpuInfo(ctx, &wg)

	wg.Wait()
	fmt.Println("监控信息完成")
}
