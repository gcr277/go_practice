package main

import (
	"fmt"
)

func main(){

	// 声明方式1：声明，用make分配内存，再赋值
	// map[keytype]valuetype  keytype必须是可以用==判断的类型
	var m1 map[int]string = make(map[int]string, 5)
	m1[0] = "aaa"
	m1[1] = "bbb"
	m1[3] = "ddd"
	m1[2] = "ccc"
	m1[4] = "eee"
	fmt.Printf("m1 : %v\n", m1)

	// 声明方式2：声明时直接赋值
	m2 := map[int]string{
		10: "aaaaa",
		11: "bbbbb",
		12: "ccccc",
		13: "ddddd",
	}
	fmt.Printf("m2 : %v\n", m2)
	fmt.Println()

	// 增删改查
	m1[5] = "fff" // key存在则为更新，key不存在则为增加
	fmt.Printf("m1 : %v len = %v\n", m1, len(m1))

	delete(m1, 1) // key存在则删除
	delete(m1, 6) // key不存在不会报错，不作任何操作
	fmt.Printf("m1 : %v len = %v\n", m1, len(m1)) 
	// 没有清空map的函数，可以遍历删除或者make一个新的空间

	value, found := m1[6] // 查找
	if found {
		fmt.Printf("%v\n", value)
	}else{
		fmt.Printf("没有此元素\n")
	}


	// 遍历
	for key, value := range m1{
		fmt.Printf("key = %v, value = %v\n", key, value)
	}
}