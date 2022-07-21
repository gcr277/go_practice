package main

import(
	"fmt"
	"go_practice/prac07/packageA"
	"go_practice/prac07/packageC"
)

type B struct{
	packageA.A
	mem3 int
	packageC.C
	c packageC.C
}
func main(){
	b1 := B{}
	b1.Mem1 = 1		// b1.A.Men1 = 1 的简写，编译器就近访问层层向上找，方法也同理
	//b1.mem2 = 2  // 跨包继承不能直接访问被继承结构体的私有成员
	b1.mem3 = 3
	//b1.sayOK()
	fmt.Println(b1)

	b2 := B{}
	b2.c.Mem4 = 4	//当嵌套有名结构体时，属于组合，必须加上包含的结构名来访问其字段
	//b2.c.mem5 = 5
	fmt.Println(b2)

	b3 := B{
		packageA.A{Mem1:1},
		3,
		packageC.C{Mem4:4}, // 匿名
		packageC.C{Mem4:6}, // 有名
	}
	fmt.Println(b3)
	fmt.Println(b3.Mem4) // 访问了继承的packageC.Mem4, 而不是组合的c.Mem4

}