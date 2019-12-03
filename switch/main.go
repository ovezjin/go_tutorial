package main

import "fmt"

// switch
func main() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入")
	}

	//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	//分支还可以使用表达式
	//fallthrough语法可以执行满足条件的case的下一个case
}
