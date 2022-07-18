package main

import(
	"fmt"
	"time"
	"math/rand"
)

func main(){
	n := 11
	intSli := make([]int, n, n)
	rand.Seed(time.Now().Unix())
	for index, _ := range intSli{
		intSli[index] = rand.Intn(1000)
	}
	bubbleSort(intSli)
	fmt.Printf("%v\n", intSli)
	var dstNum int
	var indexFind int = -1
	for {
		fmt.Scanf("%d", &dstNum)
		indexFind = binaryFind(intSli, dstNum, 0, len(intSli)-1)
		if indexFind != -1 {
			fmt.Printf("查找目标的下标是:%d\n", indexFind)
		}else {
			fmt.Println("没有找到")
		}
	}
}

func binaryFind(sli []int, dstNum int, leftIndex int, rightIndex int) int{
	midIndex := (leftIndex + rightIndex) / 2
	if leftIndex > rightIndex {
		return -1
	}

	if sli[midIndex] < dstNum{
		leftIndex = midIndex + 1
		return binaryFind(sli, dstNum, leftIndex, rightIndex)
	}else if sli[midIndex] > dstNum{
		rightIndex = midIndex - 1
		return binaryFind(sli, dstNum, leftIndex, rightIndex)
	}else {
		return midIndex
	}
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