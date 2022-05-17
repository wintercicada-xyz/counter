package counter

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var syncMap sync.Map

func Init() {
}

func Flush2broker(resetDuration uint64, FlushCb func()) {
	go func() {
		for {
			time.Sleep(time.Duration(resetDuration) * time.Millisecond)
			FlushCb()
			syncMap.Range(func(key, value any) bool {
				v1, loaded := syncMap.LoadAndDelete(key)
				if loaded {
					v := v1.(*uint64)
					fmt.Println(key, *v)
				}
				return true
			})
		}
	}()
}

func Incr(key string, count uint64) {
	actual, loaded := syncMap.LoadOrStore(key, &count)
	if loaded {
		num, _ := actual.(*uint64)
		atomic.AddUint64(num, count)
	}
}
