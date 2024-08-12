package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	adress := "localhost:" + os.Args[1]
	addr, err := net.ResolveUDPAddr("udp", adress)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server listening on", addr)
	fmt.Println(addr)

	for {
		buffer := make([]byte, 1024)

		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			return
		}

		fmt.Printf("Received message from %s: %s\n", clientAddr, string(buffer[:n]))

		_, err = conn.WriteToUDP([]byte("Message received"), clientAddr)
		if err != nil {
			fmt.Println("Error writing to UDP connection:", err)
			return
		}
	}
}
