package main

import (
	"sync"
	"time"
)

// 无缓冲通道
// 只有receive 和 sender 都准备好后才会返回， 否则会一直阻塞
func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- `foo`
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		println(`Message:` + <-c) // sleep 1 s 后 receiver 准备好了，解除阻塞
	}()
	wg.Wait()
}
