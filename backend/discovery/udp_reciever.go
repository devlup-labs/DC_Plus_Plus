package discovery

import (
	"fmt"
	"net"
	"github.com/joho/godotenv"
)

func StartUDPListener(port string) {
	if err := godotenv.Load(".env"); err != nil {
        fmt.Printf("No .env file found, using system environment variables.")
    }
	
	addr := net.UDPAddr{
		Port: 9000, 
		IP:   net.ParseIP("0.0.0.0"), 
	}

	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	fmt.Println("Listening for UDP broadcasts on port", port)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Received from %s: %s\n", remoteAddr, message)

	}
}