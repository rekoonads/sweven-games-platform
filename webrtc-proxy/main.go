package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sweven-games/webrtc-proxy/proxy"
)

func main() {
	fmt.Println("Sweven Games WebRTC Proxy")
	fmt.Println("==========================")

	// Configuration from environment
	stunServer := os.Getenv("STUN_SERVER")
	if stunServer == "" {
		stunServer = "stun:stun.l.google.com:19302"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config := &proxy.ProxyConfig{
		StunServer: stunServer,
		Port:       port,
	}

	log.Printf("Starting WebRTC Proxy on port %s", port)
	log.Printf("Using STUN server: %s", stunServer)

	// Initialize and start proxy
	p := proxy.NewProxy(config)
	if err := p.Start(); err != nil {
		log.Fatalf("Failed to start proxy: %v", err)
	}

	// Keep running
	select {}
}
