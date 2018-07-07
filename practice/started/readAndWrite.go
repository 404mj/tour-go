package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/**
 *
 *互斥锁影响性能，
 *其实真正的互斥应该是 **读取和修改**、**修改和修改之间**，**读取和读取** 是没有互斥操作的。
 * 也就是读写锁，类似一种读者写者问题
 */

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

func main() {
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go read(i)
	}
	for i := 0; i < 5; i++ {
		go write(i)
	}
	wg.Wait()
}
func read(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n", n)
	v := count
	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	wg.Done()
	rw.RUnlock()
}
func write(n int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()
	rw.Unlock()
}
