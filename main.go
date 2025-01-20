package main

import (
	"fmt"
	"log"
	"net"
)

func readMessage(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Printf("error reading from connection")
	}
	fmt.Printf("received data: %s", string(buf))
	conn.Write([]byte("hello from server"))
	//conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("connected client: %s\n", conn.RemoteAddr().String())
		go readMessage(conn)
	}
}
