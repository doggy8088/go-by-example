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

	// ------------------------------------------------

	done := make(chan bool, 1)

	// 計時 2 秒
	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		select {
		case <-done:
			return // returning not to leak the goroutine
		case <-timer2.C:
			fmt.Println("Timer 2 fired")
		}
	}()

	// 非同步等待 1 秒
	time.Sleep(time.Second)

	// 計時器到達前停止計時，用以取消非同步執行
	stop2 := timer2.Stop() // Stop does not close the channel, to prevent a read from the channel succeeding incorrectly.
	if !stop2 {
		<-timer2.C
	} else {
		done <- true
		fmt.Println("Timer 2 stopped")
	}

	// 結束前等待兩秒
	time.Sleep(2 * time.Second)
}
