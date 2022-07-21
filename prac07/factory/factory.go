package factory

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

func (studentP *student)SetName(name string){
	(*studentP).name = name
	return
}

func (studentP *student)SetAge(age int){
	(*studentP).age = age
	return
}
func (studentP *student)GetName() string{
	return (*studentP).name
}
func (studentP *student)GetAge() int{
	return (*studentP).age
}