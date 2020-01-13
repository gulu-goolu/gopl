package main

import (
	"fmt"
	"sort"
)

// 判断两个 map 是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, val := range x {
		if yv, ok := y[k]; !ok || yv != val {
			return false
		}
	}

	return true
}

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages["bob"] = 21

	for name, age := range ages {
		fmt.Println(name, age)
	}

	var names []string
	for name, _ := range ages {
		names = append(names, name)
	}
	// 按名称排序
	sort.Strings(names)

	// 输出年纪
	for _, name := range names {
		fmt.Println(name, ages[name])
	}
}
