package main

import (
	"fmt"
	"log"
	// "net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	//***HTTP***
	// client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")

	//***TCP***
	// client, err := rpc.Dial("tcp", "127.0.0.1:1234")

	//***JSONRPC***
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")

	if err != nil {
		log.Fatal("dialing: ", err)
	}

	//同步调用
	args := Args{7, 8}
	var reply int
	//注意结果是指针类型
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	// Asynchronous call
	// quotient := new(Quotient)
	// divCall := client.Go("Arith.Divide", args, quotient, nil)
	// replyCall := <-divCall.Done // will be equal to divCall
	// check errors, print, etc.
}
