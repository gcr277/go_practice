package main

import (
	"fmt"
	"runtime"
)

func main(){
	cpuNum := runtime.NumCPU()
	fmt.Println("logic cpu num is:", cpuNum)
	
	// 可以自己设置使用多少cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	

}