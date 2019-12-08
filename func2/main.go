package main

import "fmt"

//func进阶之变量作用域
//定义全局变量
var num = 10

//定义函数
func testGlobal() {
	num := 100
	//函数中使用变量
	//1.先在在自己函数中查找变量
	//2.函数中找不到就往外层找全局变量
	name := "jin"
	fmt.Println(name)
	fmt.Println(num)
}

func add(x,y int) int {
	return x + y
}
func sub(x,y int) int {
	return x - y
}
//函数作为函数参数
func calc(x,y int, op func(int, int) int) int{
	return op(x, y)
}

func main() {
	testGlobal()
	//外层不能访问到函数里的内部变量(局部变量)
	//外层不能访问到语句块里的变量
	for i:=0;i<5;i++{
		fmt.Println(i)
	}
	//变量i此时只在for循环语句块中生效
	//函数时可以作为变量
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()
	r1 := calc(100, 200, add)
	fmt.Println(r1)
	r2 := calc(100, 200, sub)
	fmt.Println(r2)
}