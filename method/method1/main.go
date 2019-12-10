package main

//方法的定义实例
import "fmt"

// Person是一个结构体
type Person struct {
	name string
	age int8
}

// NewPerson 是person类型的构造函数
func NewPerson(name string, age int8) *Person{
	return &Person{
		name: name,
		age: age,
	}
}

// Dream 是为Person类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言! \n", p.name)
}

//指针接收者指的是接收者的类型是指针
func (p *Person) setAge(newAge int8) {
	p.age = newAge
}

//值接收者
func (p Person) setAge2(newAge int8){
	p.age = newAge
}
/*
什么时候应该使用指针类型接收者:
1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
 */
func main() {
	p1 := NewPerson("Jin", 23)
	// (*p1).Dream()
	p1.Dream()
	fmt.Println(p1.age)
	p1.setAge(int8(24))
	fmt.Println(p1.age)
	p1.setAge2(int8(25))
	fmt.Println(p1.age)
}