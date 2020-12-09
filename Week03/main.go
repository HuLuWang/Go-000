package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出

const (
	post1 = 8080
	post2 = 8081
)

// 启动服务
func startServer(ctx context.Context, port int) error {
	mux := http.NewServeMux()
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	s := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Println("start server on port", port)
	go func() {
		<-ctx.Done()
		err := s.Shutdown(ctx)
		log.Printf("shotdown server: %d, and err is %v", port, err)
	}()
	return s.ListenAndServe()
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	// 启动server 1
	g.Go(func() error {
		return startServer(ctx, post1)
	})
	// 启动server 2
	g.Go(func() error {
		return startServer(ctx, post2)
	})
	
	// 终止信号
	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-quit:
			return fmt.Errorf("kill by signal")
		case <-ctx.Done():
			log.Println(ctx.Err())
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Printf("shutdown server by %v", err)
	}
	os.Exit(1)
}
