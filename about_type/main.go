package main

import "fmt"

//自定义类型和类型别名示例

//1. 自定义类型

//MyInt 基于int类型 指向MyInt
type MyInt int

//2. 类型别名，指向原type

//NewInt int类型别名 指向int
type NewInt = int

func main(){
	var a MyInt
	var b NewInt
	fmt.Printf("type:%T value:%v\n", a, a)
	fmt.Printf("type:%T value:%v\n", b, b) 
}