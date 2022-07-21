package main

import(
	"fmt"
	"go_practice/prac07/factory"
)

func main(){
	stu1Ptr := factory.NewStudent("john", 20)
	stu1Ptr.SetName("peter")
	stu1Ptr.SetAge(25)

	fmt.Printf("name=%v, age=%v\n", stu1Ptr.GetName(), stu1Ptr.GetAge())
}