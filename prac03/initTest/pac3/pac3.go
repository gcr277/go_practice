package pac3
import (
	"fmt"
	"test/pac1"
	"test/info"
)
var NUM3 int

func init (){
	NUM3 = 30
	//pac1.NUM1 = 3000
	fmt.Printf("[%v]NUM1=%v|NUM3=%v\n", info.CurrFuncName(),pac1.NUM1,NUM3)
	
}