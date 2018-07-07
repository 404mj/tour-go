package main

import (
	"fmt"
	"math"
)

func main() {
	// sqrtIter(2.0, 4.0)
	// fmt.Println(math.Sqrt(10))
}

//使用牛顿迭代法求平方跟
func sqrtIter(guess float32, x float32) {
	if isGoodEnough(guess, x) {
		fmt.Println(guess)
	} else {
		sqrtIter(improve(guess, x), x)
	}
}

func isGoodEnough(guess float32, x float32) bool {
	g := (guess + guess/x) / 2
	f := math.Abs(float64(g - x))
	return f < 0.001
}

func improve(guess float32, x float32) float32 {
	return (guess + guess/x) / 2
}
