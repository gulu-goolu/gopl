package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 聊天服务器
// 可以在几个用户之间广播文本消息
type client chan <- string
var (
	entering =make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func main() {
	// 监听端口，接受连接客户端的网络连接。对每一个连接，创建一个新的 handleConn goroutine
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	go chatBroadcaster()

	// 处理用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go chatHandleConn(conn)
	}
}

// 广播器，记录当前连接的客户端集合
func chatBroadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <- entering:
			// 加入到 clients 中
			clients[cli] = true
        case cli:= <- leaving:
        	// 从 clients 中移除
    		delete(clients, cli)
    	case msg := <- messages:
    		// 向 clients 中的所有元素广播信息
    		for cli := range clients {
    			cli <- msg
    		}
		}
	}
}

// 处理来自客户的连接请求
// 进入聊天室时，向聊天室中的用户广播 xx 进入了聊天室
// 离开聊天室时，向聊天室中的用户广播 xx 离开了聊天室
// 其他情况下，则直接广播用户的聊天数据
func chatHandleConn(conn net.Conn) {
	// 构
	ch := make(chan string)
	go chatClientWriter(conn, ch)


	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	// 转发数据
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left."
	conn.Close()
}

func chatClientWriter(conn net.Conn, ch <- chan string) {
	for msg := range ch {
		// 忽略错误信息
		fmt.Fprintln(conn, msg)
	}
}
