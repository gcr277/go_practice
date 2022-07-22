package main

import (
	"os"
	"fmt"

)

func PathExist(path string)(bool, error){
	_, err := os.Stat(path)
	if err == nil{
		return true, nil
	}
	if os.IsNotExist(err){
		return false, nil
	}
	return false, err
}

func main(){
	inputFilePath := "./input"
	inputFileExist, _ := PathExist(inputFilePath)
	outputFilePath := "./output"  // 不存在则自动创建
	data, err := os.ReadFile(inputFilePath)
	if (!inputFileExist) || (err != nil){
		fmt.Printf("read file error---%v\n", err)
		return
	}
	err2 := os.WriteFile(outputFilePath, data, 0664)
	if err2 != nil{
		fmt.Printf("write file error---%v\n", err2)
		return
	}
}