package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func serviceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func StartNetworkDaemonwithPool() *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(1)
	p := serviceConnCache()
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer func(server net.Listener) {
			err := server.Close()
			if err != nil {
				log.Fatalf("%v", err)
			}
		}(server)

		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatalf("%v", err)
			}
			svcConn := p.Get()
			// do whatever you want to do with this conn
			_, err = fmt.Fprintln(conn, "")
			if err != nil {
				return
			}
			p.Put(svcConn)
			err = conn.Close()
			if err != nil {
				return
			}
		}
	}()
	return &wg
}
