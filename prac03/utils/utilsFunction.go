package utils

import "fmt"

func init(){
	fmt.Println("utils.init()...")
}

func SumAndSub(m float64, n float64) (float64, float64) {
	return m + n, m - n
}

func Test(){
	a := test1 //2.
	fmt.Printf("%T\n",a)
	
	fmt.Println(test2(a, 2))
}

func test1(a int) int {
	return a
}

type Test1Type func(int)int

func test2(exam1 Test1Type, b int) (ret int){  //3.   4.
	ret = exam1(b) + b
	return 
}
// 1.go函数不支持重载
// 2.go函数本身也是一种数据类型，可以赋值给变量，通过变量可以调用函数
// 3.类型为函数的变量可以作为参数, 也可以作为返回值
// 4.支持对返回值命名
// 5.go支持可变数目的参数，可变参数要放在形参列表的末尾
// func myfunc(num int, args ...type) { 
// 	for _, arg := range args {
//         fmt.Println(arg)
//     }
// }

var GAfunc = func(n1 int, n2 int)int{ // 全局匿名函数
	return n1 * n2
}