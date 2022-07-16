package main

import (
	"fmt"
)



func main(){
	a := 10
	b := 20
	fmt.Println(func(n1 int, n2 int)int{  // 匿名函数
		return n1 + n2
	}(a, b))

	afunc := func(n1 int, n2 int)int{  // 匿名函数可以赋给变量
		return n1 - n2
	}
	fmt.Println(afunc(a, b))
}


