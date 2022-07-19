package main

import (
	"fmt"
)

func main(){
	var heroSli []map[string]string = make([]map[string]string, 2)
	fmt.Println(cap(heroSli))
	
	if heroSli[0] == nil{
		heroSli[0] = make(map[string]string, 2)
		heroSli[0]["name"] = "jinx"
		heroSli[0]["job"] = "ADC"
	}
	if heroSli[1] == nil{
		heroSli[1] = make(map[string]string, 2)
		heroSli[1]["name"] = "zed"
		heroSli[1]["job"] = "MID"
	}

	fmt.Printf("%v\n", heroSli)

	newHero := make(map[string]string, 2)
	newHero["name"] = "lulu"
	newHero["job"] = "SUP"
	heroSli = append(heroSli, newHero)
	// 千万注意不要越出slice的cap，所以使用append动态地向slice里增加map
	fmt.Printf("%v\n", heroSli)
}