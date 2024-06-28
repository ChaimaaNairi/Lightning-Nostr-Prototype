package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"../go-nostr/cmd"
)

type Message struct {
	ID        string `json:"id"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

var mu sync.Mutex

func logMessage(message Message, logFile string) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(message); err != nil {
		log.Println("Error logging message:", err)
	}
}

func sendMessage(conn *websocket.Conn, message Message) {
	err := conn.WriteJSON(message)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func main() {
	keysFile := "key_charlie.json"
	publicKey, privateKey, err := loadKeys(keysFile)
	if err != nil || publicKey == "" {
		publicKey, privateKey = generateNewKeys()
		if err := saveKeys(publicKey, privateKey, keysFile); err != nil {
			log.Fatal("Error saving keys:", err)
		}
	}

	fmt.Println("Charlie's private key (sk):", privateKey)
	fmt.Println("Charlie's public key (pk):", publicKey)

	nameMap := map[string]string{
		publicKey: "Charlie",
		"72b76dd37a7fa2b9cf48081940fac13deb0d10858dd9887c7e8e0726867dbe6c": "Alice",
		"fab28a32bc209bfee51d13ab724d9e18edf40216147df79d50129886fef73759": "Bob",
	}

	relayServer := "localhost:8000"

	conn, _, err := websocket.DefaultDialer.Dial("ws://"+relayServer+"/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to relay server:", err)
	}
	defer conn.Close()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			var message Message
			if err := json.Unmarshal(msg, &message); err != nil {
				log.Println("Error unmarshalling message:", err)
				continue
			}

			if message.Recipient == publicKey {
				logMessage(message, "messages_received.json")
				fmt.Printf("Message received: %s from %s to %s\n", message.Content, nameMap[message.Sender], nameMap[message.Recipient])
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter messages to send. Type 'exit' to quit.")
	for {
		fmt.Print("Message: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}

		fmt.Print("Recipient's public key: ")
		recipientPublicKey, _ := reader.ReadString('\n')
		recipientPublicKey = strings.TrimSpace(recipientPublicKey)

		if _, ok := nameMap[recipientPublicKey]; !ok {
			fmt.Print("Enter the name of recipient: ")
			recipientName, _ := reader.ReadString('\n')
			recipientName = strings.TrimSpace(recipientName)
			nameMap[recipientPublicKey] = recipientName
		}

		message := Message{
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
			Sender:    publicKey,
			Recipient: recipientPublicKey,
			Content:   text,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		sendMessage(conn, message)
		logMessage(message, "messages_sent.json")
		fmt.Printf("Message sent: %s from %s to %s\n", message.Content, nameMap[message.Sender], nameMap[message.Recipient])
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down...")
	conn.Close()
}
