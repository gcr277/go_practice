package pac1
import (
	"fmt"
	"test/info"
)
var NUM1 int
var S1 string = "pac1S1"
func init (){
	NUM1 = 10
	fmt.Printf("[%v]NUM1=%v\n", info.CurrFuncName(),NUM1)
}
	
