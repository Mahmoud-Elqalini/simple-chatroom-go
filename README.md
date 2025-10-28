# 🗨️ Simple Chatroom (Go RPC)

## 📖 Description
This project implements a **simple multi-client chatroom** using **Go and RPC (Remote Procedure Call)**.  
It allows multiple clients to connect to a central server, exchange messages, and automatically share join/leave notifications.

The chatroom works as a **text-based console application**.

---

## Features

### 🧠 Server
- Handles multiple clients over TCP using RPC.
- Stores all messages in a shared history.
- Broadcasts join and leave notifications automatically.
- Returns the full chat history to each client after sending a message.

### 💬 Client
- Connects to the server via RPC.
- Sends and receives messages in real time.
- Displays the full chat history after each message.
- Supports graceful exit using the `exit` command.
- Shows messages like:
  ```
  Mahmoud has joined the chat.
  Mahmoud has left the chat.
  ```

---

## 🗂️ Folder Structure
```
project/
│
├── server.go        # Chat server logic (handles Join, Leave, SendMessage)
├── client.go        # Chat client logic (connects and communicates via RPC)
├── commons/
│   └── args.go   # Shared data structures (MessageArgs)
├── go.mod           # Go module file
├── .gitignore       # Git ignore configuration
└── README.md        # Project documentation
```

---

## ⚙️ How to Run

### 1️⃣ Start the Server
```bash
go run server.go
```

### 2️⃣ Start the Client (in a new terminal)
```bash
go run client.go
```

### 3️⃣ Interact with the Chat
- Type your message and press **Enter**.
- Type **exit** to leave the chatroom.
- When a user joins or leaves, everyone sees a server message.

---

## 🧩 Example Output

**Client 1 (Mahmoud):**
```
Enter your name: Mahmoud
Welcome, Mahmoud 👋
Your message: Hello everyone!

--- Chat History ---
Mahmoud has joined the chat.
Mahmoud: Hello everyone!
--------------------
Your message: exit
You have left the chat. Goodbye!
```

**Client 2 (Sara):**
```
Enter your name: Sara
Welcome, Sara 👋
Your message: Hi Mahmoud!

--- Chat History ---
Mahmoud has joined the chat.
Mahmoud: Hello everyone!
Sara has joined the chat.
Sara: Hi Mahmoud!
--------------------
```

After Mahmoud leaves:
```
Mahmoud has left the chat.
```

---

## 🧠 Technical Notes
- Uses **net/rpc** and **TCP** for communication.
- Handles each request sequentially.
- Chat messages and system notifications are stored in a slice of strings.
- Compatible with multiple simultaneous clients.

---

## 🎥 Demo Video
🔗 Watch the running application demo here: *([Add your video link here](https://drive.google.com/file/d/1J20QXFc4HrKfLxk7Yg01Jix3WCX6hF42/view?usp=drive_link))*

---

## 👨‍💻 Author
**Mahmoud Hamdi**  
Computer Engineering Student @ Tanta University  
Passionate about software engineering, data science, and AI.

---

