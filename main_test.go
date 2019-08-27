package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func testmultiplyByRand(index int) error {
	integer := rand.Intn(100)
	if integer < 5 {
		return fmt.Errorf("there was an error in the index: %v integer: %v", index, integer)
	}
	fmt.Printf("%v multiplied by %v is %v \n", index, integer, index*integer)
	return nil
}

func TestDo(t *testing.T) {
	errChan := make(chan error)
	doneChan := make(chan struct{})
	once := sync.Once{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		once.Do(func() {
			go func() {
				rand.Seed(time.Now().UnixNano())
				wg.Wait()
				doneChan <- struct{}{}
			}()
		})
		go func(index int) {
			defer wg.Done()
			if err := testmultiplyByRand(index); err != nil {
				errChan <- err
			}
		}(i)
	}
	select {
	case err := <-errChan:
		t.Error(err)
	case <-doneChan:
	}
}
