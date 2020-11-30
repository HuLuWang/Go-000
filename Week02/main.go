package main

import (
	"Go-000/Week02/service"
	"fmt"
)

//在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
//是否应该 Wrap 这个 error，抛给上层
//需要，调用kit层产生的error需要在应用层处理并在最顶层打印错误

func main() {
	err := service.DoDao()
	if err != nil {
		fmt.Println("最顶层调用打印错误")
		fmt.Printf("原始错误堆栈信息: %+v\n", err)
	}
}
