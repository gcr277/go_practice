package main

import (
	"fmt"
)


/* 1.使用内置函数close可以关闭channel，一旦关闭后，只能取出不能写入
 * 2.一般应当由发送者关闭
 * 3.在最后的值从已关闭的信道中被接收后，任何对其的接收操作都会无阻塞的成功。
 */
func main(){
	var intChan chan int = make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)

	var num int
	var ok bool
	num, ok = <- intChan
	fmt.Printf("%v, %v\n", num, ok) // 100, true
	num, ok = <- intChan
	fmt.Printf("%v, %v\n", num, ok) // 200, true
	num, ok = <- intChan
	fmt.Printf("%v, %v\n", num, ok) // 0, false
	num, ok = <- intChan
	fmt.Printf("%v, %v\n", num, ok) // 0, false

	fmt.Println("*********************************")

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i:=0; i<100; i++{
		intChan2 <- i*2
	}

	// 遍历前必须关闭管道，否则一直向后读取，形成死锁；关闭后遍历结束会正常退出
	// 遍历同时也会取出
	close(intChan2) 
	for value := range intChan2{ // 无下标！
		fmt.Printf("%v\t", value)
	}
	fmt.Printf("\n")
	fmt.Printf("遍历结束后len:%v, cap:%v\n", len(intChan2), cap(intChan2))
	
}