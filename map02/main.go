package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

//map
func main() {
	//按照某个固定顺序遍历
	var scoreMap = make(map[string]int, 100)

	//添加50个键值对
	for i:=0; i<50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//按照key值从小到大的顺序去遍历scoreMap
	//1.先取出所有的key存放在切片
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	//2.对key做排序
	sort.Strings(keys)//有序的
	//3.按照排序后的key对scoreMap排序
	for _, key := range keys{
		fmt.Println(key, scoreMap[key])
	}

	//元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8)//切片的初始化
	fmt.Println(mapSlice[0] == nil)
	//还需要将map元素初始化
	mapSlice[0] = make(map[string]int, 8)//完成map的初始化
	mapSlice[0]["jin"] = 100 
	fmt.Println(mapSlice)

	//值为切片的map
	var sliceMap = make(map[string][]int, 8)//只完成了map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		sliceMap["中国"] = make([]int, 8)//完成切片的初始化
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 100
		sliceMap["中国"][3] = 100
	}

	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	//统计单词出现的个数
	//"how do you do"
	//0.定义一个map[string]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	//1.字符串都有哪些单词
	words := strings.Split(s, " ")
	//2.遍历单词做统计
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}