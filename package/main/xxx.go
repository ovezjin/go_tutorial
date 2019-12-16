package main

//当代码都放在$GOPATH目录下的话
//我们导入包要从$GOPATH/src继续写起
import (
	"fmt"
	"go_zum.com/tutorial/package/calc"
)

func main() {
	fmt.Println("hello")
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}