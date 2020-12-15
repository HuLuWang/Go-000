package main

import (
	"sync"
	"time"
)

// goroutine1 获取锁的次数远超 goroutine2
// 解析：
// 1. goroutine1 获得锁并休眠的过程中， goroutine2试图获取锁，它将被添加入到锁的队列中(FIFO), goroutine2 进入等待状态
// 2. goroutine1 sleep完并释放锁， 将通知队列唤醒 goroutine2, 此时goroutine2标记为可运行的 并等待Go的调度程序（Go scheduler）在线程上运行
// 3. goroutine2 等待运行时， goroutine1 再次请求锁，并持有锁，（那个贪睡的goroutine2 sleep了100ms）
// 4. goroutine2 尝试获取锁 发现锁已经被其他的 goroutine持有， 重新进入到锁的队列中并进入 等待状态
func main() {
	done := make(chan bool, 1)
	// 锁
	var mu sync.Mutex
	// goroutine1 获取长时间的锁
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond) //goroutine1 获得锁，并休眠100ms
				mu.Unlock()
			}
		}
	}()
	// goroutine2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		mu.Unlock()
	}
	done <- true
}
