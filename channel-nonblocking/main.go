package main

import (
	"fmt"
)

func main() {
	// unbuffered channel
	messages := make(chan string)
	signals := make(chan bool)

	// 如果加上這行 `case msg := <-messages` 就會收到訊息
	// time.Sleep(time.Millisecond)

	// 因為 goroutine 是 non-blocking 的，所以到這行 select 還不會這麼快執行，因此會選 default 執行
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	// 因為 message channel 並沒有 buffer，所以透過 select 的時候 msg 不能直接送到 messages
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	// 只要 messages channel 加上 1 個 buffer size 就可以收到訊息！

	// 如果加上以下 3 行，那麼這兩個 channel 都有可能會被選中
	// go func() { messages <- "First" }()
	// go func() { signals <- true }()
	// time.Sleep(time.Second)

	// 因為兩個 channel 都是 non-blocking 的，因此兩個都收不到！
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
