package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	proxyAddr := "127.0.0.1:1080"  // SOCKS5 proxy address
	proxyUser := "your_username"   // Replace with your username
	proxyPass := "your_password"   // Replace with your password
	targetAddr := "127.0.0.1:9999" // UDP target server (replace with actual UDP server)

	// Establish TCP connection to SOCKS5 proxy
	tcpConn, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		log.Fatalf("Failed to connect to SOCKS5 proxy: %v", err)
	}
	defer tcpConn.Close()

	// Perform SOCKS5 authentication (Username/Password)
	_, err = tcpConn.Write([]byte{0x05, 0x01, 0x02})
	if err != nil {
		log.Fatalf("Failed to send authentication request: %v", err)
	}

	resp := make([]byte, 2)
	_, err = tcpConn.Read(resp)
	if err != nil || resp[0] != 0x05 || resp[1] != 0x02 {
		log.Fatalf("SOCKS5 proxy requires authentication or failed to accept auth method")
	}

	// Send username/password authentication
	authRequest := append([]byte{0x01, byte(len(proxyUser))}, proxyUser...)
	authRequest = append(authRequest, byte(len(proxyPass)))
	authRequest = append(authRequest, proxyPass...)

	_, err = tcpConn.Write(authRequest)
	if err != nil {
		log.Fatalf("Failed to send authentication credentials: %v", err)
	}

	_, err = tcpConn.Read(resp)
	if err != nil || resp[0] != 0x01 || resp[1] != 0x00 {
		log.Fatalf("SOCKS5 authentication failed")
	}

	fmt.Println("Authentication successful")

	// Request UDP Associate
	request := []byte{
		0x05, 0x03, 0x00, 0x01, // SOCKS5, UDP_ASSOCIATE, Reserved, IPv4
		0x00, 0x00, 0x00, 0x00, // Bind Address (0.0.0.0)
		0x00, 0x00, // Bind Port (0)
	}
	_, err = tcpConn.Write(request)
	if err != nil {
		log.Fatalf("Failed to send UDP associate request: %v", err)
	}

	resp = make([]byte, 10)
	_, err = tcpConn.Read(resp)
	if err != nil {
		log.Fatalf("Failed to read UDP associate response: %v", err)
	}

	udpRelayAddr := fmt.Sprintf("%d.%d.%d.%d:%d",
		resp[4], resp[5], resp[6], resp[7], binary.BigEndian.Uint16(resp[8:10]))

	fmt.Printf("UDP Relay Address: %s\n", udpRelayAddr)

	// Send UDP packet via the UDP relay
	udpConn, err := net.Dial("udp", udpRelayAddr)
	if err != nil {
		log.Fatalf("Failed to establish UDP connection: %v", err)
	}
	defer udpConn.Close()

	// SOCKS5 UDP header
	var udpHeader bytes.Buffer
	udpHeader.WriteByte(0x00) // Reserved
	udpHeader.WriteByte(0x00) // Reserved
	udpHeader.WriteByte(0x00) // Fragment (0 = no fragmentation)
	udpHeader.WriteByte(0x01) // Address type (IPv4)

	host, port, err := net.SplitHostPort(targetAddr)
	if err != nil {
		log.Fatalf("Invalid target address: %v", err)
	}
	targetIP := net.ParseIP(host).To4()
	targetPort, err := net.LookupPort("udp", port)
	if err != nil {
		log.Fatalf("Invalid target port: %v", err)
	}

	udpHeader.Write(targetIP)
	binary.Write(&udpHeader, binary.BigEndian, uint16(targetPort))

	// UDP Message
	message := []byte("hello udp server")
	packet := append(udpHeader.Bytes(), message...)

	// Send the UDP message
	_, err = udpConn.Write(packet)
	if err != nil {
		log.Fatalf("Failed to send UDP packet: %v", err)
	}

	fmt.Println("Message sent: hello udp server")

	// Receive response
	buffer := make([]byte, 512)
	udpConn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, err := udpConn.Read(buffer)
	if err != nil {
		log.Fatalf("Failed to receive UDP response: %v", err)
	}

	fmt.Printf("Received %d bytes: %s\n", n, string(buffer[:n]))
}
