package main

import (
    "fmt"
)

func main() {
    count:=0;
    x:=9999
    for(x != 0) {
        fmt.Println("x: ", x)
        count++
        x = x & (x - 1)
    }

    fmt.Println(count)
}

