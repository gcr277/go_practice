package main

import (
	"fmt"
	"time"
)


func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello%d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

/*
 * goroutine中使用recover来处理协程中出现的panic，避免整个程序崩溃
 */
func test(){
	defer func(){
		// 捕获test()抛出的error
		if err := recover(); err != nil{
			fmt.Printf("test()...err:%v\n", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" // panic:向空map中赋值
}
func main()  {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Printf("main...hello%d\n", i)
		time.Sleep(time.Millisecond * 100)
	}

}