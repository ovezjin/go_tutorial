package main

import (
	"fmt"
	"math"
)

//基本数据类型

func main() {
	//十进制打印二进制
	n := 10
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	//八进制
	m := 075
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)
	//十六进制
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f)
	fmt.Printf("%d\n", f)
	// uint8
	var age uint8 //0~255
	age = 255
	fmt.Println(age)

	//浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	//复数
	/*
		complex64和complex128
		复数有实部和虚部，complex64的实部和虚部为32位，
		complex128的实部和虚部为64位。
	*/

	//布尔值
	//默认为false
	//不允许将整型转化为布尔型
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	//字符串
	s1 := "hello world"
	s2 := "你好世界"
	fmt.Println(s1)
	fmt.Println(s2)

	//打印windows平台路径
	fmt.Println("c:\\code\\go")
	//多行字符串用`
}
