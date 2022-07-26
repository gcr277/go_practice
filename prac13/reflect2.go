package main

import (
	"fmt"
	"reflect"
)
/* 
 * 通过反射修改变量的值
 */
func reflectTest03(b interface{}){
	rValPtr := reflect.ValueOf(b) // 得到的是指针
	// Elem()返回b持有的接口保管的值的Value封装，或者b持有的指针指向的值的Value封装。
	// 如果b的Kind不是Interface或Ptr会panic；如果b持有的值为nil，会返回Value零值。
	rVal := rValPtr.Elem()
	rVal.SetInt(200)
	fmt.Printf("rVal=%v, (%T)\n", rVal, rVal)
}
/*
 * 反射修改结构体的字段值
 */
func reflectTest04(b interface{}){
	rValPtr := reflect.ValueOf(b)
	fmt.Printf("rVal=%v, (%T)\n", rValPtr, rValPtr)

	rVal := rValPtr.Elem()
	fmt.Printf("rVal=%v, (%T)\n", rVal, rVal)

	fieldNum := rVal.NumField()
	fmt.Printf("fieldNum=%v, (%T)\n", fieldNum, fieldNum)

	for i := 0; i < fieldNum; i++ {
		rField := rVal.Field(i)
		fmt.Printf("rField=%v, (%T) | ", rField, rField)
	}
	fmt.Printf("\n")
	rVal.Field(1).SetInt(10)
	
	//fmt.Println(rVal.NumMethod()) // 只拿到首字母大写的方法个数
	rMethod := rVal.MethodByName("GetSum")
	fmt.Printf("rMethod=%v, (%T)\n", rMethod, rMethod)

	var argIn interface{} = 22
	getSumResRValSli := rMethod.Call([]reflect.Value{reflect.ValueOf(argIn)})  // 参数和返回值放在了[]Value里
	getSumRes := getSumResRValSli[0].Int()

	vField := rVal.FieldByName("Name") // 无法导出首字母小写的字段
	//vField := rVal.Field(0)
	//fmt.Printf("vField=%v, (%T)\n", vField, vField)

	name := vField.Interface().(string)
	fmt.Printf("%v has finished the calculation, getSumRes=%v, (%T)\n", name, getSumRes, getSumRes)

}
type cal struct{ 
	Name string `json:"name"`
	Num int		`json:"num"`
}
func (this cal)GetSum(n int)int{
	return this.Num + n
}
func main(){
	var num1 int = 100
	fmt.Println("-----------------------------")
	reflectTest03(&num1)
	fmt.Printf("main: num1:%v\n", num1)
	fmt.Println("-----------------------------")
	var cal1 cal = cal{"Jerry", 0}
	reflectTest04(&cal1)
	fmt.Printf("main: cal1:%v\n", cal1)
	fmt.Println("-----------------------------")
}