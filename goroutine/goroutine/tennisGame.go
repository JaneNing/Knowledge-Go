package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	game()
}

func game() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	channel := make(chan int)
	go player("Jane", channel, &waitGroup)
	go player("Ning", channel, &waitGroup)

	fmt.Println("game start")
	channel <- 1

	waitGroup.Wait()
	fmt.Println("game done")
}

func player(name string, channel chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	fmt.Println(name + " ready")
	for {
		ball, ok := <-channel
		if !ok {
			fmt.Println(name + " win")
			return
		}

		if rand.Intn(100)%99 == 0 {
			fmt.Println(name + " lost ball")
			close(channel)
			return
		}

		fmt.Println(name + " hit ball back")

		channel <- ball
	}
}
