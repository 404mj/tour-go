package main

import (
	"fmt"
)

func main() {
	realWorld()
}

//*** 在Go当中 string底层是用byte数组存的，并且是不可以改变的。
func see() {
	s := "Go大神"
	fmt.Println("len(s):", len(s)) //chinese stored 3 bytes

	ss := rune('神')
	fmt.Println("len(rune(s)):", len(string(ss)))

	skill := []rune(s) //变成rune切片
	fmt.Println(len(skill))
}

//****rune类型是int32的别名，用于表示unicode character （https://www.cnblogs.com/moodlxs/p/4133121.html）
func seeAgain() {
	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Println("byte of abcd: ", b1)

	s2 := "中文"
	b2 := []byte(s2)
	fmt.Println("byte of 中文：", b2)

	b3 := []rune(s2)
	fmt.Println("rune of 中文：", b3)
}

//https://blog.golang.org/strings
//https://www.jianshu.com/p/01a842787637
// *** Rob Pike的blog解释 ***
func realWorld() {
	sample := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// fmt.Println(sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
		//bd(一个字节byte) b2 3d bc 20 e2 8c 98
	}

	//q--quote
	fmt.Printf("%q\n", sample)

	// 采用UTF-8打印
	fmt.Printf("%+q\n", sample)
}

//***In Go, a string is in effect a read-only slice of bytes***
// Go source code is always UTF-8.
// A string holds arbitrary bytes.
// A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
// Those sequences represent Unicode code points, called runes.
// No guarantee is made in Go that characters in strings are normalized.
// Strings are built from bytes so indexing them yields bytes, not characters.
// A string might not even hold characters. In fact, the definition of "character"
// is ambiguous and it would be a mistake to try to resolve the ambiguity by defining
// that strings are made of characters.
//*** https://blog.golang.org/normalization
func realWorld2() {

}
