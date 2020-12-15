package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

// parent ctx可以派生出很多的子ctx(cow的方式)，
// 当parent ctx退出， 其子ctx以及向光的goroutine都将退出
func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // return context deadline exceeded
	}
}
