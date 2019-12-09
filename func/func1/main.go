package main

import "fmt"

//func
//定义有一个不需要返回值的函数
func sayHello() {
	fmt.Println("Hello world!")
}

//定义一个接收string类型的参数
func sayHello2(name string) {
	fmt.Println("Hello ", name)
}

//定义接收多个参数的函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

//参数类型简写
//初始化返回值，可以简写return
func intSum2(a ,b int) (ret int) {
	ret = a + b
	return
}

//函数接收可变参数, 在参数后面加...，表示可变参数
//可变参数在函数体中时切片
func intSum3(a ...int) int {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}

//固定参数和可变参数同时出现时，可变参数要放在最后
//go语言的函数中没有默认参数
func intSum4(a int, b ...int) int {
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	return ret
}

//定义一个具有多个返回值的函数
//多返回值也支持类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	//函数调用
	sayHello()
	name := "go"
	sayHello2(name)
	r := intSum(10, 20)
	fmt.Println(r)
	intSum3(10, 20)
	r1 := intSum3()
	r2 := intSum3(10)
	r3 := intSum3(10, 20)
	r4 := intSum3(10, 20, 30)
	fmt.Println(r1, r2, r3, r4)
	x, y := calc(100, 200)
	fmt.Println(x, y)
}