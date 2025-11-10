# 🚀 Final Deployment Guide - Sweven Games Platform

## Current Status

### ✅ Already Deployed:
- **Auth Service** - Authentication (already live)

### 📦 In Repository (Ready to Deploy):
- **Browser Client** - Frontend web application
- **DevSim** - Development simulator
- **Signaling Server** - WebRTC signaling infrastructure
- **Sunshine Util** - Screen capture/encoding utilities
- **WebRTC Proxy** - Custom streaming proxy (newly created)
- **Daemon** - Backend service

### 🎯 Need to Deploy:
1. **Signaling Server** (ports 8088, 8000)
2. **WebRTC Proxy** (port 8080)
3. **Daemon** (background worker)

---

## 🚀 One-Click Deploy All 3 Services

### **Option 1: Render.com (RECOMMENDED - All at Once)** ⭐

This will deploy all 3 backend services in one go:

**Steps:**
1. Go to: **https://render.com/deploy**
2. Click: **"New Blueprint Instance"**
3. Connect repo: **`rekoonads/sweven-games-platform`**
4. Render detects `render.yaml` automatically
5. **Add your auth service URL** when prompted:
   ```
   VALIDATION_URL=<your-already-deployed-auth-url>
   ```
6. Click **"Apply"**

**Result:** All 3 services deployed together!
- ✅ Signaling Server: `https://sweven-signaling-xxxx.onrender.com`
- ✅ WebRTC Proxy: `https://sweven-webrtc-proxy-xxxx.onrender.com`
- ✅ Daemon: (background worker)

**Time:** 5-10 minutes
**Cost:** Free tier available, then $7/service/month

---

### **Option 2: Railway (Separate Services)**

Deploy each service individually for better control:

#### **Service 1: Signaling Server**
```
1. Go to: https://railway.app/new
2. Select: rekoonads/sweven-games-platform
3. Settings > Build:
   - Dockerfile Path: railway-services/signaling-server.Dockerfile
4. Environment Variables:
   WEBSOCKET_PORT=8088
   GRPC_PORT=8000
   VALIDATION_URL=<your-auth-url>
5. Deploy
```

#### **Service 2: WebRTC Proxy**
```
1. New Project: https://railway.app/new
2. Select: rekoonads/sweven-games-platform
3. Settings > Build:
   - Dockerfile Path: railway-services/webrtc-proxy.Dockerfile
4. Environment Variables:
   STUN_SERVER=stun:stun.l.google.com:19302
   PORT=8080
   SIGNALING_SERVER_URL=<url-from-service-1>
5. Deploy
```

#### **Service 3: Daemon**
```
1. New Project: https://railway.app/new
2. Select: rekoonads/sweven-games-platform
3. Settings > Build:
   - Dockerfile Path: railway-services/daemon.Dockerfile
4. Environment Variables:
   SIGNALING_SERVER_URL=<url-from-service-1>
5. Deploy
```

**Time:** 15-20 minutes
**Cost:** ~$5/service/month (~$15 total)

---

### **Option 3: Docker Compose (Local/VPS)**

Run all services on your own server:

```bash
cd "d:\cloud gaming\backend-services\sweven games"

# Edit docker-compose.yml and add your auth URL
# Then start all services:
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f
```

**Services available at:**
- Signaling: http://localhost:8088
- WebRTC Proxy: http://localhost:8080
- Daemon: (internal)

---

## 🔗 After Deployment - Connect Services

### Update Frontend Environment Variables

Once backend services are deployed, update your frontend:

```env
# .env.local or Vercel Environment Variables

# Your already-deployed auth service
NEXT_PUBLIC_AUTH_URL=<your-auth-service-url>

# Newly deployed backend services
NEXT_PUBLIC_SIGNALING_URL=<signaling-server-url>
NEXT_PUBLIC_WEBRTC_URL=<webrtc-proxy-url>

# Example:
# NEXT_PUBLIC_AUTH_URL=https://sweven-auth.railway.app
# NEXT_PUBLIC_SIGNALING_URL=https://sweven-signaling.onrender.com
# NEXT_PUBLIC_WEBRTC_URL=https://sweven-webrtc-proxy.onrender.com
```

---

## 📊 Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│                    SWEVEN GAMES PLATFORM                │
└─────────────────────────────────────────────────────────┘

Frontend (Browser Client)
    ↓
    ├──→ Auth Service (✅ Already Deployed)
    │    └── Authentication & User Management
    │
    ├──→ Signaling Server (🔴 Need to Deploy)
    │    └── WebRTC Connection Setup
    │    └── Ports: 8088 (WebSocket), 8000 (gRPC)
    │
    └──→ WebRTC Proxy (🔴 Need to Deploy)
         └── Video/Audio Streaming
         └── Port: 8080

Background:
    └──→ Daemon (🔴 Need to Deploy)
         └── Backend Processing

Utilities:
    ├── Sunshine Util (Screen capture/encoding)
    ├── DevSim (Development tools)
    └── Browser Client (Web interface)
```

---

## ✅ Quick Checklist

Before deploying:
- [ ] Have your **Auth Service URL** ready
- [ ] Choose deployment platform (Render recommended)
- [ ] Prepare to set environment variables

After deploying:
- [ ] Note down all 3 service URLs
- [ ] Update frontend environment variables
- [ ] Test the connection between services
- [ ] Deploy frontend to Vercel

---

## 🎯 Recommended Deployment Flow

### For Quick Testing:
1. **Render** - Deploy all 3 backend services (one click)
2. **Vercel** - Deploy frontend
3. Connect them via environment variables

### For Production:
1. **Railway** - Deploy each service separately for better control
2. **Vercel** - Deploy frontend with custom domain
3. Set up monitoring and logging

---

## 📞 Need Help?

**Platform Issues:**
- Render Support: https://render.com/docs
- Railway Support: https://railway.app/help

**Repository:**
- GitHub: https://github.com/rekoonads/sweven-games-platform
- Issues: https://github.com/rekoonads/sweven-games-platform/issues

---

## 🔥 Start Deploying NOW

### **Fastest Way (5 minutes):**
👉 https://render.com/deploy
- Connect: `rekoonads/sweven-games-platform`
- Add your auth URL
- Deploy!

### **Full Control (15 minutes):**
👉 https://railway.app/new
- Deploy each service separately
- Better scaling options

---

**Repository:** https://github.com/rekoonads/sweven-games-platform

**Files You Need:**
- ✅ `render.yaml` - Render deployment config
- ✅ `docker-compose.yml` - Docker compose config
- ✅ `railway-services/` - Individual Dockerfiles
- ✅ All source code for services
