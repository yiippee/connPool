package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	connPool := NewConnPool(&Options{
		Dialer: func() (net.Conn, error) {
			conn, err := net.DialTimeout("tcp", "127.0.0.1:80", 5*time.Second)
			return conn, err
		},
		PoolSize:    1024,
		PoolTimeout: 600 * time.Second,
		IdleTimeout: 600 * time.Second,
	})
	//for {
	conn, _, err := connPool.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(connPool.Stats())
	time.Sleep(1 * time.Second)
	connPool.Put(conn)
	fmt.Println(conn.RemoteAddr())
	//}

	p := &sync.Pool{
		New: func() interface{} {
			conn, _ := net.DialTimeout("tcp", "127.0.0.1:80", 5*time.Second)
			return conn
		},
	}

	a := p.Get().(net.Conn)
	p.Put(a)
	b := p.Get().(net.Conn)
	c := p.Get().(net.Conn)
	fmt.Println(a, b, c)
}
