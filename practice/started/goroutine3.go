package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println(" ..lala..  ")
		time.Sleep(300 * time.Millisecond)
	}()

	fmt.Println("master thread done!...")
}
