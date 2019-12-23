package main

import (
	"fmt"
	"sync"
)

//goroutine demo2
var wg sync.WaitGroup

func main() { //开启一个主goroutine去执行main函数
	wg.Add(10) //计数牌
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Hello goroutine", i)
			wg.Done()
		}(i)
	}
	// go hello() //开启了一个goroutine去执行hello函数
	fmt.Println("Hello main")
	// time.Sleep(time.Second)
	wg.Wait() //阻塞，等待所有goroutine
}
