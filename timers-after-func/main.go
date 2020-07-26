package main

import (
	"fmt"
	"time"
)

func main() {

	// 計時 2 秒
	timer1 := time.AfterFunc(time.Second*2, func() {
		fmt.Println("Timer fired")
	})

	// 非同步等待 1 秒
	time.Sleep(time.Second * 1)

	// 計時器到達前停止計時，用以取消非同步執行
	if timer1.Stop() {
		fmt.Println("Timer stopped")
	}

	// 結束前等待兩秒，用以驗證 callback func 沒有被觸發！
	time.Sleep(2 * time.Second)

}
