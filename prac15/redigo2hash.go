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
	fmt.Printf("connection success! \n")

	//2.向redis写入数据
	res2, err2 := conn.Do("HSET","hashmap1","007","Bond")
	if err2 != nil{
		fmt.Printf("\"hset\" error...%v\n", err2)
	}else{
		fmt.Printf("\"hset\" %v %T\n", res2,res2)
	}
	res2, err2 = conn.Do("HSET","hashmap1","009","大内密探")
	if err2 != nil{
		fmt.Printf("\"hset\" error...%v\n", err2)
	}else{
		fmt.Printf("\"hset\" %v %T\n", res2,res2)
	}

	//3.从redis读取数据
	res3, err3 := redis.Strings(conn.Do("HMGET","hashmap1","007","009")) 
		// 返回的是接口,推荐使用redis自带的方法完成类型转换
	if err3 != nil{
		fmt.Printf("\"hmget\" error...%v\n", err3)
	}else{
		fmt.Printf("\"hmget\" result:%v %T\n", res3,res3)
	}
}