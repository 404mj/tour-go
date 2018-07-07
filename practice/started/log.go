package main

import (
	"log"
)

func main() {
	log.Println("xuetuis.me", "zsx")

	//log.Fatal() 类似print和os.Exit(1)
}

//***init函数***会在main函数执行之前执行，进行一些初始化的工作
func init() {
	// 定制化的设置前缀
	log.SetPrefix("[TEST]")
	log.SetFlags(log.Ldate | log.Lshortfile)
}
