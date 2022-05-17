package main

import (
	"counter/counter"
	"fmt"
	"sync"
	"time"
)

func main() {
	waitGroup := &sync.WaitGroup{}
	counter.Init()
	counter.Flush2broker(5000, func() { fmt.Println("Flush") })
	for {
		for i := 1; i < 1000; i += 25 {
			waitGroup.Add(1)
			go func(index int) {
				for i := index; i < index+25; i++ {
					time.Sleep(1 * time.Millisecond)
					counter.Incr("lala", uint64(i))
				}
				waitGroup.Done()
			}(i)
		}
		waitGroup.Wait()
	}
}
