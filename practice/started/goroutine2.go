package main

import (
	"fmt"
)

//******select用在chan的处理上*
func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)

	go func() {
		fmt.Println(1)
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1:", v)

			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2:", v)
			}
		}
	}()

	fmt.Println(2)
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	fmt.Println("input chanel")

	close(c1)
	close(c2)

	fmt.Println("o:", <-o)
}
