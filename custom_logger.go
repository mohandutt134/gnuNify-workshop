package main

import (
	"log"
	"net"
	"os"
	"time"
)

type tcpLogger struct {
	conn net.Conn
}

func NewTcpWriter() *tcpLogger {
	t := tcpLogger{}
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Printf("Error while creating %v", err)
		return &t
	}

	t.conn = conn
	return &t
}

func (t *tcpLogger) Write(b []byte) (n int, err error) {
	if t.conn != nil {
		t.conn.Write(b)
	}

	return os.Stderr.Write(b)
}

func main() {
	timer := time.NewTicker(time.Second * 1)
	cLog := log.New(NewTcpWriter(), "", 0)

	for {
		t := <-timer.C
		cLog.Println(t)
	}
}
