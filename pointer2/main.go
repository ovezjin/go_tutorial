package main

import "fmt"

//指针new和make
func main() {
	var a *int//没有初始化 a = nil
	a = new(int)//初始化
	*a = 100
	fmt.Println(*a)

	var b map[string]int//没有初始化 b = nil
	b = make(map[string]int, 10)//make只用于slice, map, chan的内存创建
	b["jin"] = 100
	fmt.Println(b)
	//二者都是用来做内存分配的。
	//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
}