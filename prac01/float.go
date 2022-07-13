package main
import "fmt"

func main(){
	//单精度float32, 双精度float64，不指定类型默认为float64
	var num1 = 1.1
	fmt.Printf("%T\n", num1)
	var n1 = 10.1 + 'a'
	var n2 = 'a' + 10
	fmt.Println(n1)
	fmt.Println(n2)
	//浮点数默认值为0
	var b1 bool //默认为false, bool变量不能用0或1代替
	fmt.Printf("%T\n", b1)
	fmt.Printf("b1原值%v\n",b1)


}