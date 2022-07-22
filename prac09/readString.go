package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main(){
	file, err1 := os.Open("./abc")
	if err1 != nil{
		fmt.Println("error:...", err1)
		return 
	}
	defer file.Close()

	fmt.Printf("%v\n", file)

	reader := bufio.NewReader(file) // 默认缓冲区为4096个字节
	for {
		str, err3 := reader.ReadString('\n') // 读到一个换行符结束
		fmt.Printf(":%v", str)
		if err3 == io.EOF{
			fmt.Println("\nerror...", err3)
			break
		}
	}
	
}