package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"chatroom/commons"
)

// Listener: Represents the chat server and stores all messages and connected clients
type Listener struct {
	mu       sync.Mutex
	Messages []string        // To store all chat messages
	Clients  map[string]bool // To store active clients
}

// SendMessage receives a message from a client, adds it to the history, and returns all messages
func (l *Listener) SendMessage(args commons.MessageArgs, reply *[]string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Format and store the message
	msg := fmt.Sprintf("%s: %s", args.Name, args.Text)
	l.Messages = append(l.Messages, msg)

	// Copy chat history
	history := make([]string, len(l.Messages))
	copy(history, l.Messages)

	// Print message on server console
	fmt.Println(msg)

	// Return full chat history
	*reply = history
	return nil
}

// Join: Adds a system message when a client joins
func (l *Listener) Join(name string, reply *bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.Clients[name] { // Only if not already joined
		l.Clients[name] = true
		systemMsg := fmt.Sprintf("[Server]: %s has joined the chat.", name)
		l.Messages = append(l.Messages, systemMsg)
		fmt.Println(systemMsg)
	}

	*reply = true
	return nil
}

// Leave: Removes a client and notifies others
func (l *Listener) Leave(name string, reply *bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.Clients[name] {
		delete(l.Clients, name)
		leaveMsg := fmt.Sprintf("[Server]: %s has left the chat.", name)
		l.Messages = append(l.Messages, leaveMsg)
		fmt.Println(leaveMsg)
	}

	*reply = true
	return nil
}

func main() {
	// Define the TCP address and port for the server
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatalf("Error resolving TCP address: %v", err)
	}

	// Start listening for TCP connections
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Error starting TCP listener: %v", err)
	}
	defer listener.Close()

	// Create and register the RPC server
	server := &Listener{
		Clients: make(map[string]bool), // Initialize Clients map ✅
	}
	err = rpc.Register(server)
	if err != nil {
		log.Fatalf("Error registering RPC server: %v", err)
	}

	fmt.Println("Chat server is running on port 42586")
	fmt.Println("Waiting for clients to connect\n")

	// Accept and handle incoming client connections
	rpc.Accept(listener)
}
