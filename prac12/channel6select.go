package main

import (
	"fmt"
	"time"
)

func main(){
	intChan := make(chan int, 10)
	for i:=0; i<10; i++{
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello"+fmt.Sprintf("%d", i)
	}
	// for{
	// 	<-intChan
	// }

	/*
	 * select 语句解决死锁
	 */
	for{
		select{
			case v:= <-intChan:
				fmt.Printf("get from intChan %v\n", v)
				time.Sleep(time.Second * 3)
			case v:= <-stringChan:
				fmt.Printf("get from stringChan %v\n", v)
			default:
				fmt.Printf("can not get\n")
				return

		}
	}

}