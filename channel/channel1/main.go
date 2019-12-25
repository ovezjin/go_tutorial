package main

import "fmt"

//channel demo
func main() {
	/* var ch1 chan int //引用类型，需要初始化
	ch1 = make(chan int, 1) */

	//ch1 := make(chan int) //无缓冲区通道，又称为同步通道，必须要有goroutine接收
	ch1 := make(chan int, 1) //带缓冲区通道，又称为异步通道
	ch1 <- 10                //发送值
	x := <-ch1
	fmt.Println(x)
	//len(ch1) //取通道元素的数量
	//cap(ch1) //拿到通道的容量
	close(ch1)
}
