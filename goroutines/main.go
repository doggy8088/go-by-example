package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 讓主程式睡 1 秒
	time.Sleep(time.Second)
	fmt.Println("done")

	// 建議用 WaitGroup 的方式實作等待
}
