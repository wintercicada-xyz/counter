package counter2

import (
	"sync/atomic"
	"time"
)

var counter map[string]*uint64

func Init() {
	counter = make(map[string]*uint64)
}

func Flush2broker(resetDuration uint64, FlushCb func(m map[string]*uint64)) {
	go func() {
		for {
			time.Sleep(time.Duration(resetDuration) * time.Millisecond)
			FlushCb(counter)
			counter = make(map[string]*uint64)
		}
	}()
}

func Incr(key string, count uint64) {
	num, exists := counter[key]
	if !exists {
		counter[key] = &count
	} else {
		atomic.AddUint64(num, count)
	}
}
