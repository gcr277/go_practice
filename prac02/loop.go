package main

import "fmt"

func main(){
	// for 和 for-range
	for i := 0; i < 10; i++{
		fmt.Println("for:",i)
	}

	var str1 string = "hello北京"
	for i := 0; i < len(str1); i++ { 	// 传统方式遍历字节，中文会出现乱码
		fmt.Printf("%d:%c\t", i, str1[i])		// 想遍历字符可以转换为切片   []rune(str1)
	}
	fmt.Println()

	for i, value := range str1 {  // for-range遍历字符
		fmt.Printf("%d:%c\t", i, value)
	}
	fmt.Println()
	// 没有while和dowhile
	// 可以用break退出循环
	// 死循环的写法：
	// for{}

	for j := 0; j<50; j++{
		var i int   // 每一次循环的i都是独立的
		i++
		fmt.Println(i)
	}
}

