package main

import "fmt"

//for循环
func main() {
	//for基本
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//省略初始语句，但是必须保留分号
	var n = 0
	for ; n < 10; n++ {
		fmt.Println(n)
	}

	//省略初始语句和结束语句
	var m = 10
	for m > 0 {
		fmt.Println(m)
		m--
	}

	//无限循环
	/* for {
		fmt.Println("Hello")
	} */

	//break跳过整个循环
	for j := 0; j < 5; j++ {
		fmt.Println(j)
		if j == 3 {
			break
		}
	}

	//continue跳过这次循环
	for k := 0; k < 5; k++ {
		if k == 3 {
			continue
		}
		fmt.Println(k)

	}
}
