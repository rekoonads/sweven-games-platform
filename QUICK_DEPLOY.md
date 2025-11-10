# 🚀 Quick Deploy Guide - Sweven Games

## Repository
🔗 **https://github.com/rekoonads/sweven-games-platform**

---

## ⚡ Fastest Way: Render (All-in-One)

### 1-Click Deploy to Render

1. **Go to:** https://render.com/deploy
2. **Click:** "New Blueprint Instance"
3. **Connect Repository:** `rekoonads/sweven-games-platform`
4. **Deploy:** Render will automatically detect `render.yaml` and deploy all 3 services!

**Done!** You'll get 3 URLs:
- `https://sweven-webrtc-proxy.onrender.com` (WebRTC Proxy)
- `https://sweven-signaling.onrender.com` (Signaling)
- `https://sweven-daemon.onrender.com` (Daemon)

**Cost:** Free tier available, then $7/month per service

---

## 🔥 Railway (Multiple Services)

### Service 1: WebRTC Proxy

```bash
# Go to: https://railway.app/new
# Select: rekoonads/sweven-games-platform
# Settings > Build > Dockerfile Path: railway-services/webrtc-proxy.Dockerfile
# Environment Variables:
STUN_SERVER=stun:stun.l.google.com:19302
PORT=8080
```

### Service 2: Signaling Server

```bash
# Create new project: https://railway.app/new
# Select: rekoonads/sweven-games-platform
# Settings > Build > Dockerfile Path: railway-services/signaling-server.Dockerfile
# Environment Variables:
WEBSOCKET_PORT=8088
GRPC_PORT=8000
```

### Service 3: Daemon

```bash
# Create new project: https://railway.app/new
# Select: rekoonads/sweven-games-platform
# Settings > Build > Dockerfile Path: railway-services/daemon.Dockerfile
```

**Cost:** ~$5/service/month (~$15 total)

---

## 🐳 Docker Compose (Local/VPS)

```bash
cd "d:\cloud gaming\backend-services\sweven games"

# Build and start all services
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

**Services will run on:**
- WebRTC Proxy: http://localhost:8080
- Signaling: http://localhost:8088
- Daemon: Internal only

---

## 📋 Comparison

| Option | Deploy Time | Cost | Services | Difficulty |
|--------|-------------|------|----------|------------|
| **Render** | 5 min | Free-$21/mo | 3 services | ⭐ Easiest |
| **Railway** | 15 min | $15-40/mo | 3 services | ⭐⭐ Easy |
| **Docker Compose** | 5 min | VPS cost | All services | ⭐⭐⭐ Medium |

---

## 🎯 Recommended Choice

### For Testing/Demo:
✅ **Render** - Free tier, 1-click deploy, all services together

### For Production:
✅ **Railway** - Better scaling, independent services, reliable

### For Self-Hosting:
✅ **Docker Compose** - Full control, no monthly fees

---

## 🔗 After Deployment

Once deployed, update your frontend environment variables:

```env
NEXT_PUBLIC_SIGNALING_URL=https://your-signaling-url.com
NEXT_PUBLIC_WEBRTC_URL=https://your-webrtc-url.com
```

---

## 📚 Full Documentation

- **Complete Strategy:** [DEPLOYMENT_STRATEGY.md](./DEPLOYMENT_STRATEGY.md)
- **Railway Guide:** [RAILWAY_DEPLOYMENT.md](./RAILWAY_DEPLOYMENT.md)
- **Repository:** https://github.com/rekoonads/sweven-games-platform

---

## ❓ Need Help?

**Pick Render if:** You want the easiest, fastest deployment
**Pick Railway if:** You want production-ready, scalable services
**Pick Docker if:** You have your own server or want full control
