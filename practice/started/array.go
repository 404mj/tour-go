package main

/**
 * 需要注意的是数组和长度构成一个元素，不同长度不同类型的数组
 * 是不同的！
 *
 */

import (
	"fmt"
)

func assign() {
	//方法一
	var arr1 [5]int //只声明！
	arr1 = [5]int{1, 2, 3, 45, 5}

	//方法二
	arr2 := [5]int{2, 3, 3, 3, 3}

	//方法三
	arr3 := [...]int{4, 4, 4, 4, 4}

	//初始化指定索引为特定的值
	arr4 := [5]int{1: 5, 3: 5}

	//打印数组元素：原始for方法
	for i := 0; i < 5; i++ {
		fmt.Printf("index:%d, value:%d", i, arr1[i])
	}

	//打印数组元素，for-range方法
	for i, v := range arr2 {
		fmt.Printf("index:%d, value:%d", i, v)
		if i == 4 {
			fmt.Println()
		}
	}

	// 直接打印方法
	fmt.Printf("arr3:%v \n", arr3)
	// fmt.Printf("arr3:%v \n", &arr3)

	fmt.Printf("arr4:%#v \n", arr4)

	//只有相同类型的数组可以相互赋值
	var arr5 [5]int = arr4
	fmt.Printf("arr5:%v \n", arr5)

}

func deepin() {
	//指针数组
	arr_p1 := [5]*int{1: new(int), 3: new(int)}
	// Note：只能给new过索引赋值，或者可以先声明再赋值
	*arr_p1[1] = 24
	arr_p1[0] = new(int)
	*arr_p1[0] = 111
	fmt.Println(arr_p1)

}

func deepin2() {
	// 数组不需要显式的初始化；会自动初始化为其对应类型的零值
	var arr [3]int
	fmt.Println(arr)
}

//传递参数就不弄了
//不会修改原始值， 除非传递地址

func functions() {
	// len 获取长度

}

func main() {
	// assign()
	// deepin()
	deepin2()
}
