package main

import(
	"fmt"
	"go_practice/prac07/testPackage"
)

func main(){
	stu1Ptr := testPackage.NewStudent("john", 20)
	stu1Ptr.SetStudentInfo("peter", 25)
	stu1Name, stu1Age := stu1Ptr.GetStudentInfo()
	fmt.Printf("name=%v, age=%v\n", stu1Name, stu1Age)
}