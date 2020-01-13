package main

import "fmt"

// nonempty 返回一个新的 slice，slice 中的元素都是非空字符串
// 在函数调用的过程中，底层的数组元素发生了改变
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func main() {
	s := []string{"hello", "", "world"}
	for _, t := range s {
		fmt.Println(t)
	}

	fmt.Println("call nonempty proc")
	s = nonempty(s)
	for _, t := range s {
		fmt.Println(t)
	}
}
