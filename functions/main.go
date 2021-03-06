package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// Multiple Return Values
func vals() (int, int) {
	return 3, 7
}

// Variadic Functions (可變參數函式)
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 任何 _ 都是無法讀取的，等同於丟棄這個位子的值
	_, c := vals()
	fmt.Println(c)

	// Variadic Functions
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// Closures
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())

	// Recursion
	fmt.Println(fact(10))
}
