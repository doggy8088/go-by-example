package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	queue := make(chan string, 1)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(15 * time.Millisecond)
			queue <- strconv.Itoa(i)
		}
	}()

	go func() {
		time.Sleep(time.Second)
		close(queue)
	}()

	// queue 會一直讀到被 close 才會從 range 中停止！
	for elem := range queue {
		fmt.Println(elem)
	}
}
