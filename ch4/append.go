package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 如果容量足够，从 x 从新构造一个 slice
		z = x[:zlen]
	} else {
		// 分配一个新的 slice
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	x = appendInt(x, 4)
	fmt.Println(x)
}
