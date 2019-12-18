package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//file文件操作

//1.读
//读取文件
func readFormfile() {
	//打开文件
	fileObj, err := os.Open("./xx.txt") //相对路径
	if err != nil {
		fmt.Printf("file open fail err:%v\n", err)
		return
	}
	defer fileObj.Close() //关闭文件

	//读取文件内容
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("read form file failed, err:%v\n", err)
		return
	}
	fmt.Printf("read %d bytes from file.\n", n)
	fmt.Println(string(tmp))
}

//循环读取
func readAll() {
	//打开文件
	fileObj, err := os.Open("./xx.txt") //相对路径
	if err != nil {
		fmt.Printf("file open fail err:%v\n", err)
		return
	}
	defer fileObj.Close() //关闭文件

	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			//把当前读了多少字节打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read form file failed, err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

//read by bufio
func readBybufio() {
	//打开文件
	fileObj, err := os.Open("./xx.txt") //相对路径
	if err != nil {
		fmt.Printf("file open fail err:%v\n", err)
		return
	}
	defer fileObj.Close() //关闭文件
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf(line)
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio faile, err:%v\n", err)
			return
		}
		fmt.Printf(line)
	}

}

//ioutil
func readFormIoutil() {
	content, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Printf("read file by ioutil. err:%v\n", err)
	}
	fmt.Println(string(content))
}

//2.写
func write() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed err:%v\n", err)
	}
	defer fileObj.Close()
	str := "Jin"
	fileObj.Write([]byte(str))   //[]byte
	fileObj.WriteString("Hello") //string
}

//write by bufio
func writeBybufio() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed err:%v\n", err)
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("hello\n") //将数据先写入缓存
	writer.Flush()                //将缓存中的内容写入文件
}

//write by ioutil
//会先将文件清空
func writeByioutil() {
	str := "golang "
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	// readAll()
	// readBybufio()
	// readFormIoutil()
	// write()
	// writeBybufio()
	writeByioutil()
}
