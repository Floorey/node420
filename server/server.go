package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		fmt.Print("Received message :", message)

		_, err = conn.Write([]byte("Message received\n"))
		if err != nil {
			log.Println("Error sending response:", err)

		}
	}
}
func main() {
	port := "8080"
	StartServer(port)
}
