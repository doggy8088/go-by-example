package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args provides access to raw command-line arguments.
	// Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
	argsWithProg := os.Args
	fmt.Println(argsWithProg)

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	// You can get individual args with normal indexing.
	arg := os.Args[3]
	fmt.Println(arg)

	// 手動執行以下命令
	// go run . 1 2 3 4 5

}
