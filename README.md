# 🚀 Go_gRPC

[![Go Version](https://img.shields.io/badge/Go-1.20+-blue.svg)](https://golang.org)
[![gRPC](https://img.shields.io/badge/gRPC-Enabled-green.svg)](https://grpc.io/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

---

## ✨ Overview

**Go_gRPC** is a complete demonstration of building modern gRPC services in Go.  
It includes all four types of gRPC calls:

- 🔹 **Unary RPC** — single request, single response.
- 🔸 **Server Streaming RPC** — single request, multiple responses.
- 🔸 **Client Streaming RPC** — multiple requests, single response.
- 🔹 **Bidirectional Streaming RPC** — multiple requests, multiple responses simultaneously.

---

## 💡 Motivation

> "I want to understand, not to be understood."  

This project was built to **master gRPC concepts in Go** by exploring each RPC type in detail, with clean modular code and full explanation.


---

## ⚙️ Setup & Installation

### 1️⃣ Clone the repository

```bash
git clone https://github.com/Girish070/Go_gRPC.git
cd Go_gRPC
2️⃣ Install dependencies
go mod tidy
3️⃣ Generate protobuf code
protoc --go_out=. --go-grpc_out=. proto/great.proto
💻 Running the Project
Start server
cd server
go run *.go
Run client
cd client
go run *.go
⚡ Switch between RPC types
In client/main.go, choose which RPC call to execute:
callSayHello(client)
callSayHelloServerStream(client, names)
callSayHelloClientStream(client, names)
callHelloBidirectionalStream(client, names)
Uncomment the desired call and comment the others.

🖨️ Example Output
✅ Unary
pgsql
2025/07/11 Hello from the server! 👋
📤 Server Streaming
2025/07/11 Hello Girish
2025/07/11 Hello John
...
📥 Client Streaming
vbnet
2025/07/11 Sent request with name: Girish
2025/07/11 Sent request with name: John
...
2025/07/11 Server responded with: Hello to everyone!
🔄 Bidirectional Streaming
2025/07/11 Sending: Girish
2025/07/11 Received: Hello Girish
...
🛡️ Requirements
Go 1.20+

Protocol Buffers compiler (protoc)

protoc-gen-go and protoc-gen-go-grpc plugins

💬 Contributing
Pull requests are welcome! Feel free to open an issue or discuss improvements.

⭐ License
MIT License

✉️ Contact
Made with ❤️ by Girish

🎯 Quotes for Devs
"Understand deeply, code simply."

"Code is like humor. When you have to explain it, it’s bad. But here, we explain to understand, not to show off."
