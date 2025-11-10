# WebRTC Proxy

WebRTC protocol data and multimedia stream transport agent for Sweven Games.

## Overview

The webrtc-proxy handles WebRTC protocol transport for streaming game content from the server to browser clients. It manages:

- WebRTC peer connections
- Media stream transport (video/audio)
- Data channel communication
- Connection state management

## Architecture

This component works with:
- **screencoder** (sunshine-util): Receives encoded video/audio streams
- **signaling-server**: Handles WebRTC signaling and connection setup
- **browser-client**: Delivers streams to web clients

## Features

- WebRTC peer connection management
- RTP/RTCP packet handling
- Media track management
- ICE candidate handling
- Connection state monitoring

## Configuration

The proxy can be configured through environment variables or command-line arguments:

- `STUN_SERVER`: STUN server URL (default: stun:stun.l.google.com:19302)
- `TURN_SERVER`: TURN server URL if needed
- `PORT`: Proxy service port

## Usage

```bash
go run main.go
```

## Development Status

Custom implementation in progress for Sweven Games cloud gaming platform.
