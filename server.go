package main

import (
	"fmt"
	"log"
	"net"
	rpc "net/rpc"
)

// The Listener struct represents the server
type Listener struct {
	Messages []string // To store all messages
}

// SendMessage receives a message from the client and adds it to the history
// reply returns all messages after addition

func (l *Listener) SendMessage(msg string, reply *[]string) error {
	l.Messages = append(l.Messages, msg) // Save the new message
	*reply = l.Messages                  // Return all messages to the client
	fmt.Println("New message received:", msg)
	return nil
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	// Starting to listen on TCP
	inbound, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer inbound.Close()

	// Create and register the server
	listener := new(Listener)
	rpc.Register(listener)

	fmt.Println("Chat server is running on port 42586")
	rpc.Accept(inbound) 
}
