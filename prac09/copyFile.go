package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func CopyFile(dstPath string, srcPath string)(written int64, err error){
	srcFile, openErr := os.Open(srcPath)
	if openErr != nil{
		fmt.Printf("src file open error---%v\n", openErr)
		return 0, openErr
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, openErr := os.OpenFile(dstPath, os.O_WRONLY | os.O_CREATE, 0664)
	if openErr != nil{
		fmt.Printf("dst file open error---%v\n", openErr)
		return 0, openErr
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)

}

func main(){
	src := "./pictures/screenshot1.png"
	//src := "./input"
	dst := "./pictures/newPicture.jpg"
	//dst := "./output"
	_, err := CopyFile(dst, src)
	if err == nil{
		fmt.Println("done")
	}else{
		fmt.Printf("copy error---%v\n", err)
	}
		
}