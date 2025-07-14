package discovery
import (
	"net"
	"time"
	"fmt"
)

func BroadcastDiscoveryMessage() {
	
	broadcastAddr := net.UDPAddr{
		IP:   net.ParseIP("255.255.255.255"),
		Port: 9000,
	}

	conn, err := net.DialUDP("udp4", nil, &broadcastAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := []byte("DISCOVER_REQUEST")

	for {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Broadcast failed:", err)
		}

		time.Sleep(5 * time.Second) // Broadcast every 5 seconds
	}
}