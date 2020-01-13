package main

import "fmt"

// 就地反转 slice 中的元素
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = 1+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}
