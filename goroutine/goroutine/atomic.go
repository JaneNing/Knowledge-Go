package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//func main() {
//	atomicStore()
//}

func lock() {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.Unlock()
}

func atomicStore() {
	var num int64 = 1
	atomic.StoreInt64(&num, 5)
	res := atomic.LoadInt64(&num)
	fmt.Println(res)
}

func atomicAdd() {
	var num int64 = 1
	res := atomic.AddInt64(&num, 2)
	fmt.Println(res)
	fmt.Println(num)
}
