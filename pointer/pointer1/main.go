package main

import "fmt"

//指针

//指针传值改值
func modify1(x int){
	x = 100
}

func modify2(y *int){
	*y = 100 
}

func main() {
	a := 10 //int
	b := &a //*int
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, b)
	//变量b是一个int类型的指针，它存储的是变量a的内存地址
	//指针取值
	c := *b
	fmt.Printf("%v %p\n", c, &c)
	//&取地址
	//*取值
	x := 1
	modify1(x)
	fmt.Println(x)
	modify2(&x)
	fmt.Println(x)
}