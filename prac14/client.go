package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("client start dialing...")
	conn, err1 := net.Dial("tcp", "127.0.0.1:8888") // 开启一个端口
	if err1 != nil {
		panic(err1)
	}

	for {
		messageLine, err2 := bufio.NewReader(os.Stdin).ReadString('\n')
		if err2 != nil {
			fmt.Printf("[read from stdin error..%v]\n", err2)
		}
		messageLine = strings.Trim(messageLine, "\n")
		if messageLine == "exit" || messageLine == "quit"{
			break
		}
		byteNum, err3 := conn.Write([]byte(messageLine))
		if err3 != nil {
			fmt.Println("[write to connection error..%v]\n", err2)
		}
		fmt.Printf("[writen %v bytes]\n", byteNum)
	}

}
