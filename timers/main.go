package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	// 等兩秒
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	// 非同步中等待一秒
	go func() {
		// 這段會有 goroutine leak 的問題
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// 計時器被停止 (Q: 計時器停止後，第20行的 channel 會 blocking 嗎？為什麼第21行沒有執行？)
	// https://golang.org/pkg/time/#Timer.Stop
	stop2 := timer2.Stop() // Stop does not close the channel, to prevent a read from the channel succeeding incorrectly.

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 結束前等待兩秒
	time.Sleep(2 * time.Second)
}
