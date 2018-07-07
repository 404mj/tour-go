package main

import "fmt"

func main() {
	// p := person{"zsx", 12}
	// p.modify()
	// fmt.Println(p.name)

	//多值返回
	// a1, a2, a3 := returnMulti()
	// fmt.Println(a1, a2, a3)

	//可变参数
	varyingParam(1, 2, 3, 6, 5, 4)
}

func basicType() {
	//基本类型：数值，浮点，字符，bool(1byte,不可变，不能用数字代替true/false)。
	//复数类型！complex64/complex128
	//值类型：array， struct， string
	//引用类型：slice， map， chan
	//接口：interface
	//函数类型：func
	//不可变，函数传递副本
	/**
	 * 双引号用来创建 可解析的字符串字面量 (支持转义，但不能用来引用多行)；
	 *	反引号用来创建 原生的字符串字面量 ，这些字符串可能由多行组成
	 * (不支持任何转义序列)，原生的字符串字面量多用于书写多行消息、HTML以及正则表达式。
	 */
}

/**
 *引用类型
 */
func refType() {
	//它的修改可以影响到任何引用到它的变量
	// slice, map, interface, func, chan
	//函数传递引用，修改
}

/**
 * 结构体
 * ------------
 * 访问控制：
 * 一个导出的类型，包含了一个未导出的方法也一样，也是无法访问的。
 */
func structType() {
	//
	//
	//结构体传递的是其本身以及里面的值的拷贝。
}

/**
 * 自定义类型
 */
func selfDefType() {
	// Go的编译器不会像Java的那样，帮我们做隐式的类型转换。
	// so , type aaa int 和 int不一样
}

/**
 * 函数类型
 * go中函数不等于方法！！！！
 * 函数是指不属于任何结构体、类型的方法，
 * 也就是说，函数是没有接收者的；而方法是有接收者的，
 * 我们说的方法要么是属于一个结构体的，要么属于一个新定义的类型的。
 * ---------------------------------
 * 函数和方法的访问控制
 * 使用一个可导出的函数可以将一个无法导出的变量导出！！必须使用:=这样的操作符才可以
 */
func funcType() {

}

type person struct {
	name string
	age  int
}

func (p person) modify() {
	// p.name = "lalala"
	p.age = 23
}

//*多个返回值要在括号里面声明类型！
func returnMulti() (int, int, int) {
	return 1, 2, 3
}

/**
 * 可变参数 函数！！
 */
func varyingParam(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}
