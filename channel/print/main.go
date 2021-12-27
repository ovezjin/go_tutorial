package main

import "fmt"

func main() {
	num := 3
	chans := make([]chan int, num)
	done := make(chan bool)
	last := make(chan bool, 1)
	for i := 0; i < num; i++ {
		chans[i] = make(chan int)
		go func(i int) {
			for {
				select {
				case v := <-chans[i]:
					fmt.Printf("chan: %d, print: %d\n", i, v)
					last <- true
				case <-done:
					break
				}
			}
		}(i)
	}
	last <- true
	for i := 1; i <= 100; i++ {
		<-last
		if i%3 == 0 {
			chans[0] <- i
		} else if i%5 == 0 {
			chans[1] <- i
		} else {
			chans[2] <- i
		}
		if i == 100 {
			<-last
			close(done)
		}
	}
}
