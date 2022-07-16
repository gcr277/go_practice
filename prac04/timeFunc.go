package main

import(
	"fmt"
	"time"
	"strconv"
)

func main(){
	// 1.日期时间信息
	now := time.Now()
	fmt.Printf("value = %v | type = %T\n", now, now)
	fmt.Printf("value = %v | type = %T\n", now.Year(), now.Year())
	fmt.Printf("value = %v | type = %T\n", now.Month(), now.Month())
	fmt.Printf("value = %v | type = %T\n", now.Day(), now.Day())
	fmt.Printf("value = %v | type = %T\n", now.Hour(), now.Hour())
	fmt.Printf("value = %v | type = %T\n", now.Minute(), now.Minute())
	fmt.Printf("value = %v | type = %T\n", now.Second(), now.Second())

	// 2.格式化
	fmt.Println(now.Format("2006/01/02 15:04:05")) // 里面日期和时间的数字是固定的
	fmt.Println(now.Format("2006-01-02 * 15|04|05"))

	// 3.时间常量及Sleep
	for i := 1; i <= 5; i++{
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )

	// 4.time的Unix方法：Unix时间戳和纳秒时间戳
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("花费时间为%v秒\n", end-start)
}

func test(){
	str := "hello"
	for i := 0; i < 100000; i++{
		str += strconv.Itoa(i)
	}
}