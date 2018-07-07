package main

import (
	"fmt"
	"time"
)

/**
 * 其实不管是锁，还是临界资源，本质上都是进之间协作通讯的方式！
 * channel是更好的协作方式！！
 * 如果通道里没有数据的话，接收的数据是nil
 * 一通道关闭，如果发送数据，会引起painc异常。
 *  chan的大小区分有无缓冲
 */
func main() {

}

//无缓冲的channel，当做同步通道来使用，不需要sync.WaitGroup
func firstChannel() {
	// := 就是用在也声明和赋值的情况下
	ch := make(chan int)
	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			time.Sleep(200 * time.Millisecond)
			sum += i
		}
		ch <- sum
	}()

	fmt.Println("I'm wating channel value: ")
	fmt.Println(<-ch)
}

// **** 有缓冲的通道，一个队列 ***！
// make的时候制定大小，cap返回容量，len返回当时大小

// *** 单向通道！，make的时候指定方向 ****
//只有我自己觉的这里很饭人类吗？？
// var chan<- int 只能发送
// var <-chan int 只能接受
// 关闭channel使用close

/**
 *
 * 综合示例见pool.go
 */
