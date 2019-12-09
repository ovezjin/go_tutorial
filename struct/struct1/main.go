package main

import "fmt"

//定义结构体struct，类似class
type person struct{
	name, city string//相同类型可简写
	age int8
}

func main() {
	var p1 person//实例化struct
	p1.name = "Jin"
	p1.city = "xiamen"
	p1.age = 23
	fmt.Printf("p=%#v\n", p1)
	fmt.Println(p1.name, p1.city, p1.age)

	//匿名结构体
	var user struct{
		name string
		married bool
	}
	user.name = "Jin"
	user.married = false
	fmt.Printf("%#v", user)
}