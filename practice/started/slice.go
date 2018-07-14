package main

import (
	"fmt"
)

/**
 * 理解切片要在理解数组的基础上进行
 * 要记住切片是对底层数组的引用，
 * 执行底层数组的指针，一个len字段一个cap字段
 * 切片操作并不复制切片指向的元素，它创建一个新的切片并*复用*原来切片底层的数组
 * so，通过新切片修改数据会影响到原先的切片
 * 
 * 函数之间的传递修改会改变值！类比java 集合
 */

/**
 * 声明赋值打印基础用法
 */
func assign() {
	// 变量声明之后不用也会报错
	s1 := make([]int, 5)
	s2 := []int{'a', 'b', 'c', 'd', 'g'} // 长度自动判断了
	s22 := []string{"a", "b", "c", "d", "hh"}
	s3 := []int{2: 99, 4: 88}
	fmt.Println(s1)
	fmt.Printf("%c \n", s2)
	//Go语言处理字符时，97 和 a 都是指的是字符a，而 Go 语言将使用数值指代字符时，将这样的数值称呼为 rune 类型
	fmt.Println(s22) // string是双引号，这样就可以了！
	fmt.Println(s3)

	// 基于数组和现有切片
	slice2 := []int{11, 22, 33, 44, 55}
	array1 := [5]int{10, 12, 14, 16, 18}

	// 用原有slice,其实是源slice[start] --> slice[end-1] 的元素一共是end-start个元素
	slice22 := slice2[1:3]
	// 用数组构造slice，长度和上面一样，只不过第三个是参数控制容量
	slice1 := array1[1:4:5]
	fmt.Println(slice22)
	fmt.Println(slice1)

	// 两个slice引用同一个底层数组，任一个修改都会作用到其他的！
	fmt.Printf("原来：%v \n", slice2)
	slice22[1] = 666
	fmt.Printf("现在：%v \n", slice2)

	//迭代切片
	for i, v := range slice2 {
		fmt.Printf("index: %d, value: %d \n", i, v)
	}

}

/**
 * 深入原理实现的一些例子
 */
func deepin() {
	// 切片还有len和cap函数
	// 切片的生长 copy和append函数！！！
	sa := []int{1, 2, 3}
	sb := make([]int, len(sa), (cap(sa)+1)*2) // +1 in case cap() return 0

	fmt.Printf("容量翻倍：%v \n", cap(sb))
	//使用copy函数
	copy(sb, sa) // sa的数据复制到sb里面
	fmt.Printf("新建的容量翻倍的：%v \n", sb)
	sa = sb
	fmt.Printf("使用copy扩容：%v \n", cap(sa))

	// 内置了更完善的函数append
	a := make([]int, 1)
	fmt.Println(a)
	a = append(a, 1, 2, 3)
	fmt.Println(a)

	// 把一个切片追加到切片，需要注意！！！
	b := []int{33, 44, 55}
	// a = append(a, b) // Error!
	a = append(a, b...) // 应该这样写！！！！！
	fmt.Print("使用append扩容：")
	fmt.Println(a)

}

/**
 * 常用的标准库函数len，cap，copy， append
 */

/**
 * 猜不到结果系列！
 */
func wired() {
	vals := make([]int, 5)//
	// vals :=[]int{}//比较以上两种方式的异同
	// vals2 := make([]int)  // 这样是ERROR！，可以不指定容量，但是长度是必须指定！！！！
	fmt.Println(vals)
	for i := 0; i < 5; i++ {
		// vals = append(vals, i)
		vals[i] = i;
	}

	fmt.Println(vals)
}

func param(s *[]int) {
	(*s)[2] = 5;
}
func param2(s []int) {
	s[2] = 99
}

func main() {
	// wired()
	// assign()

	//函数间传递切片是值传递，复制的是切片的数据结构，而不是底层数组，
	//所以不需要指针，函数中的修改会作用到原始切片上的！

	// deepin()
	s := []int{1,2,3}
	// param(&s)
	param2(s)
	fmt.Printf("values is %v", s)
}
