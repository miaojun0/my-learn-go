package context_cancel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func cpuInfo(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	go memInfo(ctx, wg)
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

func memInfo(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("内存监控退出")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("内存信息读取完毕")
		}
	}
}

func TestContextCancel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	go cpuInfo(ctx, &wg)
	//go memInfo(ctx, &wg)

	time.Sleep(time.Second * 6)
	cancel()
	wg.Wait()
	fmt.Println("监控信息完成")
}
