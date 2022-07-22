package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	filePath := "./abc"
	byteSli, err := ioutil.ReadFile(filePath) // 一次性读取整个文件，内含打开和关闭操作,适合小文件
	if err != nil{							  // v1.16 改为直接调用os.ReadFile(filePath)
		fmt.Println("error---", err)
	}
	for _, value := range byteSli{
		fmt.Printf("%c", value)
	}
	fmt.Printf("\n")
}