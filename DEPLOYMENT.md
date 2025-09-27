# AI API Server Deployment Examples

## Deploy to Railway

1. Connect your GitHub repository to Railway
2. Railway will automatically detect the Dockerfile and deploy
3. Set environment variables if needed:
   - `PORT` (optional, Railway sets this automatically)
   - `AI_MOCK_MODE=false` (to use real AI service)

## Deploy to Fly.io

1. Install flyctl: `curl -L https://fly.io/install.sh | sh`
2. Login: `flyctl auth login`
3. Create app: `flyctl launch`
4. Deploy: `flyctl deploy`

## Deploy to Heroku

### Option 1: Using Docker
```bash
# Install Heroku CLI and login
heroku login

# Create app
heroku create your-ai-api

# Set container stack
heroku stack:set container

# Deploy
git push heroku main
```

### Option 2: Using Heroku Go Buildpack
```bash
# Create Procfile
echo "web: ./ai-server" > Procfile

# Create app with Go buildpack
heroku create your-ai-api --buildpack heroku/go

# Deploy
git push heroku main
```

## Deploy to Google Cloud Run

```bash
# Build and push to Google Container Registry
gcloud builds submit --tag gcr.io/PROJECT_ID/ai-server

# Deploy to Cloud Run
gcloud run deploy ai-server \
  --image gcr.io/PROJECT_ID/ai-server \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated
```

## Deploy to AWS ECS/Fargate

1. Build and push Docker image to ECR
2. Create ECS task definition
3. Create ECS service with Fargate launch type
4. Configure load balancer and security groups

## Local Docker Deployment

```bash
# Build image
docker build -t ai-server .

# Run container
docker run -d \
  --name ai-api \
  -p 8081:8081 \
  -e AI_MOCK_MODE=false \
  ai-server

# Check logs
docker logs ai-api

# Stop container
docker stop ai-api
```

## Environment Variables

- `PORT`: Server port (default: 8081)
- `AI_MOCK_MODE`: Set to "true" for mock responses, "false" for real AI (default: false)

## Health Check

All deployment platforms can use the `/health` endpoint for health checks:
- URL: `http://your-domain/health`
- Expected response: `200 OK` with JSON status