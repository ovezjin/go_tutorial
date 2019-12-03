package main

import "fmt"

//函数外的每个语句都必须以关键字开始（var、const、func等）
var x = 100
var y = "oveJin"

func foo() (int, string) {
	return 22, "ovezjin"
}

func main() {
	//标准声明格式
	/*整型和浮点型变量的默认值为0。
	字符串变量的默认值为空字符串。
	布尔型变量默认为false。
	切片、函数、指针变量的默认为nil
	*/
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	//批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	//声明变量同时指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)
	var name2, age2 = "jin", 22
	fmt.Println(name2, age2)
	//支持类型推导，让编译器根据初始值推出类型
	var name3 = "Jin"
	var age3 = 12
	fmt.Println(name3, age3)
	//短变量声明（在函数内部）
	m := 10
	fmt.Println(m)
	//匿名变量
	age4, _ := foo()
	_, name4 := foo()
	fmt.Println(age4, name4)
}
