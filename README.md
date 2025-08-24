
# Go WebSocket Chat Server

This is a simple WebSocket-based chat server written in Go. Each client connects with a username and can send messages that are broadcast to all other connected clients.

## Features
- Accepts WebSocket connections on `ws://localhost:8001/v1/ws?username=yourname`
- Clients register with a username via query string
- Messages are broadcast to all connected clients (except the sender)
- Thread-safe connection registry using `sync.Mutex`
- Automatically removes disconnected clients

## Setup & Run
1. Install Go (1.20+ recommended)  
2. Clone or copy the project code  
3. Install dependencies:
   ```bash
   go mod tidy
  ```

4. Run the server:

   ```bash
   go run main.go
   ```
5. The server will start on port `8001`:

   ```
   Server running on port: 8001
   ```

## Connect a Client

* Using [websocat](https://github.com/vi/websocat):

  ```bash
  websocat ws://localhost:8001/v1/ws?username=alice
  websocat ws://localhost:8001/v1/ws?username=bob
  ```

  Messages from Alice will appear for Bob and vice versa.

* Using browser console:

  ```javascript
  let ws = new WebSocket("ws://localhost:8001/v1/ws?username=alice");
  ws.onmessage = (event) => console.log("Received:", event.data);
  ws.onopen = () => ws.send("Hello from Alice!");
  ```

## Project Structure

```
.
├── main.go       # WebSocket server
├── go.mod        # Go module file
└── README.md     # Documentation
```

## Improvements to Try

* Use JSON messages: `{"from": "alice", "msg": "hello"}`
* Add private messaging
* Add JWT authentication
* Save chat history in DB
* Build a simple frontend chat UI
