package main

import (
	"io"
	"log"
	"net"
	"testing"
)

func init() {
	daemonWg := StartNetworkDaemonwithPool()
	daemonWg.Wait()
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("%v", err)
		}
		if _, err := io.ReadAll(conn); err != nil {
			log.Fatalf("%v", err)
		}
		err = conn.Close()
		if err != nil {
			return
		}
	}
}
