package main

import "fmt"

func main(){
	
	fmt.Println("first day peaches' num = ", getTodayPeach(1))
}

func getTodayPeach(day int) int {
	if day == 10 {
		return 1
	}else{
		return (getTodayPeach(day + 1) + 1) * 2 
	}

}