package example

import "fmt"

func List() {
	fmt.Println("List")
}

// 求和
func Sum(args ...int) int {
	var n int
	for _, i := range args {
		n += i
	}

	return n
}
