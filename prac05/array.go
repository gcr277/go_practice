package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// 数组是值类型
	var array1 [5]int64 = [5]int64{1,2,3,4,5}
	var array2 [5]int64
	for i := 0; i < len(array1);i++ {
		fmt.Printf("%p  ", &array1[i])
	}
	fmt.Println()
	
	for i := 0; i < len(array2) ;i++ {
		fmt.Printf("%x  ", &array2[i])
	}
	fmt.Println()
	fmt.Printf("%p %x\n", &array1, &array2)
	//fmt.Printf("%s %s\n", array1, array2)

	//数组的初始化方式
	var numArray1 [5]int = [5]int{1,2,3,4,5}
	var numArray2 = [5]int{1,2,3,4,5}
	numArray3 := [...]int{1,2,3,4,5}
	numArray4 := [...]int{0:1, 1:2, 2:3, 3:4, 4:5}  //指定下标
	
	fmt.Println(numArray1)
	fmt.Println(numArray2)
	fmt.Println(numArray3)
	fmt.Println(numArray4)

	arrayPtr := &numArray4 // 数组是值类型，想在其他函数修改本数组可以使用指针
	(*arrayPtr)[4] = 6
	fmt.Println(numArray4)

	//数组有越界检查，越界访问会报错
	fmt.Println("**************************************")
	rand.Seed(time.Now().Unix())
	var numArray5 [10]int
	for i := 0; i < len(numArray5); i++{
		numArray5[i] = rand.Intn(100) + 1
	}
	fmt.Println(numArray5)
	arrayReverse(&numArray5)
	fmt.Println(numArray5)

}

func arrayReverse(arrPtr *[10]int){
	var tmp int
	for i := 0; i < len(*arrPtr)/2; i++{
		tmp = (*arrPtr)[len(*arrPtr) - 1 - i]
		(*arrPtr)[len(*arrPtr) - 1 - i] = (*arrPtr)[i]
		(*arrPtr)[i] = tmp
	}
}