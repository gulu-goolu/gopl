package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", true, "show verbose progress messages")

// 每隔 0.5s，程序会往控制台上打印一个输出
func main() {
	// 解析命令行参数
	flag.Parse()

	// 遍历文件树
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 设置定时器
	var tick <-chan time.Time
	if *verbose {
		fmt.Println("verbose")
		tick = time.Tick(500 * time.Millisecond)
	}

	// 遍历文件夹
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir2(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64

	// 遍历文件夹
	// 不停地循环，直至管道 fileSizes 被关闭
loop:
	for {
		select {
		case <-tick:
			printDiskUsage2(nfiles, nbytes)
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}

			nfiles++
			nbytes += size
		}
	}
	// 输出最终结果
	printDiskUsage2(nfiles, nbytes)
}

func printDiskUsage2(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir2(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents2(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir2(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents2(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}
