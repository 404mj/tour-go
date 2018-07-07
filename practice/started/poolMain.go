package main

import (
	"./pool"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxGotoutine = 5
	PoolResSiz   = 3
)

func main() {
	var wg sync.WaitGroup
	wg.Add(MaxGotoutine)

	p, err := poolutils.NewPool(createConnection, PoolResSiz)
	if err != nil {
		log.Println(err)
		return
	}

	//模拟好几个goroutine同时使用资源池查询数据
	for query := 0; query < MaxGotoutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Ready To Close Resourse")
	p.Close()
}

func dbQuery(query int, pool *poolutils.Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	//******conn.(*dbConnection).ID写法？为什么这么写？*******/
	log.Printf("This is %d query, Use the %d connection", query, conn.(*dbConnection).ID) //?????
}

type dbConnection struct {
	ID int32
}

//实现io.Closer() 接口！！
func (db *dbConnection) Close() error {
	log.Println("io.Closer: Close Connection", db.ID)
	return nil
}

var idCounter int32

// 生成数据库链接的方法，以供资源池使用
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}
