package proxy

import (
	"fmt"
	"log"

	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
)

// StreamManager handles media streaming for WebRTC connections
type StreamManager struct {
	peerConnection *webrtc.PeerConnection
	videoTrack     *webrtc.TrackLocalStaticRTP
	audioTrack     *webrtc.TrackLocalStaticRTP
}

// NewStreamManager creates a new stream manager
func NewStreamManager(pc *webrtc.PeerConnection) *StreamManager {
	return &StreamManager{
		peerConnection: pc,
	}
}

// CreateVideoTrack creates a video track for streaming
func (sm *StreamManager) CreateVideoTrack() (*webrtc.TrackLocalStaticRTP, error) {
	// Create a video track
	videoTrack, err := webrtc.NewTrackLocalStaticRTP(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeH264},
		"video",
		"sweven-video",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create video track: %w", err)
	}

	// Add track to peer connection
	_, err = sm.peerConnection.AddTrack(videoTrack)
	if err != nil {
		return nil, fmt.Errorf("failed to add video track: %w", err)
	}

	sm.videoTrack = videoTrack
	log.Println("Video track created and added")

	return videoTrack, nil
}

// CreateAudioTrack creates an audio track for streaming
func (sm *StreamManager) CreateAudioTrack() (*webrtc.TrackLocalStaticRTP, error) {
	// Create an audio track
	audioTrack, err := webrtc.NewTrackLocalStaticRTP(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus},
		"audio",
		"sweven-audio",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create audio track: %w", err)
	}

	// Add track to peer connection
	_, err = sm.peerConnection.AddTrack(audioTrack)
	if err != nil {
		return nil, fmt.Errorf("failed to add audio track: %w", err)
	}

	sm.audioTrack = audioTrack
	log.Println("Audio track created and added")

	return audioTrack, nil
}

// WriteVideoRTP writes an RTP packet to the video track
func (sm *StreamManager) WriteVideoRTP(packet *rtp.Packet) error {
	if sm.videoTrack == nil {
		return fmt.Errorf("video track not initialized")
	}

	if err := sm.videoTrack.WriteRTP(packet); err != nil {
		return fmt.Errorf("failed to write video RTP packet: %w", err)
	}

	return nil
}

// WriteAudioRTP writes an RTP packet to the audio track
func (sm *StreamManager) WriteAudioRTP(packet *rtp.Packet) error {
	if sm.audioTrack == nil {
		return fmt.Errorf("audio track not initialized")
	}

	if err := sm.audioTrack.WriteRTP(packet); err != nil {
		return fmt.Errorf("failed to write audio RTP packet: %w", err)
	}

	return nil
}

// GetVideoTrack returns the video track
func (sm *StreamManager) GetVideoTrack() *webrtc.TrackLocalStaticRTP {
	return sm.videoTrack
}

// GetAudioTrack returns the audio track
func (sm *StreamManager) GetAudioTrack() *webrtc.TrackLocalStaticRTP {
	return sm.audioTrack
}
