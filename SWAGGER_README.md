# AI Service API Documentation

This document provides instructions for setting up and using the Swagger documentation for the AI Service API.

## 🚀 Quick Start

### Prerequisites

1. **Install Swagger CLI tool** (swag):
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

### Generate Swagger Documentation

Run the following command in the project root directory:

```bash
swag init
```

This will generate:
- `docs/docs.go` - Go documentation
- `docs/swagger.json` - JSON specification  
- `docs/swagger.yaml` - YAML specification

### Start the Server

```bash
go run main.go
```

## 📚 Access Documentation

Once the server is running (default port 8081), you can access:

- **📖 Custom Documentation Page**: http://localhost:8081/docs
- **🚀 Interactive Swagger UI**: http://localhost:8081/swagger/
- **📄 JSON Specification**: http://localhost:8081/swagger/doc.json
- **🏥 Health Check**: http://localhost:8081/health

## 🔧 API Endpoints

### Health Endpoints
- `GET /health` - Service health check

### AI Processing Endpoints
- `POST /ai/chat/completions` - Chat completion with conversation history
- `POST /ai/embeddings` - Generate text embeddings
- `POST /ai/classifications` - Classify text into categories
- `POST /ai/summarization` - Summarize text content
- `POST /ai/sentiment` - Analyze text sentiment

### Legacy Endpoints
- `POST /api/message` - Simple message processing (legacy)

## 🎯 Quick Test Examples

### Health Check
```bash
curl http://localhost:8081/health
```

### Chat Completion
```bash
curl -X POST "http://localhost:8081/ai/chat/completions" \
     -H "Content-Type: application/json" \
     -d '{
       "messages": [
         {"role": "user", "content": "What is artificial intelligence?"}
       ],
       "max_tokens": 150,
       "temperature": 0.7
     }'
```

### Text Classification
```bash
curl -X POST "http://localhost:8081/ai/classifications" \
     -H "Content-Type: application/json" \
     -d '{
       "text": "This product is absolutely amazing!",
       "categories": ["positive", "negative", "neutral"]
     }'
```

### Sentiment Analysis
```bash
curl -X POST "http://localhost:8081/ai/sentiment" \
     -H "Content-Type: application/json" \
     -d '{
       "text": "I love this new feature!"
     }'
```

### Text Summarization
```bash
curl -X POST "http://localhost:8081/ai/summarization" \
     -H "Content-Type: application/json" \
     -d '{
       "text": "Artificial intelligence (AI) is intelligence demonstrated by machines, in contrast to the natural intelligence displayed by humans and animals. Leading AI textbooks define the field as the study of intelligent agents...",
       "max_length": 100
     }'
```

## 🔄 Regenerating Documentation

Whenever you update the API endpoints or modify the Swagger annotations, regenerate the documentation:

```bash
swag init
```

## 📝 Swagger Annotations

The API uses Swagger annotations in the handler functions. Example:

```go
// HandleMessage processes incoming messages and returns AI responses
//
//	@Summary		Process a message
//	@Description	Send a message to the AI and get a response
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.MessageRequest	true	"Message request"
//	@Success		200		{object}	models.MessageResponse	"Successful response"
//	@Failure		400		{object}	models.ErrorResponse	"Bad request"
//	@Failure		500		{object}	models.ErrorResponse	"Internal server error"
//	@Router			/ai/message [post]
func (h *AIHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
    // Implementation...
}
```

## 🛠 Customization

### Update API Information

Edit the header in `docs/docs.go`:

```go
//	@title			AI Service API
//	@version		1.0.0
//	@description	Your custom description
//	@host			localhost:8081
//	@BasePath		/
```

### Add New Endpoints

1. Add Swagger annotations to your handler function
2. Run `swag init` to regenerate documentation
3. Restart the server

## 🌐 Production Deployment

For production deployment:

1. Update the host in `docs/docs.go`:
   ```go
   //	@host		your-domain.com:8081
   ```

2. Regenerate documentation:
   ```bash
   swag init
   ```

3. Deploy with proper HTTPS configuration

## 📊 Features

- ✅ Complete API documentation
- ✅ Interactive testing with Swagger UI
- ✅ JSON and YAML specifications
- ✅ Request/response examples
- ✅ Error handling documentation
- ✅ Type-safe models
- ✅ Custom HTML documentation page

## 🔍 Troubleshooting

### Common Issues

1. **Swagger generation fails**:
   - Ensure `swag` CLI is installed
   - Check that all imports are correct
   - Verify annotation syntax

2. **Documentation not updating**:
   - Run `swag init` after changes
   - Restart the server
   - Clear browser cache

3. **Missing endpoints in documentation**:
   - Ensure handlers have proper Swagger annotations
   - Check that the route is properly registered

## 📈 Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Swagger UI    │    │  Custom Docs    │    │   JSON/YAML     │
│  /swagger/      │    │    /docs        │    │ Specifications  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │   AI Service    │
                    │     API         │
                    │   (Port 8081)   │
                    └─────────────────┘
                                 │
                    ┌─────────────────┐
                    │   LLM Server    │
                    │  (Port 8082)    │
                    │   DistilGPT-2   │
                    └─────────────────┘
```

## 📚 Additional Resources

- [Swagger/OpenAPI Specification](https://swagger.io/specification/)
- [swaggo Documentation](https://github.com/swaggo/swag)
- [Go Swagger Annotations](https://github.com/swaggo/swag#declarative-comments-format)

---

**Happy coding! 🚀**