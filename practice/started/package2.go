package main

import (
	// 修改gopkg里面的两个package声明为gopkglala，引入的时候还是gopkg，是以目录引入
	//但是使用的时候可以看到：是用目录下的声明的package来引用的！！！
	// 也就是说：import后面的最后一个元素应该是路径，就是目录，并非包名。
	// NOTE：还需要注意的一点是同一个目录下的go文件的package声明必须是一样的！
	//NOTE: 另外我之前提出的引入一个包内的指定文件的假设是不存在的，包内方法是无法重复定义的！！
	"xuetuis.me/codebits/gopkg"
)

// 该例子表明：一个包内可以包含多个mian包！！！但是运行的话单独运行，一起运行会报redeclared错误！
func main() {
	gopkglala.Say()
	gopkglala.Say2()
}
