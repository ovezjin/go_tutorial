package main

import (
	"fmt"
	"runtime"
	"time"
)

//go maxprocs demo
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1) //只占用一个CPU核心
	go a()
	go b()
	time.Sleep(time.Second)
}
