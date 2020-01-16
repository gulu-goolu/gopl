package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func main() {
	// do nothing
}

// 为从通道 filenames 收到的文件名生成缩略图，并返回总的字节数
func makeThumbnail6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()

			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}

			info, _ := os.Stat(thumb)

			sizes <- info.Size()
		}(f)

	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
