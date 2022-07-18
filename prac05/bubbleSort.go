package main

import(
	"fmt"
	"time"
	"math/rand"
)

func main(){
	n := 20
	intSli := make([]int, n, n)
	rand.Seed(time.Now().Unix())
	for index, _ := range intSli{
		intSli[index] = rand.Intn(1000)
	}
	fmt.Printf("%v\n", intSli)
	bubbleSort(intSli)
	fmt.Printf("%v\n", intSli)
}

func bubbleSort(sli []int){
	var tmp int
	for i := len(sli) - 1; i > 0; i--{
		for j := 0; j < i; j++{
			if sli[j] > sli[j+1]{
				tmp = sli[j]
				sli[j] = sli[j+1]
				sli[j+1] = tmp
			}
		}
	}
}