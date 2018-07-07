package main

// https://studygolang.com/articles/3365?t=1493776691081  处理命令行参数的库
//http://www.01happy.com/golang-command-line-arguments/  golang 获取命令行参数
import (
	"flag"
	"fmt"
	"os"
)

var para1 *string = flag.String("param1", "", "usage: --param1=value1")

func main() {
	flag.Parse()
	fmt.Println(*para1)
	args := os.Args //获取用户输入的所有参数
	if len(args) > 0 {
		fmt.Println(args[1])
	}
}
