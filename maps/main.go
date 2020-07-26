package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 當你不需要回傳值時，才用 _ 當成參數來接
	// 第二個參數 prs 用來儲存執行的結果
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 所有 map 的型別都是 map[T1]T2 來標示！
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
