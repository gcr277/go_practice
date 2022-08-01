package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"

)

func main(){
	//1.连接到redis
	conn, err1 := redis.Dial("tcp", "127.0.0.1:6379")
	if err1 != nil{
		fmt.Printf("%v\n", err1)
		return
	}
	defer conn.Close()
	fmt.Printf("connection success! %v\n", conn)

	//2.向redis写入数据
	res2, err2 := conn.Do("SET","name","北京的jerry")
	if err2 != nil{
		fmt.Printf("\"set\" error...%v\n", err2)
	}else{
		fmt.Printf("\"set\" %v %T\n", res2.(string),res2)
	}

	//3.从redis读取数据
	res3, err3 := redis.String(conn.Do("GET","name")) 
		// 返回的是接口,推荐使用redis自带的方法完成类型转换
	if err3 != nil{
		fmt.Printf("\"get\" error...%v\n", err3)
	}else{
		fmt.Printf("\"get\" result:%v\n", res3)
	}
}