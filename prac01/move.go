package main
import "fmt"

func main(){
	var a int8 = -128
	a = a >> 1
	fmt.Printf("a = %d\n", a)
}