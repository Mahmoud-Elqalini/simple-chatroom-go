package main

import (
	"bufio"
	"fmt"
	"log"
	rpc "net/rpc"
	"os"
	"strings"
)

func main() {
	// Connecting to the server
	client, err := rpc.Dial("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==================== Chat Room ====================")
		fmt.Print("Write your message (type 'exit' to quit): ")
		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Remove spaces and newline characters
		line = strings.TrimSpace(line)

		if line == "exit" {
			fmt.Println("Exiting chat")
			break
		}

		// Sending the message and receiving all messages
		var history []string
		err = client.Call("Listener.SendMessage", line, &history)
		if err != nil {
			log.Fatal(err)
		}

		// Print all messages
		fmt.Println("===== Chat History =====")
		for _, msg := range history {
			fmt.Println(msg)
		}
		fmt.Println("========================")
	}
}
