package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Second)
		messages <- "ping"
	}()

	// 底下這行會 blocking 現有程式執行，直到取回訊息
	msg := <-messages
	fmt.Println(msg)
}
