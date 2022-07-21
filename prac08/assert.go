package main

import(
	"fmt"
)

type USB interface{ // 接口内只有未实现的方法，没有变量
	Start()
	Stop()
}

/*------------------------------*/
type Phone struct{
	IMEI string
}
func (p *Phone)Start(){
	fmt.Println("phone start...")
}
func (p *Phone)Stop(){
	fmt.Println("phone stop...")
}
func (p *Phone)Call(){
	fmt.Println("phone call...")
}


func main(){
	var usb USB
	var phone1 Phone = Phone{
		IMEI : "88888888",
	}
	usb = &phone1 // 由于*Phone 实现了USB接口
	var phone2Ptr *Phone
	var ok bool
	phone2Ptr, ok = usb.(*Phone) // 将usb断言为*Phone类型，成功返回true

	fmt.Printf("%v %T\n", phone2Ptr, phone2Ptr)
	if ok{
		(*phone2Ptr).Call()
	}

}