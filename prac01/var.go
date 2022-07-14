package main
import "fmt"

func main() {
	var i int			//1.指定变量类型
	num := 11.11		//2.自动判断变量类型;省略var，:= 左侧的变量应是未声明过的，否则编译错误
	var s1 string 		//3.整数默认0,字符串默认空；string是基本数据类型

	fmt.Println("i = ", i)
	fmt.Println("num = ", num)
	fmt.Println("s1 = ", s1)

	var n1, n2, n3 int //一次声明多个变量
	fmt.Println("n1=",n1," n2=",n2," n3=",n3)
	var n4, n5, name1 = 4, 5, "tom" //一次赋值多个变量
	fmt.Println("n4=", n4, " n5=", n5, " name1=",name1)
	n7, n8 ,name2 := 7, 8, "jerry" //同样可以多个变量类型自动推导
	fmt.Println("n7=", n7, " n8=", n8, " name2="+name2)  //加号可以做字符串拼接
	
	//go的变量类型转换必须是显式的，大转小不会报错，但可能溢出
	m1 := 10
	m2 := 20
	var  m3 int
	m3 = m1
	m3 = m2
	fmt.Println(m3)
	
}
