package main

import "fmt"

func main() {

	ch := make(chan int)

	// fatal error: all goroutines are asleep - deadlock!
	ch <- 10 // 這行必須寫在 goroutine 裡面！

	i := <-ch

	fmt.Println(i)
}
