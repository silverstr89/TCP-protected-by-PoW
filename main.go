package main

import (
	"net"
	"tcp_pow/pow"
	"tcp_pow/server"
)

func main() {
	pow.GenesisBlock()
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go server.HandleConnection(conn)
	}
}
