package main

import(
	"fmt"
)

type USB interface{ // 接口内只有未实现的方法，没有变量
	Start()
	Stop()
}

type Charger interface{
	USB
	Charge()
}
/*------------------------------*/
type Phone struct{
	IMEI string
}

// 实现了USB接口声明的所有方法，即实现了USB接口
func (p *Phone)Start(){
	fmt.Println("phone start...")
}
func (p *Phone)Stop(){
	fmt.Println("phone stop...")
}

type Camera struct{
	Model string
}

func (c Camera)Start(){
	fmt.Println("camera start...")
}
func (c Camera)Stop(){
	fmt.Println("camera stop...")
}
func (c Camera)Charge(){
	fmt.Println("camera charging...")
}

func ComputerWorking(usb USB){ // 接收USB接口类型的变量, 实际传入实现了USB接口的变量的引用
	usb.Start()
	usb.Stop()
}// usb可以接受多种实现USB接口的类型体现出多态性



func main(){
	phone1 := Phone{}
	camera1 := Camera{Model : "SONY"}

	ComputerWorking(&phone1)
	ComputerWorking(camera1)
	ComputerWorking(&camera1)
	fmt.Println("////////////////////////////////")
	var u1 USB = &phone1  // 由指针类型实现，则只能赋相应的指针
	u1.Start()
	u1.Stop()

	var u2 Charger
	fmt.Printf("%v\n", u2)
	u2 = camera1
	fmt.Printf("%v\n", u2)
	u2.Charge()
}

// 实现接口的不一定是结构体，也可以是自定义类型
// 接口可以继承，如果要实现子接口，必须实现父接口
// interface是引用类型, 默认值是nil
// 任何类型都实现了空接口 interface{}, 因此可以把任何变量赋给空接口