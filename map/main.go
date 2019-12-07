package main

import "fmt"

//map
func main() {
	//光声明map，但不初始化，则map的值为nil
	//map需要初始化
	var a map[string]int
	fmt.Println(a == nil)
	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	//map中添加键值对
	a["进"] = 100
	a["jin"] = 200
	fmt.Println(a)
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("%T\n", a)
	//声明map同时完成初始化
	b := map[int]bool{
		1:true,
		2:false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("%T\n", b)

	//判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["jin"] = 100
	scoreMap["jing"] = 200
	
	value, ok := scoreMap["ji"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("ji is exist", value)
	}else{
		fmt.Println("ji is't exist")
	}

	//遍历map for 
	//map时无序的，键值对和添加的顺序无关
	for k, v := range scoreMap{
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range scoreMap{
		fmt.Println(k)
	}
	//只遍历value
	for _, v := range scoreMap{
		fmt.Println(v)
	}
}