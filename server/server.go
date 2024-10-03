package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

var clients = make(map[net.Conn]bool)
var mu sync.Mutex

func broadcastMessage(message string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message))
			if err != nil {
				log.Println("Error sending message to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func StartServer(port string) {
	listerner, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)

	}
	defer listerner.Close()

	fmt.Println("Server started on port", port)

	for {
		conn, err := listerner.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err)
			continue
		}

		mu.Lock()
		clients[conn] = true
		mu.Unlock()

		fmt.Println("New client connected:", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		fmt.Println("Received message:", message)

		broadcastMessage(message, conn)
	}
}
func main() {
	port := "8080"
	StartServer(port)
}
