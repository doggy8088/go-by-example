package main

import (
	"fmt"
	"time"
)

func main() {

	// 計時 2 秒
	timer2, done := setTimeout(func() {
		fmt.Println("Timer fired")
	}, time.Second*2)

	// 非同步等待 1 秒
	time.Sleep(time.Second * 1)

	// 計時器到達前停止計時，用以取消非同步執行
	if timer2.Stop() {
		done <- true
		fmt.Println("Timer stopped")
	}

	// 結束前等待兩秒，用以驗證 callback func 沒有被觸發！
	time.Sleep(2 * time.Second)
}

func setTimeout(callback func(), timeout time.Duration) (*time.Timer, chan bool) {

	done := make(chan bool, 1)

	t := time.NewTimer(timeout)

	go func() {
		select {
		case <-done:
			return
		case <-t.C:
			callback()
		}
	}()

	return t, done
}
