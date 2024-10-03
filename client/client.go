package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func ConnectToSever(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Failed to connect to server at %s: %v", address)
	}
	defer conn.Close()

	fmt.Println("Connected to server:", address)

	// Read messages from the user and send them to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a message:")
		message, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Println("Failed to dsend message:", err)
			continue
		}

		reply, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading server response: ", err)
			return
		}
		fmt.Println("Server response: ", reply)
	}
}
func main() {
	serverAddress := "localhost: 8080"
	ConnectToSever(serverAddress)
}
