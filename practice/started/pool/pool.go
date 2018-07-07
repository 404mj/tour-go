package poolutils

import (
	"errors"
	"io"
	"log"
	"sync"
)

/**
 *http://www.flysnow.org/2017/05/01/go-in-action-go-pool.html
 *使用有缓冲的通道实现一个资源池,管理多个goroutine之间共享的资源，常见的数据库资源池
 */

var ErrPoolClosed = errors.New("Resource Pool Has Closed")

type Pool struct {
	mutex    sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

//创建资源池
func NewPool(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("invalid size")
	}
	return &Pool{
		factory:  fn,
		resource: make(chan io.Closer, size),
	}, nil
}

//获取资源
// **NOTE: 注意select的使用！！！非阻塞IO可以获取到就获取，不能就生成一个。
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resource:
		log.Println("Acquire:....")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:new>......")
		return p.factory()
	}
}

//关闭资源pool
func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resource) // 关闭channel

	//关闭通道里面的资源
	for r := range p.resource {
		r.Close()
	}
}

//释放资源
func (p *Pool) Release(r io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resource <- r:
		log.Println("Release Resource In Pool ...")
	default:
		log.Println("Pool Is Full, give up....")
		r.Close()
	}
}
