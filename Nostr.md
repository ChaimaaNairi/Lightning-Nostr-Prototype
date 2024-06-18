# Nostr Protocol - Decentralized Social Networking

The Nostr Protocol is a decentralized social networking solution designed to provide users with a censorship-resistant and privacy-focused platform for communication.

#### Why Nostr Protocol?
Centralized social networking platforms often face issues related to censorship, data privacy, and control over user data. The Nostr Protocol addresses these concerns by decentralizing the communication infrastructure, enabling users to interact directly with each other without relying on centralized servers.

#### How Nostr Protocol Works?
The Nostr Protocol operates by establishing a network of relays that facilitate the distribution of messages between users. 

### Explanation of clients and relays
-
-


**Nostr Documentation:** [link](https://nostr.how/en/what-is-nostr)



## **I. run server**
- To create server port in each node we created `relay_server.go` in each node in `{$node_name}/go-nostr/cmd/relay_server.go` with the port succesively for alice, bob and charlie are 8001, 8002 and 8003. Here are the codes of Alice, Bob and Charlie nodes: [Alice's "relay_server.go"](https://),  [Bob's "relay_server.go"](https://) and [Charlie's "relay_server.go"](https://).

## **II. Generate keys**
- To generate the keys we created `generate_keys.go` in each node in `{$node_name}/go-nostr/cmd/generate_keys.go`. Here are the codes of Alice, Bob and Charlie nodes: [Alice's "generate_keys.go"](https://),  [Bob's "generate_keys.go"](https://) and [Charlie's "generate_keys.go"](https://).

## **III. Connection**
- To make connection between the nodes we created `connection.go` in each node in `{$node_name}/go-nostr/cmd/connection.go`. Here are the codes of Alice, Bob and Charlie nodes: [Alice's "connection.go"](https://),  [Bob's "connection.go"](https://) and [Charlie's "connection.go"](https://).

- main.go

With the relay servers running and keys generated for each node, run the connection setup code to establish connections between the nodes.

## **IV. Send msgs**

- after the connection the nodes can send the msgs to each other, 
evn offline  can receive the msgs

### Alice client:
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


### Bob client:
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


### Charlie client:
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


## single/3 relay

- at first we created `relay_server.go` in `go-nostr/cmd/relay_server.go`

- **Code:** [relay_server.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/go-nostr/cmd/relay_server.go)



