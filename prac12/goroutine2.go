package main

import (
	"fmt"
	"time"
	"sync"
)

var lock sync.Mutex

func add(n uint64, m map[uint64]uint64){
	var res uint64 = 0
	var i uint64
	for i = 1; i <= n; i++{
		res += i
	}
	lock.Lock()
	m[n] = res
	lock.Unlock()
}

func main(){
	var m1 map[uint64]uint64 = make(map[uint64]uint64)
	for i := 1; i < 100; i++{
		go add(uint64(i), m1)
	}

	time.Sleep(time.Second)

	fmt.Printf("%v\n", m1)

}