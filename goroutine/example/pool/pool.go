package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"
)

var (
	ErrorPoolClose = errors.New("pool have closed")
	ErrorPoolFull  = errors.New("pool full")
)

type Pool struct {
	mutex   sync.Mutex
	sources chan io.Closer
	factory func() (io.Closer, error)
	close   bool
}

type DBConnection struct {
	id int64
}

func (c DBConnection) Close() error {
	fmt.Println("close")
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	group := sync.WaitGroup{}
	group.Add(10)

	pool := NewPool(func() (io.Closer, error) {
		return &DBConnection{
			id: int64(rand.Intn(10000)),
		}, nil
	}, 10)

	for i := 0; i < 10; i++ {
		go func() {
			resource, err := pool.Acquire()
			if err != nil {
				fmt.Println(err)
				return
			}

			connection := resource.(*DBConnection)
			fmt.Println(fmt.Sprintf("%d", connection.id) + " doing thing")
			time.Sleep(2 * time.Second)
			pool.Release(resource)

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("done")
}

func (pool *Pool) Close() {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.close {
		return
	}

	pool.close = true

	close(pool.sources)

	for source := range pool.sources {
		source.Close()
	}
}

func (pool *Pool) Release(closer io.Closer) error {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.close {
		return ErrorPoolClose
	}

	select {
	case pool.sources <- closer:
	default:
		return ErrorPoolFull
	}

	return nil
}

func (pool *Pool) Acquire() (io.Closer, error) {
	select {
	case source, ok := <-pool.sources:
		if !ok {
			return nil, ErrorPoolClose
		}
		return source, nil
	default:
		return pool.factory()
	}
}

func NewPool(fn func() (io.Closer, error), size int64) *Pool {
	return &Pool{
		mutex:   sync.Mutex{},
		sources: make(chan io.Closer, size),
		factory: fn,
		close:   false,
	}
}
