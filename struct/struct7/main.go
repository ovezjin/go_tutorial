package main

import "fmt"

//嵌套结构体的字段冲突

type Address struct{
	Province string
	City string
	UpdateTime string
}

type Email struct{
	Addr string
	UpdateTime string
}

type Person struct {
	Name string
	Gender string
	Age int
	//Address Address 嵌套一个结构体
	Address//匿名嵌套
	Email
}

func main() {
	p1 := Person{
		Name: "Jin",
		Gender: "男",
		Age: 23,
		Address: Address{
			Province: "重庆",
			City: "重庆",
			UpdateTime: "2019-12-10",
		},
		Email: Email{
			Addr: "jin@jin.com",
			UpdateTime: "2020-12-10",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Address)
	//fmt.Println(p1.UpdateTime)编译不通过
	//嵌套结构体包含多个相同名字的字段
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)
}