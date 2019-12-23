package main

import (
	"fmt"
	"sync"
)

//goroutine demo1
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("Hello Goroutine", i)
	wg.Done() //通知wg把计数器-1
}

func main() { //开启一个主goroutine去执行main函数
	wg.Add(10) //计数牌
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	// go hello() //开启了一个goroutine去执行hello函数
	fmt.Println("Hello main")
	// time.Sleep(time.Second)
	wg.Wait() //阻塞，等待所有goroutine
}
