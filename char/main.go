package main

import "fmt"

//字符

func main() {
	//byte uint8的别名 ASCII码
	//rune	int32的别名 UTF-8编码
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	s := "hello 世界"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	//处理中英文混合时
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}

	var (
		str   = "hello 沙河小王子"
		char  = []rune(str)
		count = 0
	)

	for _, v := range char {
		if v > 256 {
			count++
			fmt.Println(string(v))
		}
	}
	fmt.Println(count)
}
