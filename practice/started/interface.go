package main

/**
 * interface类型，本质是抽象和多态！
 *接口用来定义行为！接口声明行为，
 *具体的类型去实现这个行为，
 * 调用的时候就会有多态了
 *
 *如果是值接收者，实体类型的值和指针都可以实现对应的接口；如果是指针接收者，那么只有类型的指针能够实现对应的接口。
 */
import "fmt"

type animal interface {
	printInfo()
}

type cat string
type dog string

func (c cat) printInfo() {
	fmt.Println("a CAT")
}
func (d dog) printInfo() {
	fmt.Println("a DOG")
}

func main() {
	var i_a animal
	var c cat
	i_a = c
	i_a.printInfo()

	//多态
	var d dog
	i_a = d
	i_a.printInfo()
}

// ***方法集***
