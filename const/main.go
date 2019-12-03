package main

import "fmt"

const (
	pi = 3.1415
	ei = 2.714
	n1 = 10
	n2
	n3
	m1 = iota
	m2
	_
	m3 = 100
	m4 = iota
)

const (
	_  = iota
	KB = 1 << (10 * iota) //1<<10
	MB = 1 << (10 * iota) //1<<20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

//const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
/*
iota是go语言的常量计数器，只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0。
const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
使用iota能简化定义，在定义枚举时很有用。
*/

func main() {
	fmt.Println(pi, ei, n1, n2, n3, m1, m2, m3, m4)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e, f)
}
