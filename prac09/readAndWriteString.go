package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main(){

	filePath := "./ddd"
	file, err := os.OpenFile(filePath, 
		os.O_CREATE | os.O_RDWR | os.O_APPEND,
		0666)
	if err != nil {
		fmt.Printf("error---", err)
		return 
	}
	defer file.Close()

	reader := bufio.NewReader(file) // 默认缓冲区为4096个字节
	for {
		str, err3 := reader.ReadString('\n') // 读到一个换行符结束
		fmt.Printf(":%v", str)
		if err3 == io.EOF{
			fmt.Printf("\n")
			break
		}
	}

	writer := bufio.NewWriter(file)
	for i:=0; i<5; i++{
		writer.WriteString("\r\nthis is a new line!!!")
	}
	writer.Flush() // 将缓存中的内容写入硬盘

}

	//文件打开模式
	/*const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	)*/