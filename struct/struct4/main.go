package main

import "fmt"

//结构体构造函数：构造一个结构体实例的函数
//结构体是值类型
type person struct {
	name, city string
	age int8
}

// 构造函数通常返回指针（保证性能）
/* 
func newPerson(name, city string, age int8) person{
	return person{
		name:name,
		city:city,
		age:age,
	}
}
 */
func newPerson(name, city string, age int8) *person{
	return &person{
		name:name,
		city:city,
		age:age,
	}
}

func main() {
	p1 := newPerson("jin", "xiamen", 23)
	fmt.Printf("type:%T value:%#v\n", p1, p1)
}