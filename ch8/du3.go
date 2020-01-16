package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// 统计指定文件夹下地文件数以及字节数
// 通过 goroutine 来加速

var done = make(chan struct{}) // 通知取消事件

func cancelled() bool {
	// 通过轮询来判断结果
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	// 解析命令参数
	flag.Parse()

	// 从命令行参数中读取参数
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 读取取消标记
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// 遍历
	fileSizes := make(chan int64)
	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir3(root, &n, fileSizes)
	}

	// 所有操作完成后，关闭管道
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			{
				// 管道中的数据读取完毕，跳出循环
				if !ok {
					break loop
				}

				nfiles++
				nbytes += size
			}
		case <-done:
			// 耗尽 fileSizes 以允许已有的 goroutine 结束
			for range fileSizes {
				// 不执行任何操作
			}
		}
	}

	// 打印信息
	showInfo(nfiles, nbytes)
}

func showInfo(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB", nfiles, float64(nbytes)/1e9)
}

// @param dir 路径
// @param n 用于同步的信号
// @param fileSizes 文件数目
func walkDir3(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	// 如果管道关闭，操作终止
	if cancelled() {
		return
	}

	for _, entry := range dirInfos(dir) {
		if entry.IsDir() {
			// 增加一个 goroutine 来遍历
			n.Add(1)
			// 继续遍历
			subDir := filepath.Join(dir, entry.Name())
			go walkDir3(subDir, n, fileSizes)
		} else {
			// 输出文件信息
			fileSizes <- entry.Size()
		}
	}
}

// 读取文件夹下的信息
func dirInfos(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return entries
}
