package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func counter(wg *sync.WaitGroup) {
	wg.Done()
	slice := []int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func main() {
	runtime.GOMAXPROCS(5)
	traceProfile := "traceprofile"
	flag.Parse()

	f, err := os.Create(traceProfile)
	if err != nil {
		log.Fatal(err)
	}
	trace.Start(f)
	defer f.Close()
	defer trace.Stop()

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go counter(&wg)
	}
	wg.Wait()
}
