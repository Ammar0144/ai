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

// HandleComplete handles text completion requests
//
//	@Summary		Text completion
//	@Description	Complete text based on a given prompt. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.CompleteRequest		true	"Complete request"
//	@Success		200		{object}	models.CompleteResponse		"Successful text completion"
//	@Failure		400		{object}	models.ErrorResponse		"Bad request"
//	@Failure		429		{object}	models.ErrorResponse		"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse		"Internal server error"
//	@Router			/ai/complete [post]
func (h *AIHandler) HandleComplete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.CompleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Prompt == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Prompt cannot be empty")
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

	log.Printf("Received complete request from user %s", req.UserID)

	aiResponse, err := h.aiService.GetComplete(req.Prompt, maxTokens, temperature)
	if err != nil {
		log.Printf("Error getting completion: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get completion")
		return
	}

	response := models.CompleteResponse{
		Response:  aiResponse,
		UserID:    req.UserID,
		Timestamp: time.Now(),
		Model:     h.aiService.GetModel(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleGenerate handles text generation requests
//
//	@Summary		Text generation
//	@Description	Generate text based on a given prompt. Rate limited to 30 requests per minute per IP address.
//	@Tags			AI Processing
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.GenerateRequest		true	"Generate request"
//	@Success		200		{object}	models.GenerateResponse		"Successful text generation"
//	@Failure		400		{object}	models.ErrorResponse		"Bad request"
//	@Failure		429		{object}	models.ErrorResponse		"Rate limit exceeded"
//	@Failure		500		{object}	models.ErrorResponse		"Internal server error"
//	@Router			/ai/generate [post]
func (h *AIHandler) HandleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Prompt == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Prompt cannot be empty")
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

	log.Printf("Received generate request from user %s", req.UserID)

	aiResponse, err := h.aiService.GetGenerate(req.Prompt, maxTokens, temperature)
	if err != nil {
		log.Printf("Error getting generation: %v", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get generation")
		return
	}

	response := models.GenerateResponse{
		Response:  aiResponse,
		UserID:    req.UserID,
		Timestamp: time.Now(),
		Model:     h.aiService.GetModel(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleHealth handles health check requests
//
//	@Summary		Health check
//	@Description	Check the health status of the AI service. Rate limited to 200 requests per minute per IP address.
//	@Tags			System
//	@Produce		json
//	@Success		200	{object}	models.HealthResponse	"Service is healthy"
//	@Failure		429	{object}	models.ErrorResponse	"Rate limit exceeded"
//	@Router			/health [get]
func (h *AIHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	response := models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleModelInfo returns information about the current AI model
//
//	@Summary		Get model information
//	@Description	Get detailed information about the current AI model being used. Rate limited to 100 requests per minute per IP address.
//	@Tags			System
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"Model information"
//	@Failure		429	{object}	models.ErrorResponse	"Rate limit exceeded"
//	@Failure		500	{object}	models.ErrorResponse	"Internal server error"
//	@Router			/ai/model-info [get]
func (h *AIHandler) HandleModelInfo(w http.ResponseWriter, r *http.Request) {
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
