package proxy

import (
	"fmt"
	"log"
	"sync"

	"github.com/pion/webrtc/v3"
)

// ProxyConfig holds configuration for the WebRTC proxy
type ProxyConfig struct {
	StunServer string
	TurnServer string
	Port       string
}

// Proxy represents the WebRTC proxy service
type Proxy struct {
	config      *ProxyConfig
	connections map[string]*webrtc.PeerConnection
	mutex       sync.RWMutex
}

// NewProxy creates a new WebRTC proxy instance
func NewProxy(config *ProxyConfig) *Proxy {
	return &Proxy{
		config:      config,
		connections: make(map[string]*webrtc.PeerConnection),
	}
}

// Start initializes and starts the proxy service
func (p *Proxy) Start() error {
	log.Println("WebRTC Proxy initialized")
	log.Printf("Configuration: STUN=%s, Port=%s", p.config.StunServer, p.config.Port)

	// TODO: Initialize WebRTC API and setup listeners
	// This is a placeholder for the actual implementation

	return nil
}

// CreatePeerConnection creates a new WebRTC peer connection
func (p *Proxy) CreatePeerConnection(sessionID string) (*webrtc.PeerConnection, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Configure WebRTC
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{p.config.StunServer},
			},
		},
	}

	// Create peer connection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create peer connection: %w", err)
	}

	// Store connection
	p.connections[sessionID] = peerConnection

	// Setup connection handlers
	p.setupConnectionHandlers(sessionID, peerConnection)

	return peerConnection, nil
}

// setupConnectionHandlers configures event handlers for a peer connection
func (p *Proxy) setupConnectionHandlers(sessionID string, pc *webrtc.PeerConnection) {
	// Handle ICE connection state changes
	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log.Printf("Session %s: ICE Connection State: %s", sessionID, state.String())

		if state == webrtc.ICEConnectionStateFailed ||
		   state == webrtc.ICEConnectionStateClosed ||
		   state == webrtc.ICEConnectionStateDisconnected {
			p.CloseConnection(sessionID)
		}
	})

	// Handle ICE candidates
	pc.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			log.Printf("Session %s: New ICE candidate: %s", sessionID, candidate.String())
		}
	})

	// Handle peer connection state changes
	pc.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		log.Printf("Session %s: Connection State: %s", sessionID, state.String())
	})
}

// CloseConnection closes and removes a peer connection
func (p *Proxy) CloseConnection(sessionID string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	pc, exists := p.connections[sessionID]
	if !exists {
		return fmt.Errorf("connection not found for session: %s", sessionID)
	}

	if err := pc.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	delete(p.connections, sessionID)
	log.Printf("Session %s: Connection closed", sessionID)

	return nil
}

// GetConnection retrieves a peer connection by session ID
func (p *Proxy) GetConnection(sessionID string) (*webrtc.PeerConnection, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	pc, exists := p.connections[sessionID]
	return pc, exists
}

// Stop gracefully shuts down the proxy
func (p *Proxy) Stop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	log.Println("Stopping WebRTC Proxy...")

	// Close all connections
	for sessionID, pc := range p.connections {
		if err := pc.Close(); err != nil {
			log.Printf("Error closing connection %s: %v", sessionID, err)
		}
	}

	p.connections = make(map[string]*webrtc.PeerConnection)
	log.Println("WebRTC Proxy stopped")

	return nil
}
