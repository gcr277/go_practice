package main

import (
	"fmt"
	"strings"
)



func main(){
	f := closePackageTest()
	f()
	f()
	f()
	/**********************************/
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("winter"))
	fmt.Println(f1("summer.jpg"))
}

func closePackageTest() func() {
	var str string = "closePackageTest.returnFunc..."
	var i int = 0
	return func(){
		i++
		fmt.Println(str, i)
	}
}

func makeSuffix(suffix string) func(string)string {

	return func(name string)string{
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}



// 闭包：可理解为返回的函数与其相关的引用环境组合的一个整体
// 闭包的意义在于不使用全局变量而达到全局变量的效果