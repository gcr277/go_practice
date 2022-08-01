package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"

)

var pool *redis.Pool
func init(){
	pool = &redis.Pool{
		MaxIdle: 8,
		MaxActive: 0,
		IdleTimeout: 100,
		Dial: func()(redis.Conn, error){
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main(){
	conn := pool.Get()
	defer conn.Close()

	res, err := conn.Do("SET","name3","北京的tom")
	if err != nil{
		fmt.Printf("\"set\" error...%v\n", err)
	}else{
		fmt.Printf("\"set\" %v %T\n", res,res)
	}

	res2, err2 := redis.String(conn.Do("GET","name3")) 
	if err2 != nil{
		fmt.Printf("\"get\" error...%v\n", err2)
	}else{
		fmt.Printf("\"get\" result:%v\n", res2)
	}
	pool.Close() // 一定确保用完再关闭连接池
}