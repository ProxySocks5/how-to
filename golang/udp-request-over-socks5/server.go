package main

import (
	"fmt"
	"net"
)

func main() {
	// Define the UDP server address and port
	addr := ":8080" // Use any desired port number
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP socket
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error creating UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("UDP server listening on %s\n", addr)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	// Continuously listen for incoming connections
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			continue
		}

		// Print client IP and the received message
		fmt.Printf("Received message from %s: %s\n", clientAddr.String(), string(buffer[:n]))

		// Optionally respond to the client
		_, err = conn.WriteToUDP([]byte("Message received!"), clientAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
