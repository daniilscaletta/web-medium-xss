package main

import (
	"example/v3/bot/handlers"
	"fmt"
	"net"
	"os"
)

func main() {
	// Убедитесь, что сервер запущен на нужном порту
	port := "12345"
	fmt.Println("Starting server on port:", port)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on port:", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handlers.HandlerConnection(conn)
	}
}
