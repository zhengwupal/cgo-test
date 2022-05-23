package pack1

import (
	"fmt"
)

func Test4() {
	done := make(chan int, 10) // 带 10 个缓存

	// 开 N 个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("你好, 世界")
			done <- 1
		}()
	}

	// 等待 N 个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
