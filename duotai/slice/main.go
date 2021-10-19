package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	var count int
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mx.Lock()
			count++
			mx.Unlock()
		}()
		//wg.Done()
	}
	wg.Wait()
	fmt.Print(count)
}
