package main

import (
	"fmt"
	"encoding/json"
)

//结构体字段的可见性与JSON序列化

//如果一个Go语言包中定义的标识符是大写的，那么就是对外可见的
//如果一个结构体中的字段名是大写的，那么就是对外可见的
type student struct {
	ID int
	Name string
}

func newStudent(id int, name string) student {
	return student{
		ID: id,
		Name: name,
	}
}

//Tag
type class struct {
	Title string `json:"title"`
	Students []student `json:"stu_list"`
}

func main() {
	//创建一个班级变量c1
	c1 := class{
		Title: "火箭101",
		Students: make([]student, 0, 20),
	}
	//往班级c1中添加学生
	for i:= 0; i<10; i++ {
		//添加十个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}	
	fmt.Printf("%#v\n", c1)
	//JSON序列化:Go语言中的数据 -> JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err: ", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	//JSON序列化:JSON格式的字符串->Go语言中的数据
	jsonStr := `{"Title":"火箭101","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed, err: ", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}