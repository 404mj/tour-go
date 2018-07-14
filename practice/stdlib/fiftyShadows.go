package main

//***********原文http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
//###########中文http://colobu.com/2015/09/07/gotchas-and-common-mistakes-in-go-golang/

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func sliceRange() {
	x := []string{"a", "b", "c"}

	for v := range x {
		fmt.Println(v) //prints 0, 1, 2
	}
	for _, v := range x {
		fmt.Println(v) //prints a, b, c
	}
}

//二维数组！
func twoDArray() {
	//直接初始化赋值
	var a = [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	// 循环嵌套访问！
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(" ", a[i][j])
		}
		fmt.Println()
	}

	//先声明后赋值！！！
	var b [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			b[i][j] = i + j
		}
	}
	fmt.Println(b[2][3])
	fmt.Println(b[2])

}

//二维切片！
// ***就不一样了！****
func twoDSlice() {
	//直接赋值初始化
	var slice1 = [][]string{
		{"a", "b", "c", "d", "e"},
		{"1", "2", "3", "4", "5"},
		{"I", "II", "III", "IV", "V"},
	}
	//循环嵌套访问,是可以的！
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(" ", slice1[i][j])
		}
		fmt.Println()
	}

// ############update 2018年07月14日15:37:51 这里问题的关键是make方式的使用或者说是二维slice的声明方式！！！！！不同而采用不同的方式初始化@##############
	//这个问题，让我反思之前学习slice，再次深入了解了slice！！！赋值二维slice坑很多啊啊啊啊！
	// 先声明再赋值初始化
	// var slice2 = make([][]int, 0, 3) // len()=0,cap=3
	var slice2 = make([][]int, 3) //len=3,cap=3
	// var slice2 = [][]int{}
	//@@@@@@@@@@@@@以上三种方式针对方法一有不同的输出，
	// @@@@@@@@@@因为slice的make有原始值，执行len为0和空{}，消除了影响！！！！
	//@@@@@@@@@@@@因此，只有执行len为3的方式适用方法二！！！for-range的前提是你有空的啊！！！
	//@@@@@@@也解释了方法一使用ij下标出错的原因！
	// var slice2 [][]int//这种只声明没有初始化不适合方法而！
	/*for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			slice2[i][j] = i * j
		}
		//#########这样是错误的！！
	}*/

	// https://golang.org/doc/effective_go.html#two_dimensional_slices
	//二维slice声明的是行！
	slice2 = init2dslice(slice2) // 使用for-i
	// slice2 = init2dslice2(slice2) // for-range
	fmt.Println(slice2)

}

//正确的赋值二维slice的方法一！
//https://golangtc.com/t/56faadf0b09ecc66b90001d1
func init2dslice(s [][]int) [][]int {
	for i := 0; i < 3; i++ {
		ss := make([]int, 0, 5)
		for j := 0; j < 5; j++ {
			// ss = append(ss, i+j)
			ss[j] = i + j//这样不行！！
		}
		// s = append(s, ss)
		s[i] = ss
	}
	return s
}

//正确赋值二维slice的方法二！！！https://coldfunction.com/go-bot/p/26
func init2dslice2(s [][]int) [][]int {
	for i := range s {
		subS := make([]int, 5)
		for j := range subS {
			subS[j] = j * j
		}
		s[i] = subS
	}
	return s
}

/*
 * https://stackoverflow.com/questions/39561140/go-how-is-two-dimensional-arrays-memory-representation
 *
 * https://blog.golang.org/slices
 *
 *
 *
 */

func stringImmutable() {
	x := "text"
	xbytes := []byte(x)
	xbytes[0] = 'T'

	fmt.Println(string(xbytes)) //prints Text
}

func notUtf8String() {
	//***use package unicode/utf8
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) //prints: true
	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) //prints: false
}

func stringLength() {
	// The built-in len() function returns the number of bytes
	//****also use package unicode/utf8
	data := "♥"
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1
}

func updateValsInRange() {
	data := []int{1, 2, 3}
	for i, _ := range data {
		// v *= 10 //original item is not changed
		data[i] *= 10
	}

	fmt.Println("data:", data) //prints data: [10 20 30]

	//****but,如果有指针类型，该类型的下一层可以改变
	data2 := []*struct{ num int }{{1}, {2}, {3}}

	for _, v := range data2 {
		v.num *= 10
	}

	fmt.Println(data2[0], data2[1], data2[2]) //prints &{10} &{20} &{30}
}

func reslice() {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) //prints: 10000 10000 <byte_addr_x>
	data := raw[:3]
	//会引用原始slice的底层的array
	fmt.Println(len(data), cap(data), &data[0]) //prints: 3 10000 <byte_addr_x>

	//***you SHOULD!
	res := make([]byte, 3)
	copy(res, raw[:3])
	fmt.Println(len(res), cap(res), &res[0]) //prints: 3 3 <byte_addr_y>

}

//****这个坑在closure.go里面说过了
type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func forClosures1() {

	data1 := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data1 {
		// a := v//不加这个不行，这种方式还不能传参数！！
		fmt.Println("---" + v.name)
		// time.Sleep(1 * time.Second)
		// go a.print()
		go v.print()
	}
	time.Sleep(1 * time.Second)
	//####################################
	//一直说是闭包的问题，为什么会出现这个问题，这里是闭包一种最简单的闭包！！！
	//出现这样的结果是为什么！！！？？？没有说
	//我猜测是因为goroutine依赖于迭代的环境，本次迭代的是v1,你还没执行print的时候
	//已经切换到下一次迭代的值去了，为了验证，我在迭代里面加了如注释所示的sleep，发现
	//输出了正常值！！！！
	//其实回头看fifty这篇文章写的应该是这个意思，是我没有理解明白
	//his is the most common gotcha in Go. The iteration variables in for statements are reused in each
	// iteration. This means that each closure (aka function literal) created in your for loop will
	// reference the same variable (and they'll get that variable's value at the time those goroutines start executing).
	// ###################################
}
func forClosures() {
	data := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		go v.print()
	}

	time.Sleep(1 * time.Second)
}

func deferFuncArgCall() {
	var i int = 1
	//必须一行吗？
	//以为是4结果是2？？
	//Arguments are evaluated when the defer statement is evaluated
	// (not when the function is actually executing)
	//可以理解为编译的时候确定的而不是执行的时候
	defer fmt.Println("result=> ", func() int { return i * 2 }())
	i++
}

type user struct {
	name string
}

func updateMapEle() {
	u := map[string]user{"u1": {"user1"}}
	// u["x"].name = "userOne" // cannot assign to struct field u["x"].name in map
	//针对以上问题，两种方式：
	//-----1、-----
	r := u["u1"]
	r.name = "userOne"
	u["u1"] = r
	fmt.Println(u["u1"])

	//-----2、------使用指针！！！
	up := map[string]*user{"1": {"1111"}}
	up["1"].name = "222"
	fmt.Println(up["1"])
}
func updateSliceEle() {
	us := []user{{"one"}}
	us[0].name = "two"
	fmt.Println(us)
}

func nilInterfaces() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil) //prints: <nil> true
	fmt.Println(in, in == nil)     //prints: <nil> true

	in = data
	fmt.Println(in, in == nil) //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'
}

func main() {
	//sliceRance()
	// twoDArray()
	// twoDSlice()
	// stringImmutable()
	// notUtf8String()
	// stringLength()

	//log.Fatal log.Panic 还会结束app
	//map无序，for-range返回不同的结果
	//switch-case匹配执行推出，所以不必多个case累加，相同的类别直接一个case 1,2,3：不要case 1：case 2：do something
	//没有++i
	//按位非不用～而使用^
	//json ,xml等解析的是偶字段应该都是首字母大写！

	// updateValsInRange()
	// reslice()

	//声明一个已有类型的类型，该类型不会继承已有类型的方法
	//如果实在是需要这个类型里面的方法，可以type一个struct或者interface

	forClosures1()
	// forClosures()

	// deferFuncArgCall()
	//另一个关于defer的问题是，defer在包含他的函数结束之前调用而不是包含他的代码
	//块，，这点很重要！不要弄混！

	//***It's OK to call a pointer receiver method on a value as long as the value
	// is addressable. In other words, you don't need to have a value receiver version of the method in some cases.
	// Map elements are not addressable. Variables referenced through interfaces are also not addressable.
	//只要一个值可被寻址！使用该value调用 以value指针类型的方法是可以的！！
	//map中的元素不可寻址！
	// updateMapEle()
	//slice的元素可以寻址！！！
	// updateSliceEle()

	//"nil" Interfaces and "nil" Interfaces Values
	// nilInterfaces()

	//read and write reordering

	//preemptive scheduling 抢先调度

}
