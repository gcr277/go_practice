package main

import (
	"fmt"
	"encoding/json"
)

type Cat struct{ // 同样用大小写控制使用范围
	Name string		`json:"name"`
	Age int			`json:"age"`
	Color string	`json:"color"`
}

func main(){
	// 创建实例方式1：
	var cat1 Cat
	cat1.Name = "no1"
	cat1.Age = 1
	cat1.Color = "white"
	fmt.Printf("%v\n", cat1)

	// 方式2：
	cat2 := Cat{Name : "no2" , Age : 2}
	cat2.Color = "orange"
	fmt.Printf("%v\n", cat2)

	// 方式3：
	var cat3Ptr *Cat = new(Cat)
	(*cat3Ptr).Name = "no3"
	(*cat3Ptr).Age = 3
	(*cat3Ptr).Color = "black"
	fmt.Printf("%v\n", *cat3Ptr)

	// 方式4：
	var cat4Ptr *Cat = &Cat{Name : "no4" , Age : 4, Color : "blue"}
	fmt.Printf("%v\n", *cat4Ptr)


	// 1. 结构体的字段在内存中是连续的
	// 2. 结构体强制转换需要字段的名字、类型、数量完全相同
	// 3. type MyCat Cat , 编译器认为MyCat和Cat是不相同的类型，但是可以强转
	// 4. 结构体的字段可以写上一个tag，该tag可以通过反射机制获取。常用于序列化和反序列化
	jsonStr, err := json.Marshal(cat2) // 此函数中用到反射
	if err != nil{
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", string(jsonStr))
	

}