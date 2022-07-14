package main
import "fmt"

func main(){
	var b int = 10
	if a := b; a < 9 {  // golang里条件判断语句不能有a = b这样的赋值语句
		fmt.Println("pass")
	} else{ // else必须紧跟着if结束的括号
		fmt.Println("not pass") 
	}
}