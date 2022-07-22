package main

import(
	"fmt"
	"flag"
)

func main(){
	// 定义几个变量用于接收命令行的参数
	var user string
	var passwd string
	var host string
	var port int

	// u 就是 -u 参数
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&passwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// 重要，必须调用这个方法
	// 从os.Args[1:]中解析注册的flag。
	// 必须在所有flag都注册好而未访问其值时执行。
	// 未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	fmt.Printf("%v\t%v\t%v\t%v\n", user, passwd, host, port)

}