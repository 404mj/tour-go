package main

import (
	"context"
	"fmt"
	"time"
)

/**
 *传统中控制并发的方式是：WaitGroup（很像一个barrier！）,Channel+select
 *比如一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，
 *这些goroutine又可能会开启其他的goroutine。
 *
 */

/**
 * Context使用原则！！！
 * 1、 不要把Context放在结构体中，要以参数的方式传递
 * 2、 以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
 * 3、 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
 * 4、Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
 * 5、 Context是县城安全的，可以放心的在多个goroutine中传递
 *
 */
func main() {
	// demo()
	// controlMultiRoutines()
	withValContext()
}

// 使用context的实例。控制一个goroutine。
func demo() {
	ctx, cancel := context.WithCancel(context.Background()) //返回一个空的context，作为整个context树的root。创建一个可以cancel的
	//context。使用cancel函数发送结束标志
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Watching Quit, Stop......")
				return
			default:
				fmt.Println("goroutin Watching....")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("It's OK, Notice Watching Stop....")
	cancel()
	time.Sleep(5 * time.Second)
}

// 控制多个goroutine
//但是，这多个goroutine属于一个context。怎么更精细的控制？？？
func controlMultiRoutines() {
	ctx, cancel := context.WithCancel(context.Background())

	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

//context接口
//context的继承衍生:WithCancel, WithDeadline, WithTimeOut, WithValue
var key = ""

func withValContext() {
	ctx, cancel := context.WithCancel(context.Background())

	//附加值
	valueCtx := context.WithValue(ctx, key, "lalal")

	go watchWithVal(valueCtx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watchWithVal(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
