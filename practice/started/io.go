package main

// http://www.flysnow.org/2017/05/08/go-in-action-go-reader-writer.html

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//定义零值Buffer类型
	var b bytes.Buffer

	//使用Write方法写入字符串
	b.Write([]byte("你好"))

	// 拼接字符串到Buffer里
	fmt.Fprint(&b, ",", "http://www.xuetuis.me")

	b.WriteTo(os.Stdout)
	fmt.Println("\n.......................")
	var p [100]byte
	n, err := b.Read(p[:])
	fmt.Println(n, err, string(p[:]))

}
