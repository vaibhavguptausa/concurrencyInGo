package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var benchFlag string

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func StartNetworkDaemon() *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(1)
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
			connectToService()
			_, err = fmt.Fprintln(conn, "")
			if err != nil {
				return
			}
			err = conn.Close()
			if err != nil {
				return
			}
		}
	}()
	return &wg
}
