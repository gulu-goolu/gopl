package main

import "fmt"

// 向管道中写入数据，所有的数写入完毕后，关闭管道
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// 将输入管道 in 中的数取平方后放入输出管道 out 中
// 数据处理完后，关闭输出管道 out
func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

// 打印输入管道 in 中的数
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}
