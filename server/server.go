package server

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"math/rand"
	"net"
	"tcp_pow/pow"
)

var (
	quotes = []string{
		"The best and most beautiful things in the world cannot be seen or even touched - they must be felt with the heart.",
		"It is during our darkest moments that we must focus to see the light.",
		"The only way to do great work is to love what you do.",
		"The best way out is always through.",
		"The only limit to our realization of tomorrow will be our doubts of today.",
	}
)

func HandleConnection(conn net.Conn) {
	client := conn.RemoteAddr()
	server := conn.LocalAddr()
	connection := pow.NewConnection(client.String(), server.String())
	newBlock, err := pow.CreateBlock(*connection)

	if err != nil {
		conn.Write([]byte(fmt.Sprintf("Invalid proof of work from %s. Close connection by reason %s\n",
			client.String(), err.Error())))
		conn.Close()
		return
	} else {
		log.Println("Connection confirmed by Pow:")
		spew.Dump(newBlock)
	}

	// serve quote or disconnect if pow invalid
	conn.Write([]byte(quotes[rand.Intn(len(quotes))]))
}
