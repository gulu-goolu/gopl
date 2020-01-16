package main

import (
	"fmt"
	"io"
	"sort"
	"time"
)

func compareBetweenInterface() {
	var x interface{} = time.Now()
	fmt.Println(x == x)

	// 不能比较不可比较类型
	var y interface{} = []int{1, 2, 3}
	// fmt.Println(y == y)
	fmt.Println(x == y) // 不同类型可以比较

	// 打印动态类型名
	fmt.Println("%T\n", &x)
}

func f(out io.Writer) {
	fmt.Println(out)
	// out 是一个包含空指针的非空接口
	// 其包括 io.Writer 的具体类型和底层的值
	if out != nil {
		out.Write([]byte("done\n"))
	}
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortSlice() {
	s := []string{"99893", "456", "789"}
	sort.Sort(StringSlice(s))
	fmt.Println(s)
}

func main() {
	sortSlice()
}
