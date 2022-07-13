package main
import "fmt"

func main(){
	//支持判断前赋值的语法
	if a := 10.0; a > 9 { // go的判断不需要写括号
		fmt.Println("pass")
	}
}