package main

import (
	"counter/counter1"
	"counter/counter2"
	"sync"
	"testing"
)

const THREADS = 8
const COUNTS = 1e6

func BenchmarkCounter1(b *testing.B) {
	b.ResetTimer()
	waitGroup := &sync.WaitGroup{}
	counter1.Init()
	counter1.Flush2broker(5000, func() {})
	for i := 0; i < 8; i++ {
		waitGroup.Add(1)
		go func() {
			for i := 0; i < COUNTS; i++ {
				counter1.Incr("lala", uint64(i))
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func BenchmarkCounter2(b *testing.B) {
	b.ResetTimer()
	waitGroup := &sync.WaitGroup{}
	counter2.Init()
	counter2.Flush2broker(5000, func(m map[string]*uint64) {})
	for i := 0; i < 8; i++ {
		waitGroup.Add(1)
		go func() {
			for i := 0; i < COUNTS; i++ {
				counter2.Incr("lala", uint64(i))
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}
