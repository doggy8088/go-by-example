package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	// 同時 50 個非同步 goroutine 作業
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// 每個 goroutine 對 ops 執行 1000 次 +1 原子操作
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
