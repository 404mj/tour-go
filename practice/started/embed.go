package main

import (
	"fmt"
)

/**
 * 1、内部类型实现了某个接口，那么外部类型也被认为实现了这个接口
 * 2、外部类型也可以声明同名的字段或者方法，来覆盖内部类型的，这种情况方法比较多
 * 3、我们在初始化的时候，采用的是字面值的方式，所以要按其定义的结构进行初始化
 * 4、
 *
 */
func main() {

	type user struct {
		name string
		age  int
	}

	type admin struct {
		user
		level string
	}

	type role struct {
		u     user
		level string
	}

	admin1 := admin{user{"zhangSan", 23}, "admin"}
	fmt.Println("直接调用嵌入类型的字段，名字为: ", admin1.name)
	fmt.Println("通过类型名简介调用嵌入类型的字段，名字为: ", admin1.user.name)

	fmt.Println("嵌入的时候指定名字：")
	admin2 := role{user{"liSi", 23}, "manager"}
	// 必须这样！！！
	fmt.Println("直接调用嵌入类型的字段，名字为: ", admin2.u.name)

}
