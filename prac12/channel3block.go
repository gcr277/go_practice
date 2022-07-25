package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int){
	for i := 0; i < 50; i++{  // 写满了，写方会阻塞，但只要有协程还在读出就不会死锁
		intChan <- i
		fmt.Printf("write %v\n", i)
		time.Sleep(time.Millisecond * 100) // 写得慢，读方会阻塞，只要有协程还在写入就不会死锁
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool){
	var v int
	var ok bool
	for {
		v, ok = <- intChan
		if !ok{
			break
		}
		fmt.Printf("read %v, %v\n", v, ok)
	}
	exitChan <- true
	close(exitChan)
}

func main(){
	intChan := make(chan int, 5)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	
	var flag bool
	var ok bool
	for {
		flag, ok = <- exitChan
		if !ok{
			break
		}
		fmt.Printf("flag = %v, ok = %v\n", flag, ok)
	}
	fmt.Printf("flag = %v, ok = %v\n", flag, ok)
	// for{
	// 	time.Sleep(time.Second)
	// 	flag, ok = <- exitChan
	// 	fmt.Printf("%v, %v\n", flag, ok)
	// }

}