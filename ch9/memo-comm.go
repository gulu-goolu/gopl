package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// 通过通信共享的方式来实现缓存
// 特性：
// 1. 采用 C/S 架构，当接收到一个来自 Client 的请求后，开启一个新的 goroutine 来处理请求，并将响应通过管道返回给 Client
// 2. 限制后台的请求数。

// 性能分析
// 服务端限制连接数：16
// 1.20s 1.37s 1.29s
// 服务端连接数限制：1
// 1.38s 1.29s
// 限制服务端的并行连接数

// 获取 HTTP 响应的 Body 部分
func commHttpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// 缓存
type CommFunc func(url string) (interface{}, error)
type CommData struct {
	value interface{}
	err   error
}
type CommReq struct {
	// 请求
	url string
	// 用于写入返回值的管道
	rep chan CommData
}

type CommCache struct {
	// 函数
	f CommFunc
	// 维持连接请求
	reqs chan CommReq
}
type CommEntry struct {
	ready chan struct{}
	data  CommData
}

func (entry *CommEntry) call(f CommFunc, url string) {
	// 获取资源
	res, err := f(url)
	entry.data = CommData{res, err}

	// 通知操作已完成
	close(entry.ready)
}

func (entry *CommEntry) deliver(response chan<- CommData) {
	<-entry.ready

	response <- entry.data
}

func commCreateCache(f CommFunc, n int) *CommCache {
	// 用于共享通信，可以允许多个 Client 同时请求
	c := CommCache{f, make(chan CommReq, n - 1)}

	// 开启一个 goroutine 来处理用户的请求
	go func() {
		m := make(map[string]*CommEntry)

		// 通过多路复用来处理请求
		for {
			select {
			case t := <-c.reqs:
				// 处理来自用户的请求
				entry, ok := m[t.url]
				// 如果缓存中没有数据，通过 f 函数加载
				if !ok {
					entry = &CommEntry{ready: make(chan struct{}), data: CommData{nil, nil}}
					m[t.url] = entry

					// 调用函数获取资源
					go m[t.url].call(c.f, t.url)
				}

				// 使用一个新的 goroutine 来处理用户的请求
				go m[t.url].deliver(t.rep)
			}
		}
	}()

	return &c
}

// CommCache 的 GET 方法
func (c *CommCache) Get(url string) (interface{}, error) {
	// 构造一个请求
	ch := make(chan CommData, 1)
	req := CommReq{url, ch}

	// 发送请求
	c.reqs <- req

	// 从管道中读取响应
	data := <-ch

	return data.value, data.err
}

var n = flag.Int("n", 1, "goroutine count")

func main() {
	// 待访问的 url
	urls := []string{
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.bing.com",
		"https://www.bing.com",
		"https://www.bing.com",
		"https://www.bing.com",
	}

	flag.Parse()
	// 构造 CommCache 结构
	c := commCreateCache(commHttpGetBody, *n)

	// 访问 url，并打印访问每个连接所花费的时间
	var n sync.WaitGroup

	total := time.Now()
	for _, url := range urls {
		// 同时开启多个 goroutine 并发执行
		n.Add(1)
		go func(url string) {
			// 延迟退出
			defer n.Done()

			start := time.Now()
			res, _ := c.Get(url)
			fmt.Printf("GET %s, %s, %d bytes\n", url, time.Since(start), len(res.([]byte)))
		}(url)
	}
	// 等待所有的 goroutine 完成
	n.Wait()
	fmt.Printf("total %s\n", time.Since(total))
}
