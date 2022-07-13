package main
import (
	"fmt"
	"go_practice/prac01/model"
)

func main (){
	var b int = 1
	
	fmt.Println(model.HeroName, b) 
	// 如果变量名、常量名、函数名首字母大写则可以被其他包访问，否则只能被自己包访问
}