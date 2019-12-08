package main

import (
	"fmt"
	"strings"
)

//匿名函数和闭包
//定义一个函数它的返回值是一个函数
//把函数作为返回值
func a(name string) func(){
	return func(){
		fmt.Println("Hello ", name)
	}
}

//闭包进阶示例2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//闭包进阶示例3
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	func() {
		fmt.Println("匿名函数")
	}()
	//闭包 = 函数 + 外层变量的引用
	r := a("World")//r此时是一个闭包
	r()//相当于执行了a函数中的匿名函数

	ms := makeSuffixFunc(".txt")
	mr := ms("Hello world")
	fmt.Println(mr)

	ca, cs := calc(100)
	car := ca(200)//base = 100 + 200
	csr := cs(200)//base = 300 - 200
	fmt.Println(car, csr)
}