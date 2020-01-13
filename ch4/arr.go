package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a[0])        // 第一个元素
	fmt.Println(a[len(a)-1]) // 末尾元素

	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅输出元素
	for _, v := range a {
		fmt.Printf("%d", v)
	}

	// 使用字面量来初始化数组
	var q [3]int = [3]int{1, 2, 3}
	fmt.Print(q)

	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	// 简化数组长度
	p := [...]int{1, 2, 3}
	fmt.Print(p)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "#", GBP: "&", RMB: "Y"}
	fmt.Println(RMB, symbol[RMB]) // 3 "Y

	// 如果数组中的元素是可比较的，那么数组也是可比较的
	a_1 := [2]int{1, 2}
	b_1 := [...]int{1, 2}
	c_1 := [2]int{1, 3}
	fmt.Println(a_1 == b_1, a_1 == c_1, b_1 == c_1) // "true false false"

}
