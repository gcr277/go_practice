package packageC

import(
	"fmt"
)

type C struct{
	Mem4 int
	men5 int
}

func (c *C)sayOK(){
	fmt.Println("OK",(*c))
}