package main

import (
	"fmt"
	"reflect"
)

/*
 * 	┌------- args ----> interface{} --- reflect.Value() -----┐
 * 	|					 |		↑						  	 ↓
 * var					 | 		|					 	  rValue
 *  ↑					 |		|							 |
 *  └------ assert ------┘		└---- rValue.Interface() ----┘
 *
 */

func reflectTest01(b interface{}){
	// 1.通过反射获取反射类型
	rType := reflect.TypeOf(b)
	fmt.Printf("1\trType=%v, (%T)\n", rType, rType)

	// 2.通过反射获取反射值
	rVal := reflect.ValueOf(b)
	fmt.Printf("2\trVal=%v, (%T)\n", rVal, rVal)

	// 3.1反射值获取基本值
	valueNum := rVal.Int() // 如果rVal不是int, int8等会报panic
	fmt.Printf("3.1\tvalueNum=%v, (%T)\n", valueNum, valueNum)

	// 3.2反射值获取基本类型
	interfaceVal := rVal.Interface() // 先转成接口
	fmt.Printf("3.2\tval=%v, (%T)\n", interfaceVal, interfaceVal)
	val := interfaceVal.(int) // 再对接口断言
	fmt.Printf("3.2\tval=%v, (%T)\n", val, val)

}

type student struct{
	name string
	Age int
}
func reflectTest02(b interface{}){
	// 1.通过反射获取反射类型
	rType := reflect.TypeOf(b)
	fmt.Printf("1\trType=%v, (%T)\n", rType, rType)

	// 2.通过反射获取反射值
	rVal := reflect.ValueOf(b)
	fmt.Printf("2\trVal=%v, (%T)\n", rVal, rVal)

	// 3.1反射值获取反射类别
	rKind := rVal.Kind() // 如果rVal不是int, int8等会报panic
	//rKind := rType.Kind() // 也可以用反射类型(本质是一个大接口)获取反射类别
	fmt.Printf("3.1\trKind=%v, (%T)\n", rKind, rKind)

	// 3.2反射值获取基本类型
	interfaceVal := rVal.Interface() // 先转成接口
	fmt.Printf("3.2\tval=%v, (%T)\n", interfaceVal, interfaceVal)
	val := interfaceVal.(student) // 再对接口断言
	fmt.Printf("3.2\tval=%v, (%T)\n", val, val)
	fmt.Printf("3.2\tval=%v, (%T)\n", val.name, val.name)

}
func main(){
	var num1 int = 100
	fmt.Println("-----------------------------")
	reflectTest01(num1)
	fmt.Println("-----------------------------")
	var student1 student = student{"Tom", 20}
	reflectTest02(student1)
	fmt.Println("-----------------------------")
}