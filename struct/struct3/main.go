package main

import "fmt"
//结构体的初始化
type person struct {
	name, city string
	age int8
}

func main() {
	//1. 键值对的初始化
	p4 := person{
		name: "jin",
		city: "xiamen",
		age: 23,
		//可以忽略某些字段
	}
	fmt.Printf("%#v\n", p4)

	p5 := &person{
		name: "jin",
		city: "xiamen",
		age: 23,
	}
	fmt.Printf("%#v\n", p5)

	//2. 值列表初始化
	p6 := person {
		"jin",
		"xiamen",
		23,//不可忽略，必须按照顺序
	}
	fmt.Printf("%#v\n", p6)
}