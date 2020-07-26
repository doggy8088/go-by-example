package main

import (
	"fmt"
	"time"
)

func main() {

	// buffer size = 2 (同時最多只能緩存 2 個訊息)
	messages := make(chan string, 2)

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "buffered"
		time.Sleep(1 * time.Second)
		messages <- "channel"
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-messages)
}
