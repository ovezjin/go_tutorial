package main

import (
	"fmt"
	"time"
)

//time
func main() {
	now := time.Now()//时间对象
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	
	//时间戳：1970：1.1到现在的毫秒数
	timeStamp1 := now.Unix() //时间戳
	timeStamp2 := now.UnixNano()//纳秒时间戳
	fmt.Println(timeStamp1, timeStamp2)

	//将时间戳转化为具体的时间格式
	t := time.Unix(1576579217, 0)
	fmt.Println(t)

	//时间间隔
	//Duration是一个类型
	// time.Sleep(5*time.Second)
	n := 5 //type int
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("over")

	//时间操作
	//add
	// + 1hour
	t2 := now.Add(time.Hour)
	fmt.Println(t2)
	//sub
	fmt.Println(t2.Sub(now))

	//定时器
	/* for tmp := range time.Tick(time.Second){ //返回一个channel
		fmt.Println(tmp)
	} */

	//时间格式化:Y     m  d  H  M  S
	//          2006 01 02 15 04 05
	ret1 := now.Format("2016-01-02 15:04:05")
	ret2 := now.Format("2016-01-02 03:04:05 PM")//AMPM制
	fmt.Println(ret1)
	fmt.Println(ret2)

	//解析字符串类型的时间
	timeStr := "2019/12/18 00:00:00"
	//1.拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.普通解析
	timeObj1, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("parse failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj1)
	//2.根据时区解析
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("parse failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj2)

}