package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	ringDemo()
}

//使用list链表
func listDemo() {
	//创建一个链表
	demoList := list.New()

	//插入元素
	a1 := demoList.PushBack(11)
	demoList.PushBack(22)

	//头部插入元素
	demoList.PushFront(33)
	demoList.PushFront(44)

	printList(demoList)

	//取第一个元素
	firstV := demoList.Front()
	fmt.Println("first: ", firstV.Value)
	//最后一个元素Back，长度demoList.len(),

	// 在某元素之后插入
	demoList.InsertAfter(77, a1)

	printList(demoList)

	//链表之后插入新链表
	demoList2 := list.New()
	demoList2.PushBack(20)
	demoList2.PushFront(22)
	demoList2.PushBackList(demoList)

	printList(demoList2)

	//是可以放其他类型的！！！
	dl := list.New()
	zs := User{name: "zhangsan", age: 13}
	dl.PushBack(zs)
	printList(dl)

}

//其实就是循环链表！
func ringDemo() {
	demoRing := ring.New(5)
	for i := 0; i < 5; i++ {
		demoRing.Value = i * i
		demoRing = demoRing.Next()
	}
	printRing(demoRing)
	fmt.Println("len: ", demoRing.Len())

	//移动环的指针
	demoRing.Move(2)

	//从当前指针删除n元素
	demoRing.Unlink(2)
	printRing(demoRing)

	//连接两个环
	r2 := ring.New(3)
	for i := 0; i < 3; i++ {
		r2.Value = i + 10
		//取得下一个元素
		r2 = r2.Next()
	}
	printRing(r2)

	demoRing.Link(r2)
	printRing(demoRing)
}

func printRing(r *ring.Ring) {
	r.Do(func(v interface{}) {
		fmt.Print(v.(int), " ")
	})
	fmt.Println()
}

type User struct {
	name string
	age  int
}

func printList(l *list.List) {
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ")
	}
	fmt.Println()
}
