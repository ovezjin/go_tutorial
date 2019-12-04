package main

import "fmt"

//切片(slice)
func main(){
	//声明切片
	/* 
	var a []string
	var b []int

	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) 
	*/

	//基于数组得到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//切片再次得到切片
	c := b[0:len(b)]
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	//make函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	//通过len()函数获得切片的长度
	fmt.Println(len(d))
	//通过cap()函数获得切片的容量
	fmt.Println(cap(d))

	//nil
	//切片要初始化后才能使用
	var i []int //声明int切片
	var j =  []int{} //声明并且初始化
	k := make([]int, 0)
	if i == nil {
		fmt.Println("i是一个nil")
	}
	fmt.Println(i, len(i), cap(i))
	if j == nil {
		fmt.Println("j是一个nil")
	}
	fmt.Println(j, len(j), cap(j))
	if k == nil {
		fmt.Println("k是一个nil")
	}
	fmt.Println(k, len(k), cap(k))

	//切片的赋值，拷贝
	m := make([]int, 3)
	n := m
	n[0] = 100
	fmt.Println(m)
	fmt.Println(n)//指向同一个内存地址

	//切片的遍历
	s := []int{1, 2, 3, 4, 5}
	//根据索引遍历
	for i := 0; i < len(s); i ++ {
		fmt.Println(i, s[i])
	}
	//for range 遍历
	for index, value := range s {
		fmt.Println(index, value)
	}

	//切片append添加元素，并扩容
	var numSlice []int //此时它并没有申请内存
	/* 
	for i := 0; i < 10; i ++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	*/
	numSlice = append(numSlice, 1, 2, 3, 4, 5)
	fmt.Println(numSlice)
	numSlice1 := []int{11, 12, 13, 14, 15}
	numSlice = append(numSlice, numSlice1...)
	fmt.Println(numSlice)

	//切片的copy
	copySlice1 := []int{1, 2, 3, 4, 5}
	copySlice2 := make([]int, 5, 5)
	copySlice3 := copySlice2//指向同一内存
	copy(copySlice2, copySlice1)//不指向同一内存
	copySlice2[0] = 100
	fmt.Println(copySlice1)
	fmt.Println(copySlice2)
	fmt.Println(copySlice3)

	//切片删除元素
	citySlice := []string{"北京", "上海", "广州", "深圳"}
	citySlice = append(citySlice[0:2], citySlice[3:]...)
	fmt.Println(citySlice)
}