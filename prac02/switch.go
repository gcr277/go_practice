package main

import "fmt"

func main() {
	var a int64 = 3
	switch a {
	case 1:
		fmt.Println("1") // 匹配项后面无需加break, 匹配执行后自动退出switch
	case 2, 3:
		fmt.Println("2")
	case 4:
		fmt.Println("3")
	default:
		fmt.Println("not match")
	}

	switch {
	case a == 1:
		fmt.Println("1")
	case a == 2, a == 3:
		fmt.Println("2")
		fallthrough //默认只能穿透一层
	case a == 4:
		fmt.Println("3")
	default:
		fmt.Println("not match")
	}
}

// 1.case后的表达式必须和switch的表达式类型一致
// 2.case后可接多个表达式，用逗号分隔
// 3.case后的表达式如果是常量则不能重复
// 4.default语句不是必需的
// 5.switch后面可以不带表达式类似if-else分支来使用
// 6.switch后面可以接赋值表达式再加分号，然后按照5的方式使用（不推荐）
// 7.如果在case语句块后面增加fallthrough，则会继续执行下面的case（switch穿透）
// 8.switch还可用于type-switch来判断某个interface变量中实际指向的变量类型
