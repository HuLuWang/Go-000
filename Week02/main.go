package main

import (
	"Go-000/Week02/service"
	"fmt"
)

func main() {
	err := service.DoDao()
	if err != nil {
		fmt.Println("最顶层调用打印错误")
		fmt.Printf("原始错误堆栈信息: %+v\n", err)
	}
}
