package main

import "fmt"

// Go Slices: usage and internals
// https://blog.golang.org/slices-intro
// https://learnku.com/docs/go-blog/go-slices-usage-and-internals/6582

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [9][9]int
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			twoD[i-1][j-1] = i * j
		}
	}
	fmt.Println("2d: ", twoD)

	// ---

	// 陣列的長度是固定的
	c := [5]int{1, 2, 3, 4, 5}
	// 陣列內容儲存的是「值」，複製的也是「值」(傳值的成本較高)
	d := c
	d[0] = 100
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	// 陣列的長度是固定的
	e := [5]int{1, 2, 3, 4, 5}
	// 陣列可以將指標傳給變數
	f := &e
	f[0] = 100
	fmt.Println("e:", e)
	fmt.Println("f:", f)

}
