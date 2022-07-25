package main

import(
	"fmt"
	"encoding/json"
)

type student struct{
	Name string 	`json:"name"`  // 为了让大写的字段序列化后得到小写的字段，可以使用tag，原理是反射
	Age int			`json:"age"`
	Score float64	`json:"score"`
}

func StructSerialize(studentPtr *student)([]byte, error){
	
	data,err := json.Marshal(studentPtr)
	if err != nil{
		return nil, err
	}
	return data, err
}
func StructUnserialize(data []byte)(*student, error){
	var student2 student
	err := json.Unmarshal(data, &student2)
	if err != nil{
		return nil, err
	}
	return &student2, err
}

func MapSerialize(m map[string]interface{})([]byte, error){
	data, err := json.Marshal(m)
	if err != nil{
		return nil, err
	}
	return data, err
}
func MapUnserialize(data []byte)(map[string]interface{}, error){
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)  	// 反序列化无需make，make已经被封装到Unmarshal内
	if err != nil{						// 这也解释了为什么要传入&m而不是m：
		return nil, err					// m指向未分配的空间，无法存放反序列化的内容
	}									// 因此必须传入&m使m指向Unmarshal内make的空间
	return m, err
}

func main(){
	student1 := student{
		Name : "tom",
		Age : 18,
		Score : 90.0,
	}

	serialized, err := StructSerialize(&student1)
	if err != nil{
		fmt.Printf("StructSerializeFail----%v\n", err)
	}
	fmt.Printf("Serialize:%v\n", string(serialized))

	stuPtr, err := StructUnserialize(serialized)
	if err != nil{
		fmt.Printf("StructUnserializeFail----%v\n", err)
	}
	fmt.Printf("Unserialize:%v\n", *stuPtr)
	fmt.Println("--------------------------------------------")
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Bob"
	m1["age"] = 30
	m1["address"] = "New York"

	serialized2, err2 := MapSerialize(m1)
	if err2 != nil{
		fmt.Printf("MapSerializeFail----%v\n", err2)
	}
	fmt.Printf("Serialize:%v\n", string(serialized2))

	m2, err2 := MapUnserialize(serialized2)
	if err2 != nil{
		fmt.Printf("MapUnserializeFail----%v\n", err)
	}
	fmt.Printf("Unserialize:%v\n", m2)


}