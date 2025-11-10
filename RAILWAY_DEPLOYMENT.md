# Railway Deployment Guide for Sweven Games

## Prerequisites
- GitHub account
- Railway account (sign up at https://railway.app)
- Repository: https://github.com/rekoonads/sweven-games-platform

## Deployment Steps

### Option 1: Deploy via Railway Dashboard (Recommended)

1. **Go to Railway**
   - Visit https://railway.app
   - Click "Start a New Project"

2. **Deploy from GitHub**
   - Click "Deploy from GitHub repo"
   - Authorize Railway to access your GitHub account
   - Select: `rekoonads/sweven-games-platform`

3. **Configure Services**

   Railway will detect the Dockerfile and automatically configure the build.

4. **Set Environment Variables**

   Add these environment variables in Railway dashboard:
   ```
   STUN_SERVER=stun:stun.l.google.com:19302
   PORT=8080
   VALIDATION_URL=your-auth-service-url
   ```

5. **Deploy**
   - Click "Deploy"
   - Railway will build and deploy your service
   - You'll get a public URL like: `your-app.up.railway.app`

### Option 2: Deploy via Railway CLI

1. **Install Railway CLI**
   ```bash
   npm i -g @railway/cli
   ```

2. **Login to Railway**
   ```bash
   railway login
   ```

3. **Initialize Project**
   ```bash
   cd "d:\cloud gaming\backend-services\sweven games"
   railway init
   ```

4. **Link to GitHub Repo**
   ```bash
   railway link
   ```

5. **Deploy**
   ```bash
   railway up
   ```

6. **Set Environment Variables**
   ```bash
   railway variables set STUN_SERVER=stun:stun.l.google.com:19302
   railway variables set PORT=8080
   ```

## Services to Deploy

### 1. WebRTC Proxy (Main Service)
- **Port**: 8080
- **Build**: `go build -o webrtc-proxy`
- **Start**: `./webrtc-proxy`
- **Environment**:
  - `STUN_SERVER=stun:stun.l.google.com:19302`
  - `PORT=8080`

### 2. Signaling Server
- **Port**: 8088 (WebSocket), 8000 (gRPC)
- **Path**: `signaling-server/`
- **Build**: `go build -o signaling`
- **Start**: `./signaling --websocket 8088 --grpc 8000`
- **Environment**:
  - `VALIDATION_URL=your-auth-url`

### 3. Daemon Service
- **Port**: Custom
- **Path**: `daemon/`
- **Build**: `go build -o daemon`
- **Start**: `./daemon`

## Multi-Service Deployment

To deploy multiple services, create separate Railway projects for each:

1. **WebRTC Proxy** - Main streaming service
2. **Signaling Server** - WebRTC signaling
3. **Auth Service** - From `auth-temp/`
4. **Daemon** - Backend daemon

## Important Notes

⚠️ **Large Files Warning**: The repository contains large executable files (>50MB). Railway handles this, but consider using Git LFS for better performance.

⚠️ **Submodules**: Railway automatically initializes git submodules during build.

⚠️ **Windows Binaries**: The `.exe` files are Windows-specific. Railway will rebuild from source using the Dockerfile.

## Monitoring & Logs

- View logs: `railway logs`
- Check service status: `railway status`
- Open deployed app: `railway open`

## Custom Domain

1. Go to Railway dashboard
2. Select your service
3. Go to "Settings" > "Domains"
4. Click "Generate Domain" or "Custom Domain"

## Scaling

Railway automatically scales based on usage. For manual scaling:
1. Go to service settings
2. Adjust "Resources"
3. Configure memory and CPU limits

## Troubleshooting

**Build Fails:**
- Check Railway build logs
- Verify Go version (1.18+)
- Ensure all dependencies are in go.mod

**Service Won't Start:**
- Check environment variables
- Verify PORT is set correctly
- Review application logs

**Connection Issues:**
- Verify firewall rules
- Check STUN/TURN server configuration
- Ensure WebSocket connections are allowed

## Repository Structure

```
sweven-games-platform/
├── Dockerfile           # Railway build configuration
├── railway.json         # Railway deployment config
├── daemon/              # Backend daemon service
├── webrtc-proxy/        # WebRTC streaming proxy
├── signaling-server/    # WebRTC signaling
├── auth-temp/           # Authentication service
├── browser-client/      # Web client (deploy separately)
├── sunshine-util/       # Utilities
├── DevSim/              # Development tools
├── script/              # Build scripts
└── package/             # Packaging files
```

## Support

- Railway Docs: https://docs.railway.app
- GitHub Issues: https://github.com/rekoonads/sweven-games-platform/issues
- Railway Discord: https://discord.gg/railway

---

**Deployment URL**: After deployment, Railway will provide a URL like:
`https://sweven-games-production.up.railway.app`
