package main

//golang中支持rpc的库的demo:https://golang.org/pkg/net/rpc
///https://www.cnblogs.com/jkko123/p/7039675.html
// https://studygolang.com/articles/8497
// 以上三个链接学习
// go原生对rpc的支持有三个类被，http， tcp和jsonrpc一下分别示例

import (
	"errors"
	"log"
	"net"
	"net/rpc/jsonrpc"
	// "net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

//服务端监听！
func main() {
	arith := new(Arith)
	//注册rpc服务！
	//***受限于Go语言的特点，
	//我们不可能在接到客户端的请求的时候，根据反射动态的创建一个对象，像Java那样????
	//go中编译好一个map对应名字和服务
	rpc.Register(arith)

	//一****HTTP****
	/*rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	chkErr(err)*/

	//二****TCP****
	//tcp4 means ipv4!
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1234")
	chkErr(err)
	tcpServer, err := net.ListenTCP("tcp", tcpaddr)
	chkErr(err)
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			continue
		}
		//使用goroutine处理tcp请求
		// go rpc.ServeConn(conn)

		//以上都是使用golang中默认的gob编码
		//三****JSONRPC****
		go jsonrpc.ServeConn(conn)
	}

}

func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
