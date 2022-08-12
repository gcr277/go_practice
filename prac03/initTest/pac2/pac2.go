package pac2
import (
	"fmt"
	"test/pac1"
	"test/info"
)
var NUM2 int

func init (){
	NUM2 = 20
	pac1.NUM1 = 2000
	fmt.Printf("[%v]NUM1=%v|NUM2=%v\n", info.CurrFuncName(),pac1.NUM1,NUM2)
}