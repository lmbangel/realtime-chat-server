# Go WebSocket Chat Server

A simple WebSocket-based chat server written in Go.  
Clients connect with a username and can exchange messages in **broadcast** mode or send **private messages** to a specific user.

## Features
- Accepts WebSocket connections at:  
  `ws://localhost:8001/v1/ws?username=yourname`
- Registers clients with a username via query string
- Handles messages in **JSON format**:

```json
{
   "message": "Hello everyone!",
   "target": ""   // leave empty for broadcast, or set to a username for private message
}
```

* Broadcasts messages to all connected clients (except the sender)
* Supports **private messaging** when `target` is set
* Thread-safe connection registry using `sync.Mutex`
* Automatically removes disconnected clients

## Setup & Run

1. Install Go (1.20+ recommended)

2. Clone or copy this project

3. Install dependencies:

```
go mod tidy
```

4. Run the server:

```
go run main.go
```

5. The server will start on port `8001`:

```
Server running on port: 8001
```

## Connect a Client

### Using [websocat](https://github.com/vi/websocat):

```bash
# Connect as Alice
websocat ws://localhost:8001/v1/ws?username=alice

# Connect as Bob
websocat ws://localhost:8001/v1/ws?username=bob
```

Send a JSON message:

```json
{"message":"Hello Bob!","target":"bob"}
```

* If `target` is empty (`""`), the message is broadcast to everyone.
* If `target` is a username, only that user receives the message.


### Using Browser Console:

```javascript
let ws = new WebSocket("ws://localhost:8001/v1/ws?username=alice");

ws.onmessage = (event) => console.log("Received:", event.data);

ws.onopen = () => {
  // Send broadcast
  ws.send(JSON.stringify({message: "Hello all!", target: ""}));

  // Send private message
  ws.send(JSON.stringify({message: "Hi Bob!", target: "bob"}));
};
```


## Project Structure

```
.
├── main.go       # WebSocket server with JSON + private messaging
├── go.mod        # Go module file
└── README.md     # Documentation
```


## Improvements to Try Next

* Add authentication (JWT-based user verification)
* Store chat history in a database
* Handle message delivery receipts or acknowledgements
* Build a frontend chat UI (React, Vue, or plain JS)
* Implement rooms or channels for group chats