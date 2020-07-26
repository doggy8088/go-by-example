package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// i 是索引值，num 是 array 的元素
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 取得字串的每個 Unicode code points
	// 這裡的每個 c 都是一個 rune
	for i, c := range "多奇" {
		fmt.Printf("%d %d %x %U %s\n", i, c, c, c, string(c))
	}

	// Convert between byte array/slice and string
	// https://yourbasic.org/golang/convert-string-to-byte-slice/

	// Strings, bytes, runes and characters in Go
	// https://blog.golang.org/strings

	// 重新认识 Go 中的字符串
	// https://learnku.com/docs/go-blog/strings/6556

}
