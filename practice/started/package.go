package main

//1. main包是一个特殊的包,表明他不是一个库,是一个可执行文件，虽然在一个特定的包内，还是要表明main包
//2. 首字母大写是public，也是用;结尾但是可以省略
/**
 *一个包中的所有文件都必须使用相同的 名称
 * 可执行命令必须使用package main
 *
 *
 */

import (
	"fmt"
	//对go语言约定的这种包的方式不是很理解.这是引入其他的库
	"xuetuis.me/codebits/gopkg"
	// 引入就必须使用，否则报错！
	// 要是一个包下面有好几个文件呢？
	"xuetuis.me/codebits/gopkg/gopkg3"

	//还可以引入自己包下面的子包
	inpg "./innerpkg"
	//可以引入匿名包 只执行包里面的init方法，使用数据库经常这样使用，
)

func main() {
	fmt.Println("hello,this is main")
	gopkg.Say() // 打印了pkg1的方法
	// pkg2.Say()// error 如果我想 使用gopkg包中pkg2的方法怎么班？？？
	// innerpkg.Say()
	//使用别名
	inpg.Say()

	gopkg3.Say()
	//golang中同时满足main包和main函数才会被编译成可执行文件，是程序的入口。也就是说main中可以有非main函数，但是
	//go run 都会默认寻找main方法，所以这里只有非main方法会报错
	// NOTE:golang中函数！=方法要注意！！
	mainSay()
}

func mainSay() {
	fmt.Println("I'm func in main package invoke by main func")
}
