package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	len := 10

	res := FastCalculate(len)
	fmt.Println(res)
}

func FastCalculate(len int) int {
	nums := make(chan int, len)

	group := sync.WaitGroup{}
	group.Add(len)

	for i := 0; i < len; i++ {
		go func() {
			time.Sleep(2 * time.Second)
			nums <- rand.Int() % 100
			group.Done()
		}()
	}

	go func() {
		group.Wait()
		close(nums)
	}()

	res := 0
	for num := range nums {
		res += num
	}

	return res
}
