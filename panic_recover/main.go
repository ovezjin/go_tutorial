package main

import "fmt"

//panic和recover

func a(){
	fmt.Println("func a")
}

//recover()必须搭配defer使用
//defer一定要在可能引发panic的语句之前定义
func b(){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c(){
	fmt.Println("func a")
}

func main() {
	a()
	b()
	c()
}
