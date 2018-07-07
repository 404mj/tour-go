package main

import (
	"fmt"
	"reflect"
)

func main() {
	demo()
}

//任何接口都由<Value, Type>组成
type User struct {
	Name string
	Age  int
}

func demo() {
	u := User{"Zhangsan", 22}
	t := reflect.TypeOf(u)  //获取类型信息
	v := reflect.ValueOf(u) //获取值信息
	fmt.Println(t)
	fmt.Println(v)
	fmt.Printf("%T \n", u)
	fmt.Printf("%v \n", u)

	//reflect.Value转换原始类型
	u1 := v.Interface().(User)
	fmt.Println("convert:", u1)

	//获取底层数据类型
	fmt.Println("t底层数据类型：", t.Kind())

	//遍历字段和方法,***NOTE:这里是用t不是u！！
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	//动态修改字段值
	x := 2
	v1 := reflect.ValueOf(&x) //NOTE:修改值要使用地址
	v1.Elem().SetInt(100)
	fmt.Println(x)

	//动态调用方法
	method := v.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(method.Call(args))
}

func (u User) Print(prefix string) {
	fmt.Println(prefix, u)
}
