package main

import "fmt"
//为什么需要接口
//接口是一种类型，是抽象的类型
//只要实现接口内的方法，则可以成为这个接口类型

type dog struct {}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct {}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

type person struct {
	name string
}

func (p person) say(){
	fmt.Println("啊啊啊")
}

//接口不管你是什么类型，它只管你应该实现什么方法
//定义一个类型，一个抽象的类型，只要实现say()这个方法的类型,都可以称为sayer
type sayer interface {
	say()
}


func da(arg sayer){
	arg.say()//不管传进来的是什么，我都要执行say
}

func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{
		name: "jin",
	}
	da(p1)

	var s sayer
	c2 := cat{}
	s = c2
	p2 := person{name: "Jing"}
	s = p2
	fmt.Println(s)
}