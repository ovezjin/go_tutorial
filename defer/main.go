package main

import "fmt"


//return defer顺序
//1. 返回值=x
//2. 运行defer
//3. RET指令
func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Println("f1 ",x)
	}()
	return x//返回5, int = x = 5, ret = 5
}

func f2() (x int) {
	defer func() {
		x++
		fmt.Println("f2 ",x)
	}()
	return 5//返回6, x = 5, call defer(), x = 5 + 1, ret = x 
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
		fmt.Println("f3 ",x)
	}()
	return x//返回5, y = x = 5, x = 5 + 1, ret = y
}
func f4() (x int) {
	defer func(x int) {
		x++
		fmt.Println("f4 ",x)
	}(x)
	return 5//返回5, x = 1, ret = 5
}

//defer:延迟执行
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}