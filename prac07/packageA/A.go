package packageA

import(
	"fmt"
)

type A struct{
	Mem1 int
	men2 int
}

func (a *A)sayOK(){
	fmt.Println("OK",(*a))
}