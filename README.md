# Simple Chatroom (Go RPC)

## Description
This project implements a **simple chatroom** using **Go and RPC (Remote Procedure Call)**.  
It allows multiple clients to send messages to a coordinating server and fetch the **chat history**.  

The chatroom works as a **text-based console application**.

---

## Features
- **Client**
  - Connects to the server via RPC.
  - Sends messages to the server.
  - Receives the full chat history after sending each message.
  - Runs in a loop until the user types `"exit"` or presses `Ctrl+C`.
  - Reads full lines including spaces.

- **Server**
  - Stores all messages in a list.
  - Returns the full chat history to clients upon receiving a message.
  - Handles multiple clients sequentially using TCP and RPC.

---

## Folder Structure
```
project/
│
├── server.go        # Chat server code
├── client.go        # Chat client code
```

---

## How to Run

1. **Start the Server**
```bash
go run server.go
```
2. **Start the Client**
```bash
go run client.go
```
3. **Use the Chat**
- Type a message and press Enter.
- Type `exit` to quit the chat client.

---

## Notes
- The server must be running before starting the client.
- Multiple clients can connect and send messages.
- Chat history is displayed after sending each message.
- The client uses `bufio.Reader` to read full lines including spaces.

## Demo Video

- Watch the running application here:
