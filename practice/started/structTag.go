package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	// 反引号包裹的就是tag
	Name string `name`
	Age  int    `age`
}

func main() {
	// json2Struct()
	// reflectGetTag()
	multiTags()
}

func json2Struct() {
	var u User
	h := `{"name":"zhangsan", "age":15}` // 为啥是反引号？？？
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
}

//反射过去字段的TAg
func reflectGetTag() {
	var u User
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag)
	}
}

//很多时候我们的一个Struct不止具有一个功能，比如我们需要JSON的互转、还需要BSON以及
//ORM解析的互转，所以一个字段可能对应多个不同的Tag，以便满足不同的功能场景。
type User2 struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}

func multiTags() {
	var u User2
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag.Get("json"), ",", sf.Tag.Get("bson"))
	}
}
