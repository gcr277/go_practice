package main

import (
	"fmt"
)

type Student struct{
	Name string
}
func typeJudge (items ...interface{}){
	for i, item := range items{
		switch item.(type){
			case bool :
				fmt.Printf("第%v个参数是bool类型，值为%v\n", i, item)
			case int, int32, int64:
				fmt.Printf("第%v个参数是int类型，值为%v\n", i, item)
			case float64, float32:
				fmt.Printf("第%v个参数是float类型，值为%v\n", i, item)
			case string:
				fmt.Printf("第%v个参数是string类型，值为%v\n", i, item)
			case Student:
				fmt.Printf("第%v个参数是Student类型，值为%v\n", i, item)
			case *Student:
				fmt.Printf("第%v个参数是*Student类型，值为%v\n", i, item)
			default:
				fmt.Printf("第%v个参数是未知类型，值为%v\n", i, item)
		}
	}
}
func main(){
	var v1 int32 = 10
	var v2 float64 = 33.3
	var v3 string = "hello"
	var v4 bool = true
	var v5 []int = []int{5,6,7}
	var v6 Student = Student{"Tom"}
	var v7 *Student = &Student{"Bob"}
	typeJudge(v1, v2, v3, v4, v5, v6, v7)
}