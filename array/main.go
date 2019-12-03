package main

import "fmt"

//数组

func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)

	//数组的初始化
	//1.定义时使用初始值列表的方式初始化
	var city = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(city)
	fmt.Println(city[0])

	//2.编译器推导数组的长度
	var bool = [...]bool{true, false, false, true, false}
	fmt.Println(bool)

	//3.使用索引值初始化
	var lang = [...]string{1: "Golang", 3: "python", 7: "java"}
	fmt.Println(lang)
	fmt.Printf("%T\n", lang)

	//遍历数组
	//1.for循环
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}

	//2.for range
	for index, value := range city {
		fmt.Println(index, value)
	}

	for _, value := range city {
		fmt.Println(value)
	}

	//多为数组只有最外层呢给你用编译器推导
	cityArray := [4][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}
	fmt.Println(cityArray)
	fmt.Println(cityArray[2][0])

	//二位数组遍历
	for _, v1 := range cityArray {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}

func f1(a [3]int) {
	a[0] = 100
}
