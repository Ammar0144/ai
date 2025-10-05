package main

// @title AI Service API
// @version 1.0.0
// @description A comprehensive Go-based AI service providing various artificial intelligence capabilities with advanced rate limiting, CORS support, and robust error handling.
// @termsOfService https://github.com/Ammar0144/ai

// @contact.name API Support
// @contact.url https://github.com/Ammar0144/ai/issues
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/Ammar0144/ai/docs" // Import docs for swagger
	"github.com/Ammar0144/ai/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Rate limiting structures
type RateLimiter struct {
	clients map[string]*ClientLimiter
	mutex   sync.RWMutex
}

type ClientLimiter struct {
	requests []time.Time
	mutex    sync.Mutex
}

// Global rate limiter instance
var rateLimiter = &RateLimiter{
	clients: make(map[string]*ClientLimiter),
}

// Rate limiting middleware
func rateLimitMiddleware(requestsPerMinute int) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Get client IP
			clientIP := getClientIP(r)

			// Check rate limit
			if !isAllowed(clientIP, requestsPerMinute) {
				// Add rate limit headers
				w.Header().Set("X-RateLimit-Limit", strconv.Itoa(requestsPerMinute))
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(time.Minute).Unix(), 10))

				// Set CORS headers for rate limit response
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

				http.Error(w, `{"error":"Rate limit exceeded","message":"Too many requests. Please try again later.","code":429}`, http.StatusTooManyRequests)
				return
			}

			// Add rate limit headers for successful requests
			remaining := getRemainingRequests(clientIP, requestsPerMinute)
			w.Header().Set("X-RateLimit-Limit", strconv.Itoa(requestsPerMinute))
			w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(remaining))
			w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(time.Minute).Unix(), 10))

			next(w, r)
		}
	}
}

// Get client IP address
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header (for proxies/load balancers)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// Take the first IP in case of multiple
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return xri
	}

	// Fallback to RemoteAddr
	ip := strings.Split(r.RemoteAddr, ":")[0]
	return ip
}

// Check if request is allowed based on rate limit
func isAllowed(clientIP string, limit int) bool {
	rateLimiter.mutex.Lock()
	defer rateLimiter.mutex.Unlock()

	// Get or create client limiter
	if rateLimiter.clients[clientIP] == nil {
		rateLimiter.clients[clientIP] = &ClientLimiter{
			requests: make([]time.Time, 0),
		}
	}

	client := rateLimiter.clients[clientIP]
	client.mutex.Lock()
	defer client.mutex.Unlock()

	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)

	// Remove requests older than 1 minute
	var validRequests []time.Time
	for _, reqTime := range client.requests {
		if reqTime.After(oneMinuteAgo) {
			validRequests = append(validRequests, reqTime)
		}
	}
	client.requests = validRequests

	// Check if limit exceeded
	if len(client.requests) >= limit {
		return false
	}

	// Add current request
	client.requests = append(client.requests, now)
	return true
}

// Get remaining requests for a client
func getRemainingRequests(clientIP string, limit int) int {
	rateLimiter.mutex.RLock()
	defer rateLimiter.mutex.RUnlock()

	if rateLimiter.clients[clientIP] == nil {
		return limit - 1
	}

	client := rateLimiter.clients[clientIP]
	client.mutex.Lock()
	defer client.mutex.Unlock()

	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)

	// Count valid requests
	validCount := 0
	for _, reqTime := range client.requests {
		if reqTime.After(oneMinuteAgo) {
			validCount++
		}
	}

	remaining := limit - validCount
	if remaining < 0 {
		remaining = 0
	}
	return remaining
}

// Cleanup function to remove old client data (run periodically)
func cleanupOldClients() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			rateLimiter.mutex.Lock()
			now := time.Now()
			fiveMinutesAgo := now.Add(-5 * time.Minute)

			for clientIP, client := range rateLimiter.clients {
				client.mutex.Lock()
				hasRecentRequests := false
				for _, reqTime := range client.requests {
					if reqTime.After(fiveMinutesAgo) {
						hasRecentRequests = true
						break
					}
				}
				if !hasRecentRequests {
					delete(rateLimiter.clients, clientIP)
				}
				client.mutex.Unlock()
			}
			rateLimiter.mutex.Unlock()
		}
	}()
}

// CORS middleware
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// Combined middleware (CORS + Rate Limiting)
func protectedHandler(handler http.HandlerFunc, rateLimit int) http.HandlerFunc {
	return corsMiddleware(rateLimitMiddleware(rateLimit)(handler))
}

func main() {
	// Get port from environment variable or use default
	port := "8081"

	// Start cleanup routine for rate limiter
	cleanupOldClients()

	// Create handlers
	aiHandler := handlers.NewAIHandler()

	// Set up AI group routes with rate limiting and CORS
	// AI endpoints: 30 requests per minute (resource intensive)
	http.HandleFunc("/ai/chat/completions", protectedHandler(aiHandler.HandleChatCompletion, 30))
	http.HandleFunc("/ai/complete", protectedHandler(aiHandler.HandleComplete, 30))
	http.HandleFunc("/ai/generate", protectedHandler(aiHandler.HandleGenerate, 30))
	http.HandleFunc("/ai/model-info", protectedHandler(aiHandler.HandleModelInfo, 100)) // Less intensive

	// Health endpoint: Higher limit for monitoring
	http.HandleFunc("/health", protectedHandler(aiHandler.HandleHealth, 200))

	// Swagger JSON spec endpoint
	http.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.json")
	})

	// Swagger documentation UI
	http.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("doc.json"), // Relative path to swagger spec
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
	))

	// API documentation page
	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/index.html")
	})

	// Root handler for undefined routes with rate limiting
	http.HandleFunc("/", protectedHandler(func(w http.ResponseWriter, r *http.Request) {
		// Only handle root path and undefined routes
		if r.URL.Path == "/" {
			// Set CORS headers for root
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Provide API information at root
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
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
			}`))
			return
		}

		// For all other undefined routes, check if it's a swagger route
		if strings.HasPrefix(r.URL.Path, "/swagger/") {
			return // Let swagger handler deal with it
		}

		// Handle 404 for truly unknown routes
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.NotFound(w, r)
	}, 100)) // Root endpoint: 100 requests per minute

	// Start server
	log.Printf("Starting AI server on port %s with rate limiting enabled", port)
	log.Printf("📚 Documentation: http://localhost:%s/docs", port)
	log.Printf("🚀 Swagger UI: http://localhost:%s/swagger/", port)
	log.Printf("🏥 Health check: http://localhost:%s/health", port)
	log.Printf("📄 API Documentation: http://localhost:%s/swagger/doc.json", port)
	log.Printf("🌐 API Info: http://localhost:%s/", port)
	log.Printf("")
	log.Printf("⚡ Rate Limits Applied:")
	log.Printf("  • AI Endpoints: 30 requests/minute per IP")
	log.Printf("  • Health Endpoint: 200 requests/minute per IP")
	log.Printf("  • Model Info: 100 requests/minute per IP")
	log.Printf("  • Root Endpoint: 100 requests/minute per IP")
	log.Printf("")
	log.Printf("AI Endpoints available:")
	log.Printf("  - Chat Completions: http://localhost:%s/ai/chat/completions", port)
	log.Printf("  - Text Completion: http://localhost:%s/ai/complete", port)
	log.Printf("  - Text Generation: http://localhost:%s/ai/generate", port)
	log.Printf("  - Model Information: http://localhost:%s/ai/model-info", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// deploy 1
