package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 简历连接
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	// 从 conn 中读取消息
	go func() {
		// 处理来自 server 的输入
		io.Copy(os.Stdout, conn)

		// 通知 main goroutine，服务已结束
		fmt.Println("done")
		ch <- struct{}{}
	}()

	// 将来自用户的输入复制到 Stdin 中，此操作会一直执行，知道遇到 EOF 标志
	io.Copy(conn, os.Stdin)
	conn.Close()

	// 等待输出 goroutine 退出
	<-ch
}
