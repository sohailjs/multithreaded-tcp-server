package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	address string
	ln      net.Listener
	quitCh  chan struct{}
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
		quitCh:  make(chan struct{}),
	}
}
func (server *Server) start() error {
	ln, err := net.Listen("tcp", server.address)
	if err != nil {
		return err
	}
	server.ln = ln

	//accept connections
	server.Accept()
	<-server.quitCh
	return nil
}

func (server *Server) Accept() {
	for {
		conn, err := server.ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		fmt.Printf("connected client: %s", conn.RemoteAddr())
		go server.Read(conn)
	}
}

func (server *Server) Read(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("error reading from connection")
		}
		fmt.Printf("\nmessage from client %s: %s", conn.RemoteAddr(), string(buf))
	}
}

func main() {
	server := NewServer(":3000")
	err := server.start()
	if err != nil {
		log.Fatal(err)
	}
}
