package main

import(
	"fmt"
	"encoding/json"
)



// 在js语言中，一切都是对象。因此，任何的数据类型都可以通过JSON表示，
// 例如字符串、数字、对象、数组、map、结构体等

// JSON键值对是用来保存数据的一种方式
// {"name":"tom", "age":18, "hobby":["music","chess"]}

type Student struct{
	Name string 	`json:"name"`  // 为了让大写的字段序列化后得到小写的字段，可以使用tag，原理是反射
	Age int			`json:"age"`
	Score float64	`json:"score"`
}

func StructSerialize(){
	student1 := Student{
		Name : "tom",
		Age : 18,
		Score : 90.0,
	}
	data,err := json.Marshal(&student1)
	if err != nil{
		fmt.Printf("StructSerializeFail----%v\n", err)
	}
	fmt.Printf("serialize:%v\n", string(data))
}

func MapSerialize(){
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Bob"
	m1["age"] = 30
	m1["address"] = "New York"
	data, err := json.Marshal(m1)
	if err != nil{
		fmt.Printf("MapSerializeFail----%v\n", err)
	}
	fmt.Printf("serialize:%v\n", string(data))
}

func SliceSerialize(){
	var s1 []map[string]interface{}
	s1 = make([]map[string]interface{}, 3)
	s1[0] = map[string]interface{}{
		"name" : "Tom",
		"age" : 10,
		"hobby" : []string{"music", "football"},
	}
	s1[1] = map[string]interface{}{
		"name" : "Jerry",
		"age" : 20,
		"hobby" : []string{"painting", "movie", "shopping"},
	}
	s1[2] = map[string]interface{}{
		"name" : "Carvin",
		"age" : 30,
		"hobby" : []string{"coding"},
	}
	data, err := json.Marshal(s1)
	if err != nil{
		fmt.Printf("SliceSerializeFail----%v\n", err)
	}
	fmt.Printf("serialize:%v\n", string(data))
}

func main(){
	StructSerialize()
	MapSerialize()
	SliceSerialize()
}