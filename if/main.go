package main

import "fmt"

//if判断
func main() {
	//基本写法
	/* var score = 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	} */

	//特殊写法
	//score只在if结构中
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	score := 100
	if score >= 95 {
		fmt.Println("A++")
	} else if score > 90 {
		fmt.Println("A")
	}
}
