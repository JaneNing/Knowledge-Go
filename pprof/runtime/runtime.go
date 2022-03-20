package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func counter() {
	//优化前
	//slice := make([]int, 0)
	//c := 1
	//for i := 0; i < 100000; i++ {
	//	c = i + 1 + 2 + 3 + 4 + 5
	//	slice = append(slice, c)
	//}
	//优化后
	//slice := [100000]int{0}
	//c := 1
	//for i := 0; i < 100000; i++ {
	//	c = i + 1 + 2 + 3 + 4 + 5
	//	slice[i] = c
	//}
}

func workOnce(wg *sync.WaitGroup) {
	counter()
	wg.Done()
}

func main() {
	cpuProfile := "cpuprofile"
	memProfile := "memprofile"
	flag.Parse()
	//采样cpu运行状态
	f, err := os.Create(cpuProfile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go workOnce(&wg)
	}

	wg.Wait()
	//采样memory状态
	create, err2 := os.Create(memProfile)
	if err2 != nil {
		log.Fatal(err2)
	}
	pprof.WriteHeapProfile(create)
	create.Close()
}