package main

import(
	"fmt"
	"sort"
	"math/rand"
)

type Hero struct{
	Name string
	Age int
}

type HeroSlice []Hero
func (heroSlice HeroSlice)Len()int{  // 实现sort.Interface的三个方法
	return len(heroSlice)
}
func (heroSlice HeroSlice)Less(i,j int)bool{ // i元素是否要排到j元素之前
	return heroSlice[i].Age < heroSlice[j].Age
}
func (heroSlice HeroSlice)Swap(i,j int){
	tmp := heroSlice[i]
	heroSlice[i] = heroSlice[j]
	heroSlice[j] = tmp
	return
}

func main(){
	n := 10
	var hs1 HeroSlice = make(HeroSlice, n)
	for i:=0 ; i<len(hs1) ; i++{
		hs1[i] = Hero{
			Name : fmt.Sprintf("hero%d", i),
			Age : rand.Intn(100),
		}
	}
	fmt.Printf("%v\n", hs1)
	sort.Sort(hs1)
	fmt.Printf("%v\n", hs1)
}