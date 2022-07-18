package main

import (
	"fmt"
)

func main(){
	var arr [3][4]int = [...][4]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}} // 内层的类型不能省略
	fmt.Println(arr)
}