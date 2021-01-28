package main

import (
	"fmt"
	"net"
)

func main() {
	msgChan := make(chan string)
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("listening fail", err.Error())
	}
	fmt.Println("Start listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error", err.Error())
		}
		go readConn(conn, msgChan)
		go sendConn(conn, msgChan)
	}
}

func readConn(conn net.Conn, inputChan chan<- string) {
	defer conn.Close()
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from %s fail, msg: %s\n", conn.RemoteAddr().String(), err)
			break
		}
		inputChan <- string(buf[:n])
	}
}

func sendConn(conn net.Conn, outputChan <-chan string) {
	defer conn.Close()
	for {
		msg := <-outputChan
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("send data to %s fail, msg: %s\n", conn.RemoteAddr().String(), msg)
		}
	}
}
