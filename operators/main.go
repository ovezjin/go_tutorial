package main

import "fmt"

//Go语言中的运算符
func main() {
	//算术运算符, ++, --是语句
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	//关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(4 <= 5)

	//逻辑运算符
	fmt.Println(10 > 5 && 1 == 1)
	fmt.Println(!(10 > 5))
	fmt.Println(1 > 5 || 1 == 1)

	//位运算符
	c := 1               //001
	d := 5               //101
	fmt.Println(c & d)   //001 => 1
	fmt.Println(c | d)   //101 => 5
	fmt.Println(c ^ d)   //100 => 4
	fmt.Println(1 << 2)  // 1 -> 100 => 4
	fmt.Println(4 >> 2)  // 100 -> 1 => 1
	fmt.Println(1 << 10) // 1 -> 100 => 4

	//赋值运算符
	var e int
	e = 5
	e += 5 // e = e + 5
	fmt.Println(e)
}
