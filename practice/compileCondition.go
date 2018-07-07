package main

import (
	"fmt"
	"xuetuis.me/practice/json"
)

type user struct {
	Name   string
	Age    int
	WeChat string
}

func main() {
	u := user{"zhangSan", 12, "zs_lalalla"}
	b, err := json.MarshalIndent(u, "", "")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}
