package main

import(
	"fmt"
)
/********************************************/
type people struct{
	name string
	age int
}
func (people *people)Talk(){
	fmt.Println("hello")
}
func (people *people)SetName(name string){
	(*people).name = name
	return 
}
func (people *people)SetAge(age int){
	(*people).age = age
	return
}
func (people *people)GetName()string{
	return (*people).name
}
func (people *people)GetAge()int{
	return (*people).age
}

/********************************************/
type programmer struct{
	people  // 嵌入匿名结构体以实现“继承”
	language string
}
func (programmer *programmer)Work(){
	fmt.Println("programming with", (*programmer).language)
}
func (programmer *programmer)SetLanguage(language string){
	(*programmer).language = language
}
/********************************************/
type teacher struct{
	people
	school string
}
func (teacher *teacher)Work(){
	fmt.Println("teaching in", (*teacher).school)
}
func (teacher *teacher)SetSchool(school string){
	(*teacher).school = school
}
/********************************************/
func main(){
	programmer1 := programmer{}
	(&programmer1.people).SetName("张三")
	(&programmer1.people).SetAge(24)
	(&programmer1).SetLanguage("Go")
	fmt.Println(programmer1)
	(&programmer1).Work()
	fmt.Println("---------------------")
	
	teacher1 := teacher{}
	(&teacher1.people).SetName("李四")
	(&teacher1.people).SetAge(25)
	(&teacher1).SetSchool("DLUT")
	fmt.Println(teacher1)
	(&teacher1).Work()
	fmt.Println("---------------------")

}