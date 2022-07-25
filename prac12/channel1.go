package main

import (
	"fmt"
	"time"
)

/**
 * 1.channel本质就是一个队列，先进先出
 * 2.channel本身线程安全，多个routine访问时不需要加锁
 * 3.channel是有类型的，一个string类型的channel只能放string
 * 4.channel是引用类型
 * 5.channel必须初始化才能使用，即make后使用
 * 6.channel不会自动扩容，存不能超过容量，取至少长度为1
 */

func main(){
	var intChan chan int = make(chan int, 3)

	fmt.Printf("%v\n", intChan)

	num1 := 10
	num2 := 20
	num3 := 30
	intChan <- num1
	intChan <- num2
	intChan <- num3
	//intChan <- num1  // 不能超过容量，否则会报错
	fmt.Printf("len:%v , cap:%v\n", len(intChan), cap(intChan))
	num4 := <- intChan
	fmt.Printf("%v\n", num4)
	fmt.Printf("%v\n", <- intChan)
	fmt.Printf("%v\n", <- intChan)
	//fmt.Printf("%v\n", <- intChan)  // 管道为空，不能再取，否则报deadlock
	fmt.Printf("len:%v , cap:%v\n", len(intChan), cap(intChan))

	go func(intChan chan int){  // 协程可以取
		var v int
		var ok bool
		for{
			v, ok = <- intChan
		}
		fmt.Printf("%v, %v\n", v, ok)
	}(intChan)
	time.Sleep(time.Second*10)
}