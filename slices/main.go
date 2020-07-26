package main

import "fmt"

// Go Slices: usage and internals
// https://blog.golang.org/slices-intro
// https://learnku.com/docs/go-blog/go-slices-usage-and-internals/6582

// 切片操作不會複製切片的資料。它建立一個指向原始陣列的新切片值。這使得切片操作與操作陣列索引一樣高效。

func main() {

	// 建立一個數量為 3 的 slice
	// 千萬不能寫成 [3]string 喔！（這樣就變成陣列了）
	s := make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))

	// s 長度為 5 所以先建立一個長度為 5 的 slice 然後再複製過去
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 這裡的 c 雖然是 slice，但是其內容都是 copy 過來的，所以不會改到 s 的資料
	c[0] = "100"
	fmt.Println("c:", c)
	fmt.Println("s:", s)

	// append 函式可以使用「展開運算子」(...)
	u := []string{"Will", "Huang"}
	s = append(s, u...)
	fmt.Println("new:", s)

	l := s[2:5]
	fmt.Println("sl1:", l)

	m := s[:5]
	fmt.Println("sl2:", m)

	// 切片的宣告方式與陣列的宣告方式相同，只是你不需要指定元素個數！
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t[1] = "w"
	fmt.Println("upd:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// ---------------------

	// 這個是陣列
	w := [3]string{"Лайка", "Белка", "Стрелка"}
	x := w[:] // 指向 x 儲存空間的切片
	x[1] = "Will"
	fmt.Println("w:", w)
	fmt.Println("x:", x)

	// 這個是陣列
	y := [3]string{"Лайка", "Белка", "Стрелка"}
	z := y // y 指向陣列，這段是 z 複製了 y 陣列
	z[1] = "Will"
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
