# Nostr Protocol - Decentralized Social Networking

The Nostr Protocol is a decentralized social networking solution designed to provide users with a censorship-resistant and privacy-focused platform for communication.

#### Why Nostr Protocol?
Centralized social networking platforms often face issues related to censorship, data privacy, and control over user data. The Nostr Protocol addresses these concerns by decentralizing the communication infrastructure, enabling users to interact directly with each other without relying on centralized servers.

#### How Nostr Protocol Works?
The Nostr Protocol operates by establishing a network of relays that facilitate the distribution of messages between users. Here's an overview of its functioning:

1. **Relay Network**: The Nostr Protocol relies on a network of relays distributed across the internet. These relays serve as intermediaries for message propagation, ensuring that messages reach their intended recipients.

2. **Message Distribution**: Users post messages, also known as "notes," to the network of relays. These relays propagate the messages to other relays and clients connected to the network.

3. **Decentralized Identity**: Users can create and manage their identities without the need for a central authority. This allows for greater control over personal data and enhances user privacy.

4. **Censorship Resistance**: By distributing messages across a decentralized network, the Nostr Protocol makes it difficult for any single entity to censor communication. This ensures that users can express themselves freely without fear of censorship.

5. **Privacy**: Users' interactions on the Nostr network are private and secure, as messages are distributed in a peer-to-peer fashion without passing through centralized servers.


### Explanation of clients and relays
-
-
-
-



------------------

- To generate the keys we created `generate_keys.go` in each node in `{$node_name}/go-nostr/cmd/generate_keys.go`  :
```bash
package main

import (
    "fmt"
    "github.com/nbd-wtf/go-nostr"
    "github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
    sk := nostr.GeneratePrivateKey()
    pk, _ := nostr.GetPublicKey(sk)
    nsec, _ := nip19.EncodePrivateKey(sk)
    npub, _ := nip19.EncodePublicKey(pk)

    fmt.Println("sk:", sk)
    fmt.Println("pk:", pk)
    fmt.Println(nsec)
    fmt.Println(npub)
}

```

------------------

- To create server port in each node we created `relay_server.go` in each node in `{$node_name}/go-nostr/cmd/relay_server.go` we change the port, succesively for alice, bob and charlie are 3001, 3002 and 3003 :
```bash
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/websocket"
)

// Define a struct to represent the Nostr relay server.
type NostrRelayServer struct {
    clients       map[string]*websocket.Conn
    addClient     chan *websocket.Conn
    removeClient  chan *websocket.Conn
    broadcast     chan []byte
    upgrader      websocket.Upgrader
}

// Initialize the Nostr relay server.
func NewNostrRelayServer() *NostrRelayServer {
    return &NostrRelayServer{
        clients:       make(map[string]*websocket.Conn),
        addClient:     make(chan *websocket.Conn),
        removeClient:  make(chan *websocket.Conn),
        broadcast:     make(chan []byte),
        upgrader:      websocket.Upgrader{},
    }
}

// Start the Nostr relay server.
func (s *NostrRelayServer) Start() {
    // Start a goroutine to handle incoming client connections.
    go func() {
        for {
            select {
            case conn := <-s.addClient:
                s.clients[conn.RemoteAddr().String()] = conn
            case conn := <-s.removeClient:
                delete(s.clients, conn.RemoteAddr().String())
                conn.Close()
            case msg := <-s.broadcast:
                // Broadcast the received message to all connected clients.
                for _, conn := range s.clients {
                    conn.WriteMessage(websocket.TextMessage, msg)
                }
            }
        }
    }()

    // Start a goroutine to handle HTTP requests to upgrade to WebSocket connections.
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := s.upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println("Error upgrading to WebSocket:", err)
            return
        }
        defer conn.Close()

        // Add the client to the relay server.
        s.addClient <- conn
        defer func() {
            s.removeClient <- conn
        }()

        // Listen for messages from the client.
        for {
            _, msg, err := conn.ReadMessage()
            if err != nil {
                break
            }
            // Broadcast the received message to all connected clients.
            s.broadcast <- msg
        }
    })

    // Start the HTTP server.
    go func() {
        log.Println("Starting Nostr Relay Server for Charlie on port 8003...")
        if err := http.ListenAndServe(":8003", nil); err != nil {
            log.Fatal("Error starting Nostr Relay Server:", err)
        }
    }()

    // Wait for interrupt signal (Ctrl+C) to gracefully shutdown the server.
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    log.Println("Shutting down Nostr Relay Server...")
}

func main() {
    // Create a new instance of the Nostr relay server.
    server := NewNostrRelayServer()

    // Start the Nostr relay server.
    server.Start()
}


```


-----------------

- **sk:** stands for the generated private key.
- **pk:** stands for the corresponding public key.
- **nsec:** is the encoded private key.
- **npub:** is the encoded public key.



------------

## Alice client:
- To clone the repo & build client:
```bash
Alice$ git clone https://github.com/nbd-wtf/go-nostr

Alice$ cd go-nostr
Alice$ go build
```
<img src="images/Capture d’écran -2.png" alt="">

- To generate Alice's keys:
```bash
Alice$ go run generate_keys.go
```
<img src="images/Capture d’écran -3.png" alt="">

- **Code:** [generate_keys.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/alice/go-nostr/cmd/generate_keys.go)


## Bob client:
- To clone the repo & build client:
```bash
Bob$ git clone https://github.com/nbd-wtf/go-nostr

Bob$ cd go-nostr
Bob$ go build
```
<img src="images/Capture d’écran -4.png" alt="">

- To generate Bob's keys:
```bash
Bob$ go run generate_keys.go
```
<img src="images/Capture d’écran -5.png" alt="">

- **Code:** [generate_keys.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/bob/go-nostr/cmd/generate_keys.go)


## Charlie client:
- To clone the repo & build client:
```bash
Charlie$ git clone https://github.com/nbd-wtf/go-nostr

Charlie$ cd go-nostr
Charlie$ go build
```
<img src="images/Capture d’écran .png" alt="">


- To generate Charlie's keys:
```bash
Charlie$ go run generate_keys.go
```
<img src="images/Capture d’écran -1.png" alt="">

- **Code:** [generate_keys.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/charlie/go-nostr/cmd/generate_keys.go)


## single relay

- at first we created `relay_server.go` in `go-nostr/cmd/relay_server.go` then we wrote the code 
``` bash
package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "github.com/gorilla/websocket"
)

// Define a struct to represent the Nostr relay server.
type NostrRelayServer struct {
	clients       map[string]*websocket.Conn
	addClient     chan *websocket.Conn
	removeClient  chan *websocket.Conn
	broadcast     chan []byte
	upgrader      websocket.Upgrader
}

// Initialize the Nostr relay server.
func NewNostrRelayServer() *NostrRelayServer {
	return &NostrRelayServer{
		clients:       make(map[string]*websocket.Conn),
		addClient:     make(chan *websocket.Conn),
		removeClient:  make(chan *websocket.Conn),
		broadcast:     make(chan []byte),
		upgrader:      websocket.Upgrader{},
	}
}

// Start the Nostr relay server.
func (s *NostrRelayServer) Start() {
	// Start a goroutine to handle incoming client connections.
	go func() {
		for {
			select {
			case conn := <-s.addClient:
				s.clients[conn.RemoteAddr().String()] = conn
			case conn := <-s.removeClient:
				delete(s.clients, conn.RemoteAddr().String())
				conn.Close()
			case msg := <-s.broadcast:
				// Broadcast the received message to all connected clients.
				for _, conn := range s.clients {
					conn.WriteMessage(websocket.TextMessage, msg)
				}
			}
		}
	}()

	// Start a goroutine to handle HTTP requests to upgrade to WebSocket connections.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			return
		}
		defer conn.Close()

		// Add the client to the relay server.
		s.addClient <- conn
		defer func() {
			s.removeClient <- conn
		}()

		// Listen for messages from the client.
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			// Broadcast the received message to all connected clients.
			s.broadcast <- msg
		}
	})

	// Start the HTTP server.
	go func() {
		log.Println("Starting Nostr Relay Server on port 8080...")
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal("Error starting Nostr Relay Server:", err)
		}
	}()

	// Wait for interrupt signal (Ctrl+C) to gracefully shutdown the server.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down Nostr Relay Server...")
}

func main() {
	// Create a new instance of the Nostr relay server.
	server := NewNostrRelayServer()

	// Start the Nostr relay server.
	server.Start()
}
```

- To run the server:
```bash
project$ go run relay_server.go
```
<img src="images/Capture d’écran -6.png" alt="">


- **Code:** [relay_server.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/go-nostr/cmd/relay_server.go)
