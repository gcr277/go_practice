package main

import (
	"fmt"
)

func main(){
	res := sum(10, 20)
	fmt.Println("res...", res)
} // ok3, ok2, ok1, res

func sum(n1 int, n2 int)int{
	defer fmt.Println("ok1...", n1) 
	defer fmt.Println("ok2...", n2)
	n1++ // n1 11
	n2++ // n2 21
	res := n1 + n2
	fmt.Println("ok3...", res)
	return res
}
// 为了在函数执行完毕及时释放资源，go提供 defer
// defer 将语句以及当时的变量值压入独立的defer栈, 等到函数即将返回时再执行