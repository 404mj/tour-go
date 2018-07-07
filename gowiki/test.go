package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>%s</p>", r)
}

func main() {

	//*** NOTE ***
	// go net/http包中有两种将路径对应到处理处理逻辑的方法HandleFunc 和 实现ServerHTTP接口的object
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
