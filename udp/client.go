package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	adress := "localhost:" + os.Args[1]
	serverAddr, err := net.ResolveUDPAddr("udp", adress)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error dialing UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Enter a message to send:")
	var message string
	fmt.Scanln(&message)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Server response: %s\n", string(buffer[:n]))
}
