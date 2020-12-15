package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}
type Ben struct {
	id   int
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben say, \"Hello my name is %s\"\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry say, \"Hello my name is %s\"\n", j.name)
}

// 会产生乱串的情况
// single machine word 是原子赋值，但interface 的赋值非单个machine word, 而是两个
func main() {
	var ben = &Ben{
		id:   10,
		name: "Ben",
	}
	var jerry = &Jerry{
		name: "Jerry",
	}
	//需要更新 IceCreamMaker 的Type 和 Data
	var maker IceCreamMaker = ben
	var loop0, loop1 func()
	loop0 = func() {
		maker = ben
		go loop1()
	}
	loop1 = func() {
		maker = jerry
		go loop0()
	}
	go loop0()
	for {
		maker.Hello()
	}
}
