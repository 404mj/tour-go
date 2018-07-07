package main

/**
 * @Author:zsx
 * @Date:2018-02-01 22:59
 */
import (
	"bytes"
	"fmt"
	// "unicode"
)

//bytes包中实现了大量对[]byte操作的函数和两个最主要的Reader和Buffer两个结构
//由此引发了对go中string byte rune的思考：string.go

func main() {
	tour()
}

func tour() {
	// http://www.cnblogs.com/jkko123/p/7221544.html

	// str := "aBsDLinux"
	//小写
	// fmt.Println(string(bytes.ToLower([]byte(str))))
	// fmt.Println(string(bytes.ToTitle([]byte(str))))

	//比较
	fmt.Println(bytes.Equal([]byte("a"), []byte("b")))
	fmt.Println(bytes.Compare([]byte("acm"), []byte("asm")))
	fmt.Println(bytes.EqualFold([]byte("aCm"), []byte("acM"))) //忽略大小写

	//去掉[]byte两边包含在cutset中的字符
	fmt.Println(string(bytes.Trim([]byte(" abc "), " ")))

	//去掉左边包含在cutset中的字符
	fmt.Println(string(bytes.TrimLeft([]byte(" abc "), " ")))

	//去掉两边空白字符
	fmt.Println(string(bytes.TrimSpace([]byte(" abc "))))

	//去掉前缀
	fmt.Println(string(bytes.TrimPrefix([]byte("tb_user"), []byte("tb_"))))

	//去掉后缀
	fmt.Println(string(bytes.TrimSuffix([]byte("user_idx"), []byte("_idx"))))

	//以sep为分隔符，切分为多个[]byte
	//SplitAfterN 分割最多n个子切片，超出n的部分将不进行切分
	tmp := bytes.Split([]byte("ab cd ef"), []byte(" "))
	for _, v := range tmp {
		fmt.Println(string(v))
	}

}

func readerBuffer() {
	//详见URL
}
