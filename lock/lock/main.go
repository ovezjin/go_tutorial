package main

import (
	"fmt"
	"sync"
)

//多个goroutine并发操作全局变量
//互斥锁：能够保证只有一个goroutine可以访问共享资源
var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func add(){
	for i:=0;i<500000;i++{
		lock.Lock()//加锁
		x = x + 1
		lock.Unlock()//释放锁
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}