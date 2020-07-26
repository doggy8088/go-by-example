package main

import (
	"fmt"
	"time"
)

func main() {

	num := 10

	// basic rate limiting

	requests := make(chan int, num)
	for i := 1; i <= num; i++ {
		requests <- i
	}
	close(requests)

	// rate: 0.2 seconds
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//

	burstyLimiter := make(chan time.Time, 3)

	// 先預留 3 筆，讓執行的時候前三個 job 可以先跑，剩下的依序跑完！
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, num)
	for i := 1; i <= num; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
