package main

import (
	"fmt"
)
// 常量只能是 int系列,float系列,bool,string  即基本类型
// 常量声明时必须赋值
// 仍然通过大小写控制访问范围
const ( 
	A = iota // 表示a赋值0, 后面依次+1
	B
	C
)
const ( 
	AA = iota 	// 0
	BB			// 1	
	CC			// 2
	DD, EE= iota, iota // 3, 3
)


func main(){
	fmt.Printf("A=%v B=%v C=%v\n", A, B, C)
	fmt.Printf("AA=%v BB=%v CC=%v DD=%v EE=%v\n", AA, BB, CC, DD, EE)
}