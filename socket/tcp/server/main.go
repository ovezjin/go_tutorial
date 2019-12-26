package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server demo

func process(conn net.Conn) {
	defer conn.Close() //处理完之后要关闭这个连接
	//针对当前的连接做数据的发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Printf("read from cnn failed, err:%v\n", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("接受到的数据: ", recvStr)
		conn.Write([]byte("ok")) //把收到的数据返回客户端
	}
}

func main() {
	//开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Listen failed, err:%v\n", err)
		return
	}
	for {
		//等待客户端来建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		go process(conn)
	}

}
