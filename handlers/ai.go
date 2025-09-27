package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Ammar0144/ai/models"
	"github.com/Ammar0144/ai/services"
)

// AIHandler handles AI-related HTTP requests
type AIHandler struct {
	aiService *services.AIService
}

// NewAIHandler creates a new AI handler
func NewAIHandler() *AIHandler {
	return &AIHandler{
		aiService: services.NewAIService(),
	}
}

// HandleMessage processes incoming messages and returns AI responses
//
//	@Summary		Process a message
//	@Description	Send a message to the AI and get a response. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.MessageRequest	true	"Message request"
//	@Success		200		{object}	models.MessageResponse	"Successful response"
//	@Failure		400		{object}	models.ErrorResponse	"Bad request"
//	@Failure		429		{object}	models.ErrorResponse	"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse	"Internal server error"
//	@Router			/ai/message [post]
func (h *AIHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Parse request body
	var req models.MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate message
	if req.Message == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Message cannot be empty")
		return
	}

	log.Printf("Received message from user %s: %s", req.UserID, req.Message)

	// Get AI response
	aiResponse, err := h.aiService.GetResponse(req.Message)
	if err != nil {
		log.Printf("Error getting AI response: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get AI response")
		return
	}

	// Create response
	response := models.MessageResponse{
		Response:  aiResponse,
		UserID:    req.UserID,
		Timestamp: time.Now(),
		Model:     h.aiService.GetModel(),
	}

	log.Printf("Sending response to user %s: %s", req.UserID, aiResponse)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleChatCompletion handles chat completion requests
//
//	@Summary		Chat completion
//	@Description	Generate a chat completion based on conversation history. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.ChatCompletionRequest	true	"Chat completion request"
//	@Success		200		{object}	models.ChatCompletionResponse	"Successful chat completion"
//	@Failure		400		{object}	models.ErrorResponse			"Bad request"
//	@Failure		429		{object}	models.ErrorResponse			"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse			"Internal server error"
//	@Router			/ai/chat/completions [post]
func (h *AIHandler) HandleChatCompletion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.ChatCompletionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if len(req.Messages) == 0 {
		h.sendErrorResponse(w, http.StatusBadRequest, "Messages cannot be empty")
		return
	}

	maxTokens := req.MaxTokens
	if maxTokens == 0 {
		maxTokens = 150
	}

	temperature := req.Temperature
	if temperature == 0 {
		temperature = 0.7
	}

	log.Printf("Received chat completion request from user %s with %d messages", req.UserID, len(req.Messages))

	aiResponse, err := h.aiService.GetChatCompletion(req.Messages, maxTokens, temperature)
	if err != nil {
		log.Printf("Error getting chat completion: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get chat completion")
		return
	}

	response := models.ChatCompletionResponse{
		Response:  aiResponse,
		UserID:    req.UserID,
		Timestamp: time.Now(),
		Model:     h.aiService.GetModel(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleEmbeddings handles text embeddings requests
//
//	@Summary		Generate text embeddings
//	@Description	Generate vector embeddings for the given text. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.EmbeddingsRequest	true	"Embeddings request"
//	@Success		200		{object}	models.EmbeddingsResponse	"Successful embeddings generation"
//	@Failure		400		{object}	models.ErrorResponse		"Bad request"
//	@Failure		429		{object}	models.ErrorResponse		"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse		"Internal server error"
//	@Router			/ai/embeddings [post]
func (h *AIHandler) HandleEmbeddings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.EmbeddingsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Text == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	log.Printf("Received embeddings request from user %s", req.UserID)

	embeddings, err := h.aiService.GetEmbeddings(req.Text)
	if err != nil {
		log.Printf("Error getting embeddings: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get embeddings")
		return
	}

	response := models.EmbeddingsResponse{
		Embeddings: embeddings,
		UserID:     req.UserID,
		Timestamp:  time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleClassification handles text classification requests
//
//	@Summary		Classify text
//	@Description	Classify text into predefined categories. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.ClassificationRequest	true	"Classification request"
//	@Success		200		{object}	models.ClassificationResponse	"Successful text classification"
//	@Failure		400		{object}	models.ErrorResponse			"Bad request"
//	@Failure		429		{object}	models.ErrorResponse			"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse			"Internal server error"
//	@Router			/ai/classifications [post]
func (h *AIHandler) HandleClassification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.ClassificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Text == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	if len(req.Categories) == 0 {
		h.sendErrorResponse(w, http.StatusBadRequest, "Categories cannot be empty")
		return
	}

	log.Printf("Received classification request from user %s", req.UserID)

	classification, err := h.aiService.GetClassification(req.Text, req.Categories)
	if err != nil {
		log.Printf("Error getting classification: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get classification")
		return
	}

	response := models.ClassificationResponse{
		Classification: classification,
		UserID:         req.UserID,
		Timestamp:      time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleSummarization handles text summarization requests
//
//	@Summary		Summarize text
//	@Description	Generate a summary of the provided text. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SummarizationRequest	true	"Summarization request"
//	@Success		200		{object}	models.SummarizationResponse	"Successful text summarization"
//	@Failure		400		{object}	models.ErrorResponse			"Bad request"
//	@Failure		429		{object}	models.ErrorResponse			"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse			"Internal server error"
//	@Router			/ai/summarization [post]
func (h *AIHandler) HandleSummarization(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SummarizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Text == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	maxLength := req.MaxLength
	if maxLength == 0 {
		maxLength = 100
	}

	log.Printf("Received summarization request from user %s", req.UserID)

	summary, err := h.aiService.GetSummarization(req.Text, maxLength)
	if err != nil {
		log.Printf("Error getting summarization: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get summarization")
		return
	}

	response := models.SummarizationResponse{
		Summary:   summary,
		UserID:    req.UserID,
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleSentiment handles sentiment analysis requests
//
//	@Summary		Analyze sentiment
//	@Description	Analyze the sentiment of the provided text. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SentimentRequest		true	"Sentiment request"
//	@Success		200		{object}	models.SentimentResponse	"Successful sentiment analysis"
//	@Failure		400		{object}	models.ErrorResponse		"Bad request"
//	@Failure		429		{object}	models.ErrorResponse		"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse		"Internal server error"
//	@Router			/ai/sentiment [post]
func (h *AIHandler) HandleSentiment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SentimentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Text == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	log.Printf("Received sentiment analysis request from user %s", req.UserID)

	sentiment, err := h.aiService.GetSentiment(req.Text)
	if err != nil {
		log.Printf("Error getting sentiment: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get sentiment analysis")
		return
	}

	response := models.SentimentResponse{
		Sentiment: sentiment,
		UserID:    req.UserID,
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleHealth provides a health check endpoint
//
//	@Summary		Health check
//	@Description	Returns the health status of the AI service. Rate limited to 200 requests per minute per IP address.
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	models.HealthResponse	"Service is healthy"
//	@Failure		429	{object}	models.ErrorResponse	"Rate limit exceeded"
//	@Router			/health [get]
func (h *AIHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	response := models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleModelInfo returns detailed model information
//
//	@Summary		Get model information
//	@Description	Returns detailed information about the current AI model. Rate limited to 100 requests per minute per IP address.
//	@Tags			Model Info
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"Model information"
//	@Failure		429	{object}	models.ErrorResponse	"Rate limit exceeded"
//	@Failure		500	{object}	models.ErrorResponse	"Internal server error"
//	@Router			/ai/model-info [get]
func (h *AIHandler) HandleModelInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get model information from AI service
	modelInfo, err := h.aiService.GetModelInfo()
	if err != nil {
		log.Printf("Error getting model info: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get model information")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(modelInfo)
}

// HandleAsk processes Q&A requests with enhanced response cleaning
//
//	@Summary		Ask a question
//	@Description	Send a question to the AI and get a cleaned, intelligent response. Rate limited to 30 requests per minute per IP address.
//	@Tags			Q&A
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.QuestionRequest	true	"Question request"
//	@Success		200		{object}	map[string]interface{}	"Q&A response"
//	@Failure		400		{object}	models.ErrorResponse	"Bad request"
//	@Failure		429		{object}	models.ErrorResponse	"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse	"Internal server error"
//	@Router			/ai/ask [post]
func (h *AIHandler) HandleAsk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Parse request body
	var req struct {
		Question  string `json:"question"`
		MaxLength int    `json:"max_length,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate question
	if req.Question == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Question cannot be empty")
		return
	}

	// Set default max length if not provided
	if req.MaxLength == 0 {
		req.MaxLength = 150
	}

	// Get answer from AI service
	response, err := h.aiService.GetAnswer(req.Question, req.MaxLength)
	if err != nil {
		log.Printf("Error getting answer: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to process question")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// sendErrorResponse sends a JSON error response
func (h *AIHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := models.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Code:    statusCode,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
