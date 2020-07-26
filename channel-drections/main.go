package main

import "fmt"

// chan<- 代表只能將訊息發送到 channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// <-chan 代表只能從 channel 取出訊息
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
