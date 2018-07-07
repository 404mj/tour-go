package main

//通过通信来共享内存，而不是通过共享内存来通信

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//全局变量
var (
	count int32
	g_wg  sync.WaitGroup

	//互斥锁
	mutex sync.Mutex
)

/**
 *
 * 进程	一个程序对应一个独立程序空间
 * 线程	一个执行空间，一个进程可以有多个线程
 * 逻辑处理器	执行创建的goroutine，绑定一个线程
 * 调度器	Go运行时中的，分配goroutine给不同的逻辑处理器
 * 全局运行队列	所有刚创建的goroutine都会放到这里
 * 本地运行队列	逻辑处理器的goroutine队列
 *
 */
func main() {
	// waitgroup()
	syncComp()

}

//自己需求实现一个随机数生成器！
func generateRand(n int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rint := rand.New(seed)
	return rint.Intn(n)
}

func syncComp() {
	g_wg.Add(2)

	// go inc()
	// go inc()

	// go incAtomic()
	// go incAtomic()

	go incSync()
	go incSync()

	g_wg.Wait()
	fmt.Println(count)
}

func waitgroup() {

	// 设置逻辑处理器个数和CPU核数相同！
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	// goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("A: ", i)
			// 速度太快看不出交叉运行，使用time包增加运行时间！
			time.Sleep(300 * time.Millisecond)
		}
	}()

	//goroutine 2
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("B: ", i)
			time.Sleep(500 * time.Millisecond)
		}
	}() // ***括号必不可少！***

	wg.Wait()
}

//资源竞争，count相当于共享资源
// *** go build -race go提供了工具进行检测！！！****
// *** 生成可执行文件，然后执行，就会生成报告***
func inc() {
	defer g_wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		//使用一个随机数产生不同的执行时间，更明显的看到资源竞争的现象
		n := time.Duration(generateRand(6))
		// fmt.Println(n)
		time.Sleep(n * 100 * time.Millisecond)
		value++
		count = value
	}
}

// 使用atomic操作同步资源！
func incAtomic() {
	defer g_wg.Done()
	for i := 0; i < 3; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		// n := time.Duration(generateRand(6))
		// time.Sleep(n * 100 * time.Millisecond)
		value++
		atomic.StoreInt32(&count, value)
	}
}

// 使用 同步锁 标志临界区，访问临界资源count
func incSync() {
	defer g_wg.Done()
	for i := 0; i < 3; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}
