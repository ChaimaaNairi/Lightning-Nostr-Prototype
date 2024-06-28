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
	"github.com/ChaimaaNairi/go-nostr"
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

func loadKeys(keysFile string) (string, string, error) {
	file, err := os.Open(keysFile)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	var keys struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	}
	if err := json.NewDecoder(file).Decode(&keys); err != nil {
		return "", "", err
	}

	return keys.PublicKey, keys.PrivateKey, nil
}

func generateNewKeys() (string, string) {
	privateKey := nostr.GeneratePrivateKey()
	publicKey, _ := nostr.GetPublicKey(privateKey)
	return publicKey, privateKey
}

func saveKeys(publicKey, privateKey, keysFile string) error {
	keyData := struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	}{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	jsonData, err := json.MarshalIndent(keyData, "", "    ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(keysFile, jsonData, 0644); err != nil {
		return err
	}

	return nil
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

	// WebSocket connection to relay server
	relayServer := "localhost:8000"
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+relayServer+"/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to relay server:", err)
	}
	defer conn.Close()

	// Start receiving messages from relay server
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
				fmt.Printf("Message received: %s from %s to %s\n", message.Content, message.Sender, "Charlie")
			}
		}
	}()

	// Start sending messages to other nodes via relay server
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

		message := Message{
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
			Sender:    publicKey,
			Recipient: recipientPublicKey,
			Content:   text,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		sendMessage(conn, message)
		logMessage(message, "messages_sent.json")
		fmt.Printf("Message sent: %s from Charlie to %s\n", message.Content, recipientPublicKey)
	}

	// Gracefully shutdown on interrupt signal (Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down...")
	conn.Close()
}
