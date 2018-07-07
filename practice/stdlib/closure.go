package main

import (
	"fmt"
	"sync"
)

var a = []string{"a", "b", "c"}

func main() {
	//for-range 闭包的坑
	wrong()
	// correct()

}

func wrong() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	for _, v := range a {
		go func() {
			//运行三次后，是闭包，实际v指向的是slice中最后一个元素
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func correct() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	for _, v := range a {
		//以参数的形式传进去
		go func(s string) {
			fmt.Println(s)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
