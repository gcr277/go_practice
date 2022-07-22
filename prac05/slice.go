package main

import "fmt"

func main(){
	// 
	var intArr [5]int = [...]int{1,2,3,4,5}
	var sli []int = intArr[1:3]
	sli2 := sli
	// 定义一个切片，范围是intArr下标1～3但包含3，切片使用同样不可以越界

	fmt.Printf("%p\n", &intArr[1])
	fmt.Printf("%p\n", &sli[0])		
	fmt.Printf("%p\n", &sli2[0])	
	// 切片是数组片段的引用

	fmt.Println()

	fmt.Printf("%p\n", &intArr) 	// 数组的地址
	fmt.Printf("%p\n", &sli)		// 切片这个结构体的地址
	fmt.Printf("%p\n", &sli2)		// 切片这个结构体的地址
	fmt.Printf("%p\n", sli)			// 切片指向的首个元素地址
	fmt.Printf("%p\n", sli2)		// 切片指向的首个元素地址
	fmt.Println()
	
	// slice的内存布局: first element pointer | len | cap
	
	// slice的三种定义方式
	// 1.对已有数组或切片进行引用
	var sli3 []int = intArr[:]
	fmt.Printf("%v\n", sli3)
	// 2.使用内置make函数，容量可选
	var sli4 []int = make([]int, 2, 4)
	fmt.Printf("%v\n", sli4)
	// 3.定义切片时直接指定元素，本质上指定了一个初始化了这些元素的数组
	var sli5 []int = []int{10, 20, 30, 40, 50}
	fmt.Printf("%v, %v, %v\n", sli5, len(sli5), cap(sli5))
	fmt.Println()
	
	/* func append(slice []Type, elems ...Type) []Type
	内建函数append将元素追加到切片的末尾。
	若它有足够的容量，其目标就会重新切片以容纳新的元素;
	否则，就会分配一个新的基本数组。
	append返回更新后的切片，因此必须存储追加后的结果。*/
	sli5 = append(sli5, 60, 70)
	fmt.Printf("%v, %v, %v\n", sli5, len(sli5), cap(sli5))
	sli5 = append(sli5, sli3...)
	fmt.Printf("%v, %v, %v\n", sli5, len(sli5), cap(sli5))
	fmt.Println()

	/*func copy(dst, src []Type) int
	内建函数copy将元素从来源切片复制到目标切片中，
	也能将字节从字符串复制到字节切片中。(来源和目的是独立的互不影响)
	copy返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个。(截断)
	来源和目标的底层内存可以重叠。*/
	var sli6 []int = []int{106, 107}
	var sli7 []int = []int{60, 70, 80, 90, 100}
	fmt.Printf("%p\n", sli6)			
	fmt.Printf("%p\n", sli7)
	num := copy(sli6, sli7)
	fmt.Printf("%v, %v, %v\n", sli6, len(sli6), cap(sli6))
	fmt.Printf("copy = %v\n", num)
	fmt.Printf("%p\n", sli6)			
	fmt.Printf("%v\n", sli7)
	fmt.Println()

	/* string本质也是切片，但不可变，若需改变可以转换为[]byte或[]rune，
	修改后再转回string */
	str1 := "hello北京"
	sli8 := []rune(str1)
	sli8[5] = '南'
	fmt.Printf("%c, %v, %v\n", sli8, len(sli8), cap(sli8))
	str1 = string(sli8)
	fmt.Println(str1)
}