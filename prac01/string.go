package main
import (
	"fmt"
	_ "unsafe"  // _ 表示忽略
	"strconv"
)

func main(){
	//字符串一旦赋值，不能修改


	//字符串有两种表示形式
	//1.双引号，识别转义字符
	str1 := "hello\nworld"
	fmt.Println(str1)
	//2.反引号，原生表示。更安全，防止注入攻击
	str2 := `hello\nworld`
	fmt.Println(str2)


	//字符串拼接
	str3 := "hello" + "world"
	str4 := "hello" + 
			"hi" + 
			"konnnitiha" + 
			"nihao"     //如需换行，+ 必须放在每行的后面，因为go识别到 + 不会自动添加分号
	fmt.Println(str3)
	fmt.Println(str4)


	//其他基本数据类型转string
	var str5 string = ""
	var num1 int = 99
	str5 = fmt.Sprintf("%d", num1) //第一种方法，使用Sprintf函数
	fmt.Println(str5)
	var num2 float64 = 123.123
	str5 = fmt.Sprintf("%f", num2)
	fmt.Println(str5)
	var b1 bool = true
	str5 = fmt.Sprintf("%v", b1)
	fmt.Println(str5)
	var c1 byte = 'h'
	str5 = fmt.Sprintf("%c", c1)
	fmt.Println(str5)
	
	//第二种方法，使用strconv包的FormatXxx函数
	var str6 string = ""
	str6 = strconv.FormatInt(int64(num1), 10) //变量，进制
	fmt.Println(str6)
	str6 = strconv.FormatFloat(num2, 'f', 10, 64) //变量，格式，保留位数，保证多少位不溢出
	fmt.Println(str6)
	str6 = strconv.FormatBool(b1)
	fmt.Println(str6)

	//第三种方法，使用strconv包的Itoa函数
	var num3 int64 = 100
	str6 = strconv.Itoa(int(num3))
	fmt.Println(str6)


	//string转其他基本类型，使用strconv.ParseXxxx函数
	str6 = "true"
	var b2 bool
	b2, _ = strconv.ParseBool(str6, )  // _ 可以表示忽略接收的返回值
	fmt.Printf("b2 is %T , = %v\n", b2, b2)

	str6 = "200"
	var num4 int64
	num4, _ = strconv.ParseInt(str6, 10, 32)  //参数分别为待转换字符串，进制，最大限制的位数(若溢出会在err中体现)                                  
	fmt.Printf("num4 is %T , = %v\n", num4, num4)																   
	//非数字转字符串失败会返回0
}