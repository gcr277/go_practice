package main
// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)

func main(){
	//int类型可以直接定义int8,int16,int32,int64以及uint8,uint16,uint32,uint64
	//rune类似于int32, 有符号，表示unicode
	//byte类似于int8，无符号，表示单个字符

	//查看某个变量的数据类型, 占用的字节数
	var n1 = 100
	fmt.Printf("n1的类型：%T\n", n1)
	fmt.Printf("n1占用的大小：%d\n", unsafe.Sizeof(n1))

}