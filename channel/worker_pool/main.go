package main

import (
	"fmt"
	"time"
)

//worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker:%d stop job:%d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启3个goroutine
	//控制goroutine数量
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	//发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
