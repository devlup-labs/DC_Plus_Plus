package main

import (
	"DC_PLUS_PLUS/discovery"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
        fmt.Printf("No .env file found, using system environment variables.")
    }
	go discovery.StartUDPListener(os.Getenv("UDP_PORT"))
	go discovery.BroadcastDiscoveryMessage()

	select {}
}
