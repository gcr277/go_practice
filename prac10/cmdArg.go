package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Printf("一共有%v个参数:\n", len(os.Args))
	for i, v := range os.Args{
		fmt.Printf("Args[%v]\t%v\n", i, v)
	}
}