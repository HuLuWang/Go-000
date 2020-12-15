package main

import (
	"sync"
	"time"
)

// 有缓冲通道，当通道满时将阻塞发送者，直到通道中的值被取走，发送者将消息发送出去
// 当通道为空时，将锁住接受者，直到通道有信息被发送，接受者取走
func main() {
	c := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- `foo`
		c <- `bar`
	}()
	
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		println(`message:` + <-c)
		println(`message:` + <-c)
	}()
	wg.Wait()
}
