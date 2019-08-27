package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type tools struct {
	err  chan error
	done chan struct{}
	once sync.Once
	wg   sync.WaitGroup
}

func newTools() *tools {
	return &tools{
		err:  make(chan error),
		done: make(chan struct{}),
		once: sync.Once{},
		wg:   sync.WaitGroup{},
	}
}

func main() {
	t := newTools()
	if err := t.do(); err != nil {
		log.Fatal(err)
	}
}

func (t *tools) do() error {
	for i := 0; i < 10; i++ {
		t.wg.Add(1)
		t.once.Do(func() {
			go func() {
				rand.Seed(time.Now().UnixNano())
				t.wg.Wait()
				t.done <- struct{}{}
			}()
		})
		go func(index int) {
			defer t.wg.Done()
			if err := multiplyByRand(index); err != nil {
				t.err <- err
			}
		}(i)
	}
	select {
	case err := <-t.err:
		return err
	case <-t.done:
		return nil
	}
}

func multiplyByRand(index int) error {
	integer := rand.Intn(10)
	fmt.Printf("%v multiplied by %v is %v \n", index, integer, index*integer)
	return nil
}
