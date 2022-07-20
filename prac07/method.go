package main

import (
	"fmt"
	_ "encoding/json"
)

type Cat struct{ // 同样用大小写控制使用范围
	Name string		
	Age int			
	Color string	
}

func (cat Cat)test() {  // 与Cat类型绑定, 结构体值传递
	fmt.Printf("%v is a cute cat\n", cat.Name)
}

func (cat Cat)String()string{
	res := fmt.Sprintf("Name=%v | Age=%v | Color=%v", cat.Name, cat.Age, cat.Color)
	return res
}

type SLI []int  // receiver type可以是结构体或自定义类型
func (slice SLI)testSli(){ //切片引用传递
	slice[0] = -1
}

func main(){
	cat1 := Cat{"mimi", 5, "grey"}
	cat1.test()

	s1 := SLI{1, 2, 3}
	fmt.Printf("%v\n", s1)
	s1.testSli()
	fmt.Printf("%v\n", s1)

	cat2 := Cat{"tom", 4, "blue"}
	fmt.Println(cat2) // 如果绑定实现了String()方法，则打印函数会自动调用String()
	fmt.Println(&cat2) // 自动将&cat2转化为cat2
	fmt.Printf("%v\n", cat2)
	fmt.Printf("%v\n", &cat2)

	fmt.Printf("%v\n", &cat1)
}

// 方法的自动取地址和自动取内容在函数中不可用
// 究竟传入什么，取决于方法的绑定是实例还是地址
