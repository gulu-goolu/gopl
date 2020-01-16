package main

import (
	"fmt"
	"os"
	"time"
)

// 倒计时
func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字节
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		default:
			// 轮询
			// fmt.Println("default")
		}
	}

	fmt.Println("Launch")
}
