package main

import(
	"fmt"
)

func main(){
	// 1.len

	// 2.new 用来分配内存，主要是分配值类型。new返回的是指向类型的指针
	num1 := new(int)
	fmt.Printf("type = %T, value = %v, *value = %v\n", num1, num1, *num1)
	// 3.make 用来分配内存，主要分配引用类型。make返回的是引用类型本身

	// 4.close 用来关闭管道
}