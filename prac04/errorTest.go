package main

import(
	"fmt"
	"errors"
)

func main(){
	test2()
	fmt.Println("next....")
}

func test(){
	defer func(){
		if err := recover();err != nil{
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1/num2
	fmt.Println(res)
	fmt.Println("rest of test...")
}

func test2(){
	defer func(){
		if err1 := recover();err1 != nil{ // 捕获panic抛出的error，恢复运行
			fmt.Println("recover",err1)
		}
	}()
	err := readConf("config2.ini")
	if err != nil {
		panic(err) // 抛出一个error，终止运行
	}
	fmt.Println("rest of test2...")
}

func readConf(name string) error{
	if name == "config.ini" {
		fmt.Println("read...")
		return nil
	} else {
		return errors.New("fail read...error")
	}
}