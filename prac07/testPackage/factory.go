package testPackage

import(

)

type student struct{ // 结构体类型名小写后不能直接跨包使用，可以使用工厂模式解决
	name string		// 字段名小写后也不能直接跨包使用，可以使用get和set解决
	age int
}

func NewStudent(n string, a int) *student{
	return &student{
		name : n,
		age : a,
	}
}

func (studentP *student)SetStudentInfo(name string, age int){
	(*studentP).name = name
	(*studentP).age = age
}

func (studentP *student)GetStudentInfo() (string, int){
	return (*studentP).name, (*studentP).age
}