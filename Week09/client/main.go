package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("client server start")
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	defer conn.Close()
	if err != nil {
		fmt.Println("connect fail", err.Error())
		return
	}
	input := bufio.NewReader(os.Stdin)
	var buf [1024]byte
	for {
		bytes, _, err := input.ReadLine()
		if err != nil {
			fmt.Println("read line fail", err.Error())
		}
		str := string(bytes)
		if str == "quit" {
			fmt.Println("exe quit!")
			return
		}
		n, err := conn.Write(bytes)
		if err != nil {
			fmt.Println("send data fail", err.Error())
		} else {
			fmt.Println("send data length", n)
		}
		read, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("receive data fail", err.Error())
		} else {
			fmt.Println("receive data length", read)
		}
		fmt.Printf("receive data from %s msg:%s\n", conn.RemoteAddr().String(), string(buf[:n]))
	}
}
