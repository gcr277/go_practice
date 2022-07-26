package main

import (
	"fmt"
	"sort"
	_"time"
	_"sync"
)
const IntNUM int = 8000


func putInt(intChan chan<- int){
	for i := 1; i <= IntNUM; i++{
		intChan <- i
	}
	close(intChan)
}

func movePrimeIntoChan(intChan <-chan int, primeChan chan<- int, finishChan chan<- bool){
	var tmp int
	var intChanOk bool
	for{
		tmp,intChanOk = <- intChan
		if !intChanOk{
			break
		}
		if isPrime(tmp){
			primeChan <- tmp
		}
	}
	fmt.Println("有一个movePrimeIntoChan已退出")
	finishChan <- true
}

/* 虽然slice为引用传递，但其本质是一个结构体对象，此函数中使用了append函数，
 * 修改了slice结构体内隐藏数组的指针，因此若想影响到main中的slice，此处应传入slice的地址，
 * 否则main中的slice仍指向扩容前的隐藏数组
 */
func insertPrimeFromChanIntoSli(primeChan <-chan int, primeSliPtr *([]int), exitChan chan<- bool){
	var tmp int
	var primeChanOk bool
	var index int
	for{
		//----取出管道中的素数
		tmp, primeChanOk = <- primeChan
		if !primeChanOk{  // 主线程close(primeChan)，所有元素被取出
			break
		}
		fmt.Printf("取出管道素数:%v\n", tmp)

		//----定位
		index = sort.SearchInts(*primeSliPtr, tmp)

		//----插入切片中
		*primeSliPtr = append(*primeSliPtr, 0)
		copy((*primeSliPtr)[index+1:], (*primeSliPtr)[index:])
		(*primeSliPtr)[index] = tmp
		//--------------------------------
	}
	fmt.Println("素数已全部取出并插入切片")
	//fmt.Printf("1~%v的素数为:\n%v\n%p\n", IntNUM, *primeSliPtr, *primeSliPtr)
	exitChan <- true
	close(exitChan)  // 主线程可以退出
	
}


func isPrime(n int)bool{

	if n == 1{
		return false
	}
	if n == 2{
		return true
	}
	//从2遍历到n-1，看看是否有因子
	for i := 2; i < n; i++ {
		if n%i == 0 {
			//发现一个因子
			return false
		}
	}
	return true
}

func main(){
	intChan := make(chan int, 200)
	primeChan := make(chan int, 100)
	finishChan := make(chan bool, 4)
	exitChan := make(chan bool, 1)
	primeSli := make([]int, 0, 100)
	go putInt(intChan)
	for i:=0; i<4; i++{
		go movePrimeIntoChan(intChan, primeChan, finishChan) 
	}
	go insertPrimeFromChanIntoSli(primeChan, &primeSli, exitChan)
	
	go func(){
		for i:=0; i<4; i++{
			<- finishChan
		}
		fmt.Println("got 4 finish flag. close primeChan...")
		close(primeChan)
	}()
	
	var exitChanOk bool
	for{
		_, exitChanOk = <- exitChan
		if !exitChanOk{
			break
		}
	}
	//var wg sync.WaitGroup
	//wg.Wait()
	fmt.Printf("main: 1~%v的素数为:\n%v\n%p\n", IntNUM, primeSli, primeSli)

}