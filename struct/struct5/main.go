package main

import "fmt"

//结构体的匿名字段

//匿名结构体
//一个类型只能出现一次
type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		string: "Jin",
		int: 23,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int)
}

