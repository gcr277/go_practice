package cal

import (
	_ "fmt"
	"testing" // 引入go的test框架
)

/* go test 自动测试，不带参数 -v 只输出错误信息
 * 1.testing框架将xxx_test.go文件引入
 * 2.自动调用文件中符合格式的TestXxx()函数
 * 3.测试单个文件，一定要带上被测函数源文件 go test cal_test.go cal.go
 * 4.测试单个方法，go test -test.run TestAddUpper
 */

/* 命名规范
 * func TestXxx(*testing.T)
 * 其中 Xxx 可以是任何字母数字字符串（但第一个字母不能是 [a-z]），用于识别测试例程。
 */
func TestAddUpper(t *testing.T){
	res := addUpper(10)
	if res != 55{
		t.Fatalf("addUpper(10)执行错误,期望值%v,实际值%v", 55, res)
	}
	t.Logf("addUpper(10)执行正确")
}