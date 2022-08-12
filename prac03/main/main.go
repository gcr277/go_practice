package main

import (
	"fmt"
	utls "go_practice/prac03/utils"  // 引包时自动从 $GOPATH/src 和 $GOROOT/src 下寻找
			// 为包取别名，取别名后只能用别名
)

// 执行顺序：被引包的全局变量->被引包的init函数->本文件全局变量定义->本文件init函数->main
// 同一个包多次被引用时，其init函数也只执行一次（在第一次被引时）
var A int = testGlobal()
func testGlobal() int {
	fmt.Println("Global var defination")
	return 555
}
func init(){
	fmt.Println("init()...")
}

func main(){
	var a float64 = 3.0
	var b float64 = 2.0
	fmt.Println(utls.SumAndSub(a, b))
	// utls.SumAndSub(a, b) 相当于 _, _ = utls.SumAndSub(a, b)
	utls.Test()
	fmt.Println("................")
	fmt.Println(utls.GAfunc(int(a), int(b)))

}
