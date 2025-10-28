package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"chatroom/commons"
)

func main() {
	// Connect to the server
	client, err := rpc.Dial("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer client.Close()

	in := bufio.NewReader(os.Stdin)

	// Ask for username
	fmt.Print("Enter your name: ")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Anonymous"
	}

	// Tell server that user joined
	var ack bool
	err = client.Call("Listener.Join", name, &ack)
	if err != nil {
		log.Fatalf("Error joining chat: %v", err)
	}

	fmt.Printf("\nWelcome, %s \n", name)
	fmt.Println("Type your message and press Enter.")
	fmt.Println("Type 'exit' to leave the chat.\n")

	for {
		fmt.Print("Your message: ")
		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		line = strings.TrimSpace(line)

		if strings.EqualFold(line, "exit") {
			// Notify the server that this user left
			err = client.Call("Listener.Leave", name, &ack)
			if err != nil {
				log.Printf("Error notifying leave: %v", err)
			}
			fmt.Println("You have left the chat. Goodbye")
			break
		}

		if line == "" {
			continue
		}

		// Send name + message and receive full chat history
		var history []string
		args := commons.MessageArgs{Name: name, Text: line}
		err = client.Call("Listener.SendMessage", args, &history)
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}

		// Print chat history nicely
		fmt.Println("\n--- Chat History ---")
		for _, msg := range history {
			fmt.Println(msg)
		}
		fmt.Println("--------------------\n")
	}
}
