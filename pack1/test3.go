package pack1

import (
	"fmt"
	"sync"
)

func Test3() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("你好, 世界")
		mu.Unlock()
	}()

	mu.Lock()
}
