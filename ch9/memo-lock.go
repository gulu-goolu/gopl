package main

// 通过变量上锁来实现

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Memo struct {
	f     Func
	cache map[string]result
}
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// 注意，非并发安全
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}

	return res.value, res.err
}

// 获取 HTTP 的 Body
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// 示例
func example() {
	m := New(httpGetBody)
	urls := []string{"https://wwww.baidu.com", "https://www.baidu.com", "https://www.baidu.com"}
	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func main() {
	// 单进程之心缓存
	example()
}
