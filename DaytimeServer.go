package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	buf := bytes.NewBuffer([]byte{})
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		buf.Reset()
		buf.Write([]byte(daytime))
		buf.Write([]byte("\n"))

		conn.Write(buf.Bytes()) // don't care about return value
		conn.Close()            // we're finished with this client
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
