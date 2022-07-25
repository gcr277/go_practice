package main

import (
	"fmt"
	"time"
)

func test(){
	for i:=0 ; i<10; i++{
		fmt.Println("test...",i)
		time.Sleep(time.Second)
	}
}
func main(){
	go test() // 开启一个协程
	for i:=0 ; i<10; i++{
		fmt.Println("main...",i)
		time.Sleep(time.Second)
	}
	
}
/* 如果主线程结束了，即使协程没有执行完也会退出
 * 
 * golang可以轻松开启上万个协程，其他编程语言的并发机制
 * 一般是基于线程的，开启过多的线程会消耗大量的资源，这就体现了golang在并发上的优势
 * 
 */

