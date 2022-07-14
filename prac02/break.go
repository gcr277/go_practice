package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var tmp int
	var sum int = 0
	rand.Seed(time.Now().Unix()) // time.Now().Unix() 返回一个从1970年开始的秒数
	//fmt.Println(time.Now().Unix())
	label1:
	for {
		//fmt.Println("for 0")
		// label2:
		for {
			fmt.Println("for 1")
			if tmp == 99 {
				fmt.Println("tmp =", tmp)
				fmt.Println("次数 = ", sum)
				break label1  // break可以指定跳到哪层循环，continue也可以跳转标签
			}
			tmp = rand.Intn(100) + 1
			sum++
			// if 1==0 {
			// 	break label2
			// }
		}
	
	}

}
