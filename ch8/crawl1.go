package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)
	// 从命令行读取初始参数
	go func() {
		worklist <- os.Args[1:]
	}()

	// 并发爬取 web
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}
