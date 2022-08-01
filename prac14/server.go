package main

import (
	"fmt"
	"net"
	"io"
)

func process(conn net.Conn){
	defer conn.Close()
	for{
		buf := make([]byte, 1024)
		//fmt.Printf("conn with(%v): [a new buffer is waiting message fron client...]\n",conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil{
			if err == io.EOF{
				fmt.Printf("conn with(%v): [client has left... process return]\n",conn.RemoteAddr().String())
			}else{
				fmt.Printf("conn with(%v): [read error ..%v]\n",conn.RemoteAddr().String(),err)
			}
			return
		}
		fmt.Printf("conn with(%v): %v\n", conn.RemoteAddr().String(), string(buf[0:n]))
	}
}

func main(){
	listen, err1 := net.Listen("tcp", "127.0.0.1:8888") // 开启一个端口
	if err1 != nil{
		panic(err1)
	}
	fmt.Println("main: [server start listening...]")
	
	for{
		fmt.Println("main: [server waiting for conn...]")
		conn, err2 := listen.Accept() // 等待客户端连接
		if err2 != nil{
			fmt.Printf("%v\n", err2)
		}else{
			fmt.Printf("main: [conn client ip=%v]\n", conn.RemoteAddr().String())
			go process(conn)
		}
		
	}
	
}