package main

import "fmt"

//结构体指针
type person struct {
	name, city string
	age int8
}

func main() {
	var p2 = new(person)//p2是指针
	fmt.Printf("%T\n", p2)
	
	/* (*p2).name = "jin"
	(*p2).city = "xiamen"
	(*p2).age = 23 */
	//结构体指针可以直接赋值
	p2.name = "jin"
	p2.city = "xiamen"
	p2.age = 23
	fmt.Printf("%#v\n", p2)

	//取结构体的地址
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
	p3.name = "jing"
	p3.city = "cq"
	p3.age = 22
	fmt.Printf("%#v\n", p3)
}
