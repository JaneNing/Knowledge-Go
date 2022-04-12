package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	task()
}

type Pool struct {
	wg   sync.WaitGroup
	work chan Task
}

type Task1 struct {
}

func (t Task1) task() {
	fmt.Println("do task")
	time.Sleep(1 * time.Second)
	fmt.Println("do task finish")
}

func main() {
	pool := New(2)
	for i := 0; i < 5; i++ {
		task1 := Task1{}
		pool.putTask(task1)
	}
	pool.shutDown()
	fmt.Println("done")
}

func New(size int) *Pool {
	p := &Pool{
		wg:   sync.WaitGroup{},
		work: make(chan Task),
	}
	p.wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			for task := range p.work {
				task.task()
			}
			p.wg.Done()
		}()
	}

	return p
}

func (p *Pool) putTask(task Task) {
	p.work <- task
}

func (p *Pool) shutDown() {
	close(p.work)
	p.wg.Wait()
}
