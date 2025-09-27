# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git (required for go install)
RUN apk add --no-cache git

# Install swag CLI tool for generating Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go mod files
COPY go.mod ./
# Copy go.sum if it exists (optional for projects without external dependencies)
COPY go.su[m] ./
RUN go mod download

# Copy source code
COPY . .

# Generate Swagger documentation
RUN swag init

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ai-server .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests and wget for health checks
RUN apk --no-cache add ca-certificates wget

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/ai-server .

# Copy the generated docs directory for Swagger UI
COPY --from=builder /app/docs ./docs/

# Expose port
EXPOSE 8081

# Set environment variables
ENV PORT=8081
ENV AI_MOCK_MODE=false

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8081/health || exit 1

# Run the application
CMD ["./ai-server"]