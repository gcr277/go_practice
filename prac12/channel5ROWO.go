package main

import (
	"fmt"
)

/* 1.管道默认为双向的
 * 2.可以声明为只读或只写
 */

func send(channel chan<- int){ //声明为只写
	channel <- 10
}
func recv(channel <-chan int)int{ //声明为只读
	return <-channel
}


func main(){

	var chan1 chan int = make(chan int, 5)
	
	send(chan1)
	num := recv(chan1)
	fmt.Println(num)

}