package models

import "time"

// MessageRequest represents the incoming message from a user
type MessageRequest struct {
	Message string `json:"message" binding:"required"`
	UserID  string `json:"user_id,omitempty"`
}

// MessageResponse represents the response from the AI
type MessageResponse struct {
	Response  string    `json:"response"`
	UserID    string    `json:"user_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Model     string    `json:"model,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

// ChatMessage represents a single chat message
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest represents a chat completion request
type ChatCompletionRequest struct {
	Messages    []ChatMessage `json:"messages"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	Temperature float64       `json:"temperature,omitempty"`
	UserID      string        `json:"user_id,omitempty"`
}

// ChatCompletionResponse represents a chat completion response
type ChatCompletionResponse struct {
	Response  string    `json:"response"`
	UserID    string    `json:"user_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Model     string    `json:"model,omitempty"`
}

// EmbeddingsRequest represents an embeddings request
type EmbeddingsRequest struct {
	Text   string `json:"text"`
	UserID string `json:"user_id,omitempty"`
}

// EmbeddingsResponse represents an embeddings response
type EmbeddingsResponse struct {
	Embeddings string    `json:"embeddings"`
	UserID     string    `json:"user_id,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
}

// ClassificationRequest represents a text classification request
type ClassificationRequest struct {
	Text       string   `json:"text"`
	Categories []string `json:"categories"`
	UserID     string   `json:"user_id,omitempty"`
}

// ClassificationResponse represents a text classification response
type ClassificationResponse struct {
	Classification string    `json:"classification"`
	UserID         string    `json:"user_id,omitempty"`
	Timestamp      time.Time `json:"timestamp"`
}

// SummarizationRequest represents a text summarization request
type SummarizationRequest struct {
	Text      string `json:"text"`
	MaxLength int    `json:"max_length,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}

// SummarizationResponse represents a text summarization response
type SummarizationResponse struct {
	Summary   string    `json:"summary"`
	UserID    string    `json:"user_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// SentimentRequest represents a sentiment analysis request
type SentimentRequest struct {
	Text   string `json:"text"`
	UserID string `json:"user_id,omitempty"`
}

// SentimentResponse represents a sentiment analysis response
type SentimentResponse struct {
	Sentiment string    `json:"sentiment"`
	UserID    string    `json:"user_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// QuestionRequest represents a Q&A question request
type QuestionRequest struct {
	Question string `json:"question" binding:"required"`
	UserID   string `json:"user_id,omitempty"`
}
