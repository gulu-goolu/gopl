package main

import "fmt"

// 捕获了外部变量 x，此处的 x 分配在堆上
func squares() func() int {
	var x int
	return func() int {
		x++
		return x* x
	}
}

func main() {
	f:=squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

