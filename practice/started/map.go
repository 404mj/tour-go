package main

import (
	"fmt"
)

func assign() {
	//使用make创建，[]内为key！外边是value
	d1 := make(map[string]int)
	d1["zhangSan"] = 25

	//使用字面量创建初始化
	d2 := map[string]int{"liSi": 19}
	//也可以为空
	d3 := map[string]int{}

	//nil的map，还要再初始化（也就是所谓的开辟内存）
	var d4 map[string]int
	//这里不能使用：=  否则报错，因为:=只用第一次出现的时候！
	d4 = make(map[string]int)
	d4["wangWu"] = 77

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)

	//添加值
	d1["A"] = 12
	d1["B"] = 13
	d1["C"] = 14
	//获取值和使用 ,ok方式获取，这里不多将了，草纸看文档的时候记录了
	valueA, ok := d2["A"]
	fmt.Printf("valueA:%v, ok?:%v \n", valueA, ok)
	//删除键
	delete(d1, "B")

	// 遍历map，是无序的，可以先排序
	fmt.Println("原始的d1,无序：")
	for k, v := range d1 {
		fmt.Println(k, v)
	}

	//排序，对key进行排序，先把key去出来放到slice中
	// sort排序
	// 如果对value进行排序。。https://studygolang.com/articles/10530

	modify(d1)
	fmt.Println("调用函数内修改之后：")
	fmt.Println(d1)
}

func deepin() {
	// map的容量，可以用len函数，make时候制定！

}

func modify(dict map[string]int) {
	// 函数之间传递map不会拷贝副本，所以修改会对所有可见
	dict["A"] = 11111
}

func main() {
	assign()
}
