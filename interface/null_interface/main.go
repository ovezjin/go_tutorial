package main

import "fmt"

//空接口
//接口中没有定义任何需要实现的方法时，该接口是一个空接口
//任意类型都实现了空接口 --> 空接口变量可以存储任意类型的值

//空接口一般不需要提前定义
type xxx interface{

}

//空接口的应用
//1.空接口作为函数的参数
//2.空接口可以作为map的value

func main() {
	var x interface{}
	x = "hello"
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	x = false
	fmt.Println(x)

	var m = make(map[string]interface{}, 16)
	m["name"] = "jin"
	m["age"] = 18
	m["hobby"] = []string{"篮球", "足球"}
	fmt.Println(m)

	//类型断言
	ret, ok := x.(string)//ok会返回一个布尔值，猜的不对ok=false,ret=string类型零值
	if !ok{
		fmt.Println("不是字符串类型")
	}else{
		fmt.Println("是字符串类型", ret)
	}

	//使用switch进行类型段言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型, value:%v\n", v)
	case bool:
		fmt.Printf("是bool类型，value:%v\n", v)
	case int:
		fmt.Printf("是int类型, value:%v\n", v)
	default:
		fmt.Printf("猜不到了，value:%v\n", v)
	}
}