# AI Service API

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![API Version](https://img.shields.io/badge/API-v1.0.0-blue.svg)](https://github.com/Ammar0144/ai)

> **🎓 Learning Project**: This is an educational project designed for learning AI integration and breaking the ice with AI services. Perfect for developers who want to understand how to build and integrate AI APIs in real-world applications!

A comprehensive Go-based AI service providing various artificial intelligence capabilities with advanced rate limiting, CORS support, and robust error handling. Built for production use with integrated LLM backend and comprehensive security features.

## 📚 About This Project

This project is created as a **learning resource** for developers interested in:
- 🧠 **AI Integration**: Learn how to integrate LLM services into your applications
- 🏗️ **Microservices Architecture**: Understand how to build scalable AI gateway services
- 🔧 **Go Development**: Practice building production-ready APIs with Go
- 🐳 **DevOps**: Learn Docker, CI/CD, and deployment strategies

### 🎯 Project Goals
- ✅ Provide a practical example of AI service integration
- ✅ Demonstrate best practices for API design and security
- ✅ Offer a starting point for your own AI projects
- ✅ Create a learning community around AI development

### 💡 Why This Project?
Breaking into AI development can be intimidating. This project aims to make it accessible by providing:
- Clear, well-documented code
- Practical examples you can run immediately
- A foundation you can build upon
- Real-world patterns and practices

### 🔮 Future Roadmap
This project is **continuously evolving** as we learn more about AI integration:
- 🚀 More AI capabilities and endpoints
- 📈 Performance optimizations
- 🔐 Enhanced security features
- 📖 More comprehensive documentation
- 🎓 Tutorial content and guides
- 🤝 Community contributions and best practices

### 🤝 Join the Learning Journey!
We **encourage you** to:
- ⭐ Star this repo if you find it helpful
- 🐛 Report issues and suggest improvements
- 💬 Share your thoughts and experiences
- 🤲 Contribute your own enhancements
- 📣 Share this with others learning AI integration

**Your feedback and contributions help everyone learn!** Whether you're a beginner or experienced developer, your perspective is valuable.

## � Companion Project: LLM Backend Server

**This AI Gateway works best with our LLM Backend Service!**

### 🤝 How They Work Together

This **AI Service** acts as a **public-facing gateway** that:
- ✅ Handles client requests with rate limiting and security
- ✅ Provides a clean, documented API interface
- ✅ Routes requests to the internal LLM backend
- ✅ Manages authentication and access control

The **[LLM Backend Service](https://github.com/Ammar0144/llm)** is the **internal AI engine** that:
- 🧠 Runs the actual DistilGPT-2 language model
- 🔒 Stays isolated from public internet (security)
- ⚡ Processes text generation requests
- 🎯 Optimized for low-resource environments

### 🏗️ Architecture Overview

```
┌─────────────┐         ┌──────────────┐         ┌──────────────┐
│   Clients   │────────▶│ AI Gateway   │────────▶│ LLM Backend  │
│  (Public)   │  HTTP   │ (This Repo)  │  HTTP   │  (Python)    │
│             │         │ Port 8081    │         │  Port 8082   │
└─────────────┘         │              │         │              │
                        │ - Go-based   │         │ - FastAPI    │
                        │ - Rate limit │         │ - DistilGPT-2│
                        │ - CORS       │         │ - IP access  │
                        │ - Public API │         │ - Internal   │
                        └──────────────┘         └──────────────┘
```

### 💡 Why This Architecture?

**Learn Real-World Patterns:**
- 🏢 **Microservices**: Separate concerns (gateway vs processing)
- 🔒 **Security Layers**: Public gateway + internal service
- ⚖️ **Load Distribution**: Gateway handles traffic, backend handles AI
- 🔧 **Technology Choice**: Right tool for the job (Go for API, Python for ML)

### 🚀 Try Both Projects!

**Explore the LLM Backend** to learn:
- 🐍 Python and FastAPI development
- 🤖 ML model integration with Transformers
- 🔐 IP-based access control
- 📦 Containerizing ML services

**👉 Check it out**: [github.com/Ammar0144/llm](https://github.com/Ammar0144/llm)

### 📦 Quick Start with Both Services

```bash
# Clone both repositories
git clone https://github.com/Ammar0144/ai.git
git clone https://github.com/Ammar0144/llm.git

# Run with Docker Compose (easiest way)
cd ai/
docker-compose up -d

# Or run separately
# Terminal 1 - LLM Backend
cd llm/
python server.py

# Terminal 2 - AI Gateway
cd ai/
go run main.go
```

**Together, they demonstrate a complete AI service architecture!** 🎯

## �💻 What You'll Learn

This project is packed with **real-world software engineering concepts** and best practices:

### 🎯 Core Concepts You'll Master
- **🏗️ API Design**: RESTful patterns, resource modeling, HTTP methods
- **⚡ Rate Limiting**: IP-based limiting, token bucket algorithm, concurrent handling
- **🔒 Security**: CORS, proxy headers, access control, input validation
- **🚀 CI/CD**: GitHub Actions, multi-platform builds, automated deployment
- **📖 API Documentation**: Swagger/OpenAPI, auto-generated docs, interactive testing
- **🐳 Containerization**: Docker, multi-stage builds, Docker Compose orchestration
- **🔗 Microservices**: Gateway pattern, service communication, health checks
- **🧪 Testing**: Unit tests, integration tests, test-driven development
- **📊 Monitoring**: Health endpoints, logging, error tracking
- **⚙️ Middleware**: Request interception, response modification, chain of responsibility

### 📚 Detailed Learning Path

Want to dive deeper? Check out our comprehensive guide:
**[📖 Software Engineering Concepts Guide](SOFTWARE_ENGINEERING_CONCEPTS.md)**

This guide includes:
- ✅ Detailed explanations of each concept
- ✅ Code examples with line numbers
- ✅ Learning paths for beginner/intermediate/advanced
- ✅ Hands-on exercises to practice
- ✅ Real-world applications
- ✅ Additional resources and tutorials

### 🎓 Learning by Experience Level

**Beginners**: Focus on API structure, Docker basics, error handling  
**Intermediate**: Study rate limiting, middleware patterns, CI/CD workflows  
**Advanced**: Explore microservices architecture, performance optimization, production deployment

**Every file teaches something valuable!** Explore the codebase with curiosity.

## 🚀 Features

### Core AI Capabilities
- 🤖 **Chat Completions**: Multi-turn conversation-based AI interactions
- ✍️ **Text Completion**: Complete text based on prompts
- 🎨 **Text Generation**: Generate creative text content
- ℹ️ **Model Information**: Query current AI model capabilities

### Security & Performance
- 🔒 **Advanced Rate Limiting**: IP-based rate limiting with configurable limits
- 🌐 **CORS Support**: Cross-origin resource sharing configuration
- 🛡️ **Proxy Header Support**: X-Real-IP and X-Forwarded-For handling
- 🧹 **Memory Management**: Automatic IP cache cleanup to prevent leaks
- 📊 **Comprehensive Monitoring**: Health checks and service metrics

### Development & Documentation
- � **Interactive Swagger UI**: Complete API documentation and testing
- 🧪 **Comprehensive Test Coverage**: Unit and integration tests
- 🐳 **Docker Support**: Containerized deployment
- ⚡ **GitHub Actions CI/CD**: Automated testing and deployment
- 📝 **Structured Logging**: Detailed operation logging

## 🔒 Security & Rate Limiting

### Rate Limits (per IP address per minute)

| Endpoint Category | Rate Limit | Endpoints |
|-------------------|------------|-----------|
| **AI Processing** | 30 req/min | `/ai/chat/completions`, `/ai/complete`, `/ai/generate` |
| **Health Check** | 200 req/min | `/health` |
| **Model Info** | 100 req/min | `/ai/model-info` |
| **Root Info** | 100 req/min | `/` |

### Security Features
- **IP-based Rate Limiting**: Prevents abuse and ensures fair usage
- **CORS Configuration**: Configurable cross-origin resource sharing
- **Proxy Support**: Handles load balancer and proxy headers
- **Error Handling**: Comprehensive error responses with proper HTTP status codes
- **Memory Safety**: Automatic cleanup of rate limiting data structures

## 📖 API Documentation

### Interactive Documentation
- **Swagger UI**: [http://localhost:8081/swagger/](http://localhost:8081/swagger/)
- **Documentation Page**: [http://localhost:8081/docs](http://localhost:8081/docs)
- **OpenAPI Spec**: [http://localhost:8081/swagger/doc.json](http://localhost:8081/swagger/doc.json)

### Core API Endpoints

#### 🤖 AI Processing Endpoints (Rate: 30/min)

##### POST /ai/chat/completions
Generate chat completions based on conversation history.

**Request:**
```json
{
  "messages": [
    {"role": "user", "content": "Hello, how are you?"}
  ],
  "max_tokens": 150,
  "temperature": 0.7,
  "user_id": "optional-user-id"
}
```

**Response:**
```json
{
  "response": "Hello! I'm doing well, thank you for asking. How can I help you today?",
  "user_id": "optional-user-id",
  "timestamp": "2025-10-05T04:00:00Z",
  "model": "distilgpt2"
}
```

##### POST /ai/complete
Complete text based on a given prompt.

**Request:**
```json
{
  "prompt": "Docker containers are",
  "max_tokens": 100,
  "temperature": 0.7,
  "user_id": "optional-user-id"
}
```

**Response:**
```json
{
  "response": "portable, lightweight, and efficient for deploying applications...",
  "user_id": "optional-user-id",
  "timestamp": "2025-10-05T04:00:00Z",
  "model": "distilgpt2"
}
```

##### POST /ai/generate
Generate creative text content based on a prompt.

**Request:**
```json
{
  "prompt": "The future of cloud computing will",
  "max_tokens": 100,
  "temperature": 0.8,
  "user_id": "optional-user-id"
}
```

**Response:**
```json
{
  "response": "be shaped by advances in artificial intelligence and automation...",
  "user_id": "optional-user-id",
  "timestamp": "2025-10-05T04:00:00Z",
  "model": "distilgpt2"
}
```

#### 🏥 Utility Endpoints

##### GET /health (Rate: 200/min)
Service health and status check.

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2025-10-05T04:00:00Z",
  "version": "1.0.0"
}
```

##### GET /ai/model-info (Rate: 100/min)
Current AI model information and capabilities.

**Response:**
```json
{
  "model_name": "distilgpt2",
  "model_type": "GPT-2",
  "model_size": "82M parameters",
  "description": "DistilGPT-2 optimized for text generation and completion tasks",
  "optimized_for": [
    "text_generation",
    "text_completion",
    "chat_conversations"
  ],
  "supported_endpoints": [
    "/generate - Text generation (primary strength)",
    "/complete - Text completion (primary strength)",
    "/chat/completions - Chat-style conversations",
    "/health - Service health check",
    "/ - Basic status"
  ]
}
```

##### GET / (Rate: 100/min)
Service information and available endpoints.

**Response:**
```json
{
  "service": "AI Service API",
  "version": "1.0.0",
  "documentation": {
    "swagger_ui": "/swagger/",
    "docs": "/docs",
    "openapi_spec": "/swagger/doc.json"
  },
  "endpoints": {
    "health": "/health",
    "chat_completions": "/ai/chat/completions",
    "complete": "/ai/complete",
    "generate": "/ai/generate",
    "model_info": "/ai/model-info"
  }
}
```

### Error Responses
All endpoints return standardized error responses:

```json
{
  "error": "Rate limit exceeded",
  "code": 429,
  "message": "Too many requests. Please try again later."
}
```

**HTTP Status Codes:**
- `200` - Success
- `400` - Bad Request (invalid input)
- `429` - Rate Limit Exceeded
- `500` - Internal Server Error

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or later
- LLM backend service (DistilGPT-2 server)
- Internet connection (for external fallbacks)

### Local Development

1. **Clone the repository:**
```bash
git clone https://github.com/Ammar0144/ai.git
cd ai
```

2. **Install dependencies:**
```bash
go mod download
```

3. **Generate Swagger documentation:**
```bash
# Install swag if not already installed
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g main.go --output ./docs
```

4. **Run the server:**
```bash
go run main.go
```

5. **Test the API:**
```bash
# Health check
curl http://localhost:8081/health

# Chat completion
curl -X POST http://localhost:8081/ai/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "messages": [{"role": "user", "content": "Hello!"}],
    "temperature": 0.7
  }'

# Text completion
curl -X POST http://localhost:8081/ai/complete \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "The future of AI is",
    "max_tokens": 50,
    "temperature": 0.7
  }'

# Text generation
curl -X POST http://localhost:8081/ai/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "Once upon a time",
    "max_length": 100,
    "temperature": 0.8
  }'
```

6. **Access Documentation:**
- Swagger UI: http://localhost:8081/swagger/
- Documentation: http://localhost:8081/docs
- Service info: http://localhost:8081/

### 🐳 Docker Deployment

1. **Build the Docker image:**
```bash
docker build -t ai-service .
```

2. **Run with Docker Compose (Recommended):**
```bash
# Runs both AI service and LLM backend
docker-compose up -d
```

3. **Or run standalone:**
```bash
docker run -p 8081:8081 \
  -e LLM_SERVICE_URL=http://localhost:8082 \
  ai-service
```

### ⚙️ Configuration

#### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8081` | Server port |
| `LLM_SERVICE_URL` | `http://localhost:8082` | LLM backend URL |
| `AI_MOCK_MODE` | `false` | Use mock responses for testing |
| `CORS_ORIGINS` | `*` | Allowed CORS origins |
| `RATE_LIMIT_AI` | `30` | Rate limit for AI endpoints (per minute) |
| `RATE_LIMIT_HEALTH` | `200` | Rate limit for health endpoint (per minute) |
| `RATE_LIMIT_INFO` | `100` | Rate limit for info endpoint (per minute) |

#### Rate Limiting Configuration
```bash
# Custom rate limits
export RATE_LIMIT_AI=50
export RATE_LIMIT_HEALTH=300
export RATE_LIMIT_INFO=150

go run main.go
```

## Development

### Running Tests
```bash
go test ./...
```

### Running Tests with Coverage
```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Code Formatting
```bash
gofmt -s -w .
```

### Linting
```bash
go vet ./...
```

## 📁 Project Structure

```
ai/
├── main.go                 # Main server entry point with middleware
├── go.mod                 # Go module dependencies and versions
├── go.sum                 # Dependency checksums
│
├── handlers/              # HTTP request handlers
│   ├── ai.go             # AI endpoints with Swagger annotations
│   └── ai_test.go        # Comprehensive handler tests
│
├── services/             # Business logic services
│   ├── ai_service.go     # LLM backend integration service
│   └── ai_service_test.go # Service layer tests
│
├── models/               # Data structures and schemas
│   └── message.go        # Request/response models for all endpoints
│
├── docs/                 # API documentation
│   ├── docs.go           # Generated Swagger configuration
│   ├── swagger.json      # OpenAPI specification (JSON)
│   ├── swagger.yaml      # OpenAPI specification (YAML)
│   └── index.html        # Custom documentation page
│
├── .github/workflows/    # CI/CD automation
│   └── ci.yml           # GitHub Actions pipeline
│
├── Dockerfile           # Container configuration
├── docker-compose.yml   # Multi-service deployment
├── .gitignore          # Git ignore patterns
├── .dockerignore       # Docker ignore patterns
└── README.md           # This documentation
```

### Key Components

#### 🚀 Main Server (`main.go`)
- **Rate Limiting**: IP-based request throttling with configurable limits
- **CORS Middleware**: Cross-origin resource sharing configuration
- **Swagger Integration**: API documentation and testing interface
- **Health Monitoring**: Service health and metrics endpoints
- **Error Handling**: Comprehensive error handling and logging

#### 🎯 Handlers (`handlers/ai.go`)
- **AI Processing**: All AI-related endpoint implementations
- **Swagger Annotations**: Complete API documentation annotations
- **Input Validation**: Request validation and sanitization
- **Response Formatting**: Standardized JSON responses
- **Error Management**: Proper HTTP status codes and error messages

#### 🧠 Services (`services/ai_service.go`)
- **LLM Integration**: Backend service communication
- **Fallback Logic**: Intelligent mock response system
- **Request Processing**: LLM request formatting and handling
- **Response Processing**: LLM response parsing and validation

#### 📋 Models (`models/message.go`)
- **Request Models**: Structured input validation
- **Response Models**: Consistent output formatting
- **Error Models**: Standardized error responses
- **Schema Definitions**: OpenAPI-compatible structures

## 🏗️ Architecture

### System Overview
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Client Apps   │───▶│   AI Service     │───▶│   LLM Backend   │
│   (Web, Mobile) │    │   (Port 8081)    │    │   (Port 8082)   │
└─────────────────┘    └──────────────────┘    └─────────────────┘
                              │
                              ▼
                       ┌──────────────────┐
                       │  Rate Limiting   │
                       │  CORS & Security │
                       │  Request Routing │
                       └──────────────────┘
```

### 🤖 LLM Backend Integration

The AI service integrates with a **dedicated LLM server** running DistilGPT-2:

#### Integration Features
- **Internal Communication**: Connects to LLM server via `http://llm-server:8082`
- **Secure Access**: LLM server restricts access to internal networks only
- **Intelligent Fallbacks**: Mock responses when LLM service unavailable
- **Load Balancing**: Distributes requests across multiple LLM instances
- **Health Monitoring**: Continuous health checks of LLM backend

#### 🔒 Security Architecture

```
Internet → AI Service (Public) → LLM Service (Internal Only)
          Port 8081              Port 8082 (IP restricted)
          │                      │
          ├─ Rate Limiting       ├─ Access Control
          ├─ CORS Policy         ├─ Internal Networks Only
          ├─ Request Validation  ├─ Health Monitoring
          └─ Response Filtering  └─ Resource Management
```

#### Multi-Layer Security
- **AI Service Layer**: Public-facing with comprehensive rate limiting and CORS
- **LLM Service Layer**: Internal-only with IP access control and resource limits
- **Network Layer**: Secure communication between services
- **Application Layer**: Input validation and response sanitization

### Fallback System

#### Mock Response Categories
- **Conversational**: Greetings, farewells, casual conversation
- **Informational**: Time, date, help requests, system status
- **Emotional**: Expressions of gratitude, well-being inquiries
- **Technical**: API information, capability descriptions
- **Generic**: Intelligent fallback responses for unmatched inputs

#### Fallback Logic
1. **Primary**: LLM backend processing
2. **Secondary**: Category-based mock responses
3. **Tertiary**: Generic helpful responses
4. **Error**: Graceful error handling with user-friendly messages

## ⚡ CI/CD Pipeline

### Optimized Build Process

The GitHub Actions workflow automatically:

1. **Build Phase (Optimized):**
   - Creates multi-platform binaries (Linux, Windows, macOS)
   - Builds Docker image for containerized deployment
   - Performs static analysis with `go vet`
   - Validates Go code compilation

2. **Artifact Management:**
   - **Smart Uploads**: Only uploads artifacts on main branch pushes
   - **Short Retention**: 7-day retention period to manage storage quotas
   - **Unique Naming**: Artifacts named with run numbers to avoid conflicts
   - **Conditional Processing**: Skips artifact uploads for pull requests

3. **Automated Cleanup:**
   - **Daily Cleanup**: Automatic removal of artifacts older than 7 days
   - **Workflow Management**: Keeps only the last 50 workflow runs
   - **Storage Optimization**: Prevents GitHub storage quota issues
   - **Manual Cleanup**: PowerShell and Bash scripts for immediate cleanup

4. **Deployment Ready:**
   - Artifacts available for download from successful builds
   - Docker images ready for container platform deployment
   - Automated cleanup prevents storage quota exhaustion

## 🚀 Deployment Options

### Production Deployment

#### 🔧 Full Stack Deployment (Recommended)
```bash
# Deploy both AI service and LLM backend
git clone https://github.com/Ammar0144/ai.git
cd ai

# Using Docker Compose
docker-compose up -d

# Verify deployment
curl http://localhost:8081/health
curl http://localhost:8082/health
```

#### ☁️ Cloud Platforms

##### **Railway**
```bash
# Connect GitHub repository or upload Docker image
railway login
railway link
railway up
```

##### **Fly.io**
```bash
# Deploy with flyctl
flyctl launch
flyctl deploy
```

##### **Google Cloud Run**
```bash
# Build and deploy
gcloud builds submit --tag gcr.io/PROJECT-ID/ai-service
gcloud run deploy --image gcr.io/PROJECT-ID/ai-service --port 8081
```

##### **AWS ECS/Fargate**
```bash
# Build and push to ECR
aws ecr build-and-push --repository ai-service
# Deploy via ECS console or CLI
```

##### **Digital Ocean App Platform**
```bash
# Connect GitHub repository for automatic deployment
# Configure environment variables via DO console
```

#### 🏠 Self-Hosted

##### **Direct Binary Deployment**
```bash
# Build for production
CGO_ENABLED=0 GOOS=linux go build -o ai-service main.go

# Transfer and run on server
scp ai-service user@server:/opt/ai-service/
ssh user@server "cd /opt/ai-service && ./ai-service"
```

##### **Docker Deployment**
```bash
# Single service
docker run -d \
  --name ai-service \
  -p 8081:8081 \
  -e LLM_SERVICE_URL=http://llm-backend:8082 \
  ai-service:latest

# With reverse proxy (nginx/traefik)
docker-compose -f docker-compose.prod.yml up -d
```

##### **Systemd Service**
```bash
# Create systemd service file
sudo tee /etc/systemd/system/ai-service.service << EOF
[Unit]
Description=AI Service API
After=network.target

[Service]
Type=simple
User=ai-service
WorkingDirectory=/opt/ai-service
ExecStart=/opt/ai-service/ai-service
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

# Enable and start
sudo systemctl enable ai-service
sudo systemctl start ai-service
```

### 📊 Production Configuration

#### Environment Variables for Production
```bash
# Production settings
export PORT=8081
export LLM_SERVICE_URL=http://internal-llm:8082
export CORS_ORIGINS=https://yourapp.com,https://api.yourapp.com
export RATE_LIMIT_AI=30
export RATE_LIMIT_HEALTH=200
export RATE_LIMIT_INFO=100

# Logging and monitoring
export LOG_LEVEL=info
export METRICS_ENABLED=true
```

#### Load Balancing
```nginx
# nginx configuration
upstream ai_backend {
    server ai-service-1:8081;
    server ai-service-2:8081;
    server ai-service-3:8081;
}

server {
    listen 80;
    server_name api.yourapp.com;
    
    location / {
        proxy_pass http://ai_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

#### Monitoring & Health Checks
```bash
# Health check endpoint
curl -f http://localhost:8081/health || exit 1

# Prometheus metrics (if enabled)
curl http://localhost:8081/metrics

# Docker health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8081/health || exit 1
```

## 🧹 Artifact Management

### Automated Cleanup

The repository includes automated cleanup workflows to prevent GitHub storage quota issues:

#### Daily Cleanup Workflow
- **Schedule**: Runs daily at 2 AM UTC
- **Retention**: Removes artifacts older than 7 days
- **Workflow Runs**: Keeps only the last 50 workflow runs
- **Storage Management**: Automatically frees up storage space

#### Manual Cleanup Tools

For immediate cleanup when storage quota is exceeded:

##### PowerShell Script (Windows)
```powershell
# Run the cleanup script
.\cleanup-artifacts.ps1
```

##### Bash Script (Linux/macOS)
```bash
# Make executable and run
chmod +x cleanup-artifacts.sh
./cleanup-artifacts.sh
```

#### Storage Quota Tips

1. **Monitor Usage**: Check at [GitHub Billing Settings](https://github.com/settings/billing)
2. **Efficient Builds**: Artifacts only upload on main branch pushes
3. **Regular Cleanup**: Use manual scripts when quota is approached
4. **Retention Limits**: All artifacts expire after 7 days automatically

### GitHub CLI Cleanup (Quick Fix)

If you hit storage quota immediately:

```bash
# Install GitHub CLI if not already installed
# https://cli.github.com/

# Login to GitHub CLI
gh auth login

# List current artifacts
gh api repos/Ammar0144/ai/actions/artifacts

# Delete specific artifact by ID
gh api repos/Ammar0144/ai/actions/artifacts/ARTIFACT_ID -X DELETE

# Or use the automated cleanup scripts provided
```

## 🤝 Contributing

**We Welcome All Contributors!** Whether you're a beginner or experienced developer, your contributions help everyone learn.

### Ways to Contribute

#### 🐛 Report Issues
- Found a bug? Let us know!
- Something unclear? Ask questions!
- Every issue helps improve the project for learners

#### 💡 Share Ideas
- Suggest new features or improvements
- Share what you learned using this project
- Propose better ways to do things

#### 📖 Improve Documentation
- Fix typos or unclear explanations
- Add examples that helped you
- Translate docs to other languages

#### 💻 Code Contributions
1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature-name`
3. **Make your changes** and add tests
4. **Ensure all tests pass**: `go test ./...`
5. **Format your code**: `gofmt -s -w .`
6. **Commit your changes**: `git commit -am 'Add feature: description'`
7. **Push to the branch**: `git push origin feature-name`
8. **Submit a pull request** with a clear description

#### 🎓 Share Your Experience
- Write a blog post about using this project
- Create tutorial videos
- Share your learning journey
- Help others in issues and discussions

### 💬 Get Involved
- ⭐ Star the repo to show support
- 👀 Watch for updates
- 🍴 Fork and experiment
- 💬 Join discussions
- 📣 Share with others learning AI

**Remember**: No contribution is too small! Even fixing a typo helps the community.

## License

This project is open source and available under the MIT License.

## 🏷️ Version Information

- **Current Version**: 1.0.0
- **API Version**: v1
- **Swagger Version**: 1.0.0
- **Go Version**: 1.21+
- **Last Updated**: September 2025

### Initial Release (v1.0.0)
- ✅ Complete AI service with advanced rate limiting
- ✅ Comprehensive Swagger API documentation
- ✅ Production-ready microservices architecture
- ✅ Docker containerization and CI/CD pipeline
- ✅ Security features with IP-based access control
- ✅ Integration with LLM backend service

---

### 🚀 Quick Links
- [🔧 LLM Backend Repository](https://github.com/Ammar0144/llm)
- [📖 API Documentation](http://localhost:8081/swagger/) (when running)
- [🏥 Health Check](http://localhost:8081/health)
- [📊 Service Info](http://localhost:8081/)
