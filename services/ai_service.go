package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Ammar0144/ai/models"
)

// AIService handles communication with AI providers
type AIService struct {
	client     *http.Client
	llmBaseURL string
	model      string
}

// LLMRequest represents the request to local LLM server
type LLMRequest struct {
	Prompt      string  `json:"prompt"`
	MaxLength   int     `json:"max_length,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	DoSample    bool    `json:"do_sample,omitempty"`
}

// LLMResponse represents the response from local LLM server
type LLMResponse struct {
	GeneratedText string `json:"generated_text"`
	Prompt        string `json:"prompt"`
}

// NewAIService creates a new AI service instance
func NewAIService() *AIService {
	// Get LLM server URL from environment variable or use default localhost
	llmURL := "http://llm-server:8082"

	log.Printf("AI Service initialized with LLM server: %s", llmURL)

	return &AIService{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		llmBaseURL: llmURL,
		model:      "distilgpt2",
	}
}

// GetResponse sends a message to the AI service and returns the response
func (s *AIService) GetResponse(message string) (string, error) {
	if message == "" {
		log.Printf("GetResponse: empty message provided")
		return "", fmt.Errorf("message cannot be empty")
	}

	log.Printf("GetResponse: processing message of length %d", len(message))

	// Use local LLM service
	response, err := s.getLLMResponse(message)
	if err != nil {
		log.Printf("GetResponse: LLM request failed: %v", err)
		return "", fmt.Errorf("failed to get AI response: %w", err)
	}

	log.Printf("GetResponse: successfully generated response of length %d", len(response))
	return response, nil
}

// getLLMResponse makes a request to local LLM server
func (s *AIService) getLLMResponse(message string) (string, error) {
	log.Printf("getLLMResponse: preparing request for LLM server %s", s.llmBaseURL)

	// Prepare the request
	request := LLMRequest{
		Prompt:      message,
		MaxLength:   100,
		Temperature: 0.7,
		TopP:        0.9,
		DoSample:    true,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Printf("getLLMResponse: failed to marshal request: %v", err)
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := s.llmBaseURL + "/generate"
	log.Printf("getLLMResponse: making request to %s", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("getLLMResponse: failed to create request: %v", err)
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("getLLMResponse: request failed: %v", err)
		return "", fmt.Errorf("failed to connect to LLM server at %s: %w", s.llmBaseURL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("getLLMResponse: failed to read response: %v", err)
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	log.Printf("getLLMResponse: received response with status %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		log.Printf("getLLMResponse: API request failed with status %d: %s", resp.StatusCode, string(body))
		return "", fmt.Errorf("LLM server responded with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var response LLMResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("getLLMResponse: failed to unmarshal response: %v", err)
		return "", fmt.Errorf("failed to parse LLM response: %w", err)
	}

	// Clean up the response text
	generatedText := response.GeneratedText
	if generatedText == "" {
		log.Printf("getLLMResponse: received empty generated text")
		return "", fmt.Errorf("LLM server returned empty generated text")
	}

	log.Printf("getLLMResponse: successfully generated text of length %d", len(generatedText))
	return generatedText, nil
}

// GetChatCompletion handles chat completion requests
func (s *AIService) GetChatCompletion(messages []models.ChatMessage, maxTokens int, temperature float64) (string, error) {
	if len(messages) == 0 {
		log.Printf("GetChatCompletion: no messages provided")
		return "", fmt.Errorf("messages cannot be empty")
	}

	log.Printf("GetChatCompletion: processing %d messages", len(messages))

	// Use the new chat/completions endpoint with proper structure
	request := map[string]interface{}{
		"messages":    messages,
		"max_tokens":  maxTokens,
		"temperature": temperature,
	}

	response, err := s.callLLMEndpointWithJSON("/chat/completions", request)
	if err != nil {
		log.Printf("GetChatCompletion: LLM call failed: %v", err)
		return "", fmt.Errorf("chat completion failed: %w", err)
	}

	// Parse response for content field
	if content, ok := response["content"].(string); ok {
		log.Printf("GetChatCompletion: successfully generated response")
		return strings.TrimSpace(content), nil
	}

	// Fallback to generated_text if available
	if generatedText, ok := response["generated_text"].(string); ok {
		return strings.TrimSpace(generatedText), nil
	}

	return "", fmt.Errorf("unexpected response format from LLM server")
}

// GetEmbeddings handles text embeddings using dedicated LLM endpoint
func (s *AIService) GetEmbeddings(text string) (string, error) {
	if text == "" {
		log.Printf("GetEmbeddings: empty text provided")
		return "", fmt.Errorf("text cannot be empty")
	}

	log.Printf("GetEmbeddings: generating embeddings for text of length %d", len(text))

	// Use the new dedicated embeddings endpoint
	request := map[string]interface{}{
		"text": text,
	}

	response, err := s.callLLMEndpointWithJSON("/embeddings", request)
	if err != nil {
		log.Printf("GetEmbeddings: LLM call failed: %v", err)
		return "", fmt.Errorf("embeddings generation failed: %w", err)
	}

	// Parse response for embeddings field
	if embeddings, ok := response["embeddings"].(string); ok {
		log.Printf("GetEmbeddings: successfully generated embeddings")
		return embeddings, nil
	}

	return "", fmt.Errorf("unexpected response format from embeddings endpoint")
}

// GetClassification handles text classification
func (s *AIService) GetClassification(text string, categories []string) (string, error) {
	if text == "" {
		log.Printf("GetClassification: empty text provided")
		return "", fmt.Errorf("text cannot be empty")
	}

	if len(categories) == 0 {
		log.Printf("GetClassification: no categories provided")
		return "", fmt.Errorf("categories cannot be empty")
	}

	log.Printf("GetClassification: classifying text of length %d into %d categories", len(text), len(categories))

	// Use the new classify endpoint
	request := map[string]interface{}{
		"text":   text,
		"labels": categories,
	}

	response, err := s.callLLMEndpointWithJSON("/classify", request)
	if err != nil {
		log.Printf("GetClassification: LLM call failed: %v", err)
		return "", fmt.Errorf("classification failed: %w", err)
	}

	// Parse response for prediction field
	if prediction, ok := response["prediction"].(string); ok {
		log.Printf("GetClassification: successfully classified text as %s", prediction)
		return prediction, nil
	}

	return "", fmt.Errorf("unexpected response format from classification endpoint")
}

// GetSummarization handles text summarization
func (s *AIService) GetSummarization(text string, maxLength int) (string, error) {
	if text == "" {
		log.Printf("GetSummarization: empty text provided")
		return "", fmt.Errorf("text cannot be empty")
	}

	log.Printf("GetSummarization: summarizing text of length %d with max length %d", len(text), maxLength)

	// Use the new summarize endpoint
	request := map[string]interface{}{
		"text":       text,
		"max_length": maxLength,
	}

	response, err := s.callLLMEndpointWithJSON("/summarize", request)
	if err != nil {
		log.Printf("GetSummarization: LLM call failed: %v", err)
		return "", fmt.Errorf("summarization failed: %w", err)
	}

	// Parse response for summary field
	if summary, ok := response["summary"].(string); ok {
		log.Printf("GetSummarization: successfully generated summary")
		return summary, nil
	}

	return "", fmt.Errorf("unexpected response format from summarization endpoint")
}

// GetSentiment handles sentiment analysis
func (s *AIService) GetSentiment(text string) (string, error) {
	if text == "" {
		log.Printf("GetSentiment: empty text provided")
		return "", fmt.Errorf("text cannot be empty")
	}

	log.Printf("GetSentiment: analyzing sentiment of text with length %d", len(text))

	// Use the new sentiment endpoint
	request := map[string]interface{}{
		"text": text,
	}

	response, err := s.callLLMEndpointWithJSON("/sentiment", request)
	if err != nil {
		log.Printf("GetSentiment: LLM call failed: %v", err)
		return "", fmt.Errorf("sentiment analysis failed: %w", err)
	}

	// Parse response for sentiment field
	if sentiment, ok := response["sentiment"].(string); ok {
		log.Printf("GetSentiment: successfully analyzed sentiment as %s", sentiment)
		return sentiment, nil
	}

	return "", fmt.Errorf("unexpected response format from sentiment endpoint")
}

// callLLMEndpointWithJSON is a helper method for endpoints that return JSON responses
func (s *AIService) callLLMEndpointWithJSON(endpoint string, request interface{}) (map[string]interface{}, error) {
	log.Printf("callLLMEndpointWithJSON: making request to %s%s", s.llmBaseURL, endpoint)

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Printf("callLLMEndpointWithJSON: failed to marshal request: %v", err)
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := s.llmBaseURL + endpoint
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("callLLMEndpointWithJSON: failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("callLLMEndpointWithJSON: request failed: %v", err)
		return nil, fmt.Errorf("failed to connect to LLM server at %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("callLLMEndpointWithJSON: failed to read response: %v", err)
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	log.Printf("callLLMEndpointWithJSON: received response with status %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		log.Printf("callLLMEndpointWithJSON: API request failed with status %d: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("LLM server responded with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the JSON response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("callLLMEndpointWithJSON: failed to unmarshal response: %v", err)
		log.Printf("callLLMEndpointWithJSON: raw response body: %s", string(body))
		return nil, fmt.Errorf("failed to parse LLM response: %w", err)
	}

	log.Printf("callLLMEndpointWithJSON: successfully received JSON response")
	return response, nil
}

// callLLMEndpoint is a helper method to make requests to different LLM endpoints
func (s *AIService) callLLMEndpoint(endpoint string, request LLMRequest) (string, error) {
	log.Printf("callLLMEndpoint: making request to %s%s", s.llmBaseURL, endpoint)

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Printf("callLLMEndpoint: failed to marshal request: %v", err)
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := s.llmBaseURL + endpoint
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("callLLMEndpoint: failed to create request: %v", err)
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("callLLMEndpoint: request failed: %v", err)
		return "", fmt.Errorf("failed to connect to LLM server at %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("callLLMEndpoint: failed to read response: %v", err)
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	log.Printf("callLLMEndpoint: received response with status %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		log.Printf("callLLMEndpoint: API request failed with status %d: %s", resp.StatusCode, string(body))
		return "", fmt.Errorf("LLM server responded with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var response LLMResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("callLLMEndpoint: failed to unmarshal response: %v", err)
		log.Printf("callLLMEndpoint: raw response body: %s", string(body))
		return "", fmt.Errorf("failed to parse LLM response: %w", err)
	}

	// Clean up the response text
	generatedText := response.GeneratedText
	if generatedText == "" {
		log.Printf("callLLMEndpoint: received empty generated text")
		return "", fmt.Errorf("LLM server returned empty generated text")
	}

	log.Printf("callLLMEndpoint: successfully received response of length %d", len(generatedText))
	return generatedText, nil
}

// GetModelInfo returns detailed information about the current model
func (s *AIService) GetModelInfo() (map[string]interface{}, error) {
	log.Printf("GetModelInfo: requesting model information")

	// Call the model-info endpoint
	url := s.llmBaseURL + "/model-info"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("GetModelInfo: failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("GetModelInfo: request failed: %v", err)
		return nil, fmt.Errorf("failed to connect to LLM server: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("GetModelInfo: failed to read response: %v", err)
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("GetModelInfo: API request failed with status %d: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("LLM server responded with status %d: %s", resp.StatusCode, string(body))
	}

	var modelInfo map[string]interface{}
	if err := json.Unmarshal(body, &modelInfo); err != nil {
		log.Printf("GetModelInfo: failed to unmarshal response: %v", err)
		return nil, fmt.Errorf("failed to parse model info: %w", err)
	}

	log.Printf("GetModelInfo: successfully retrieved model information")
	return modelInfo, nil
}

// GetAnswer handles enhanced Q&A using the ask endpoint
func (s *AIService) GetAnswer(question string, maxLength int) (map[string]interface{}, error) {
	if question == "" {
		log.Printf("GetAnswer: empty question provided")
		return nil, fmt.Errorf("question cannot be empty")
	}

	log.Printf("GetAnswer: processing question of length %d", len(question))

	// Use the enhanced ask endpoint
	request := map[string]interface{}{
		"question":   question,
		"max_length": maxLength,
	}

	response, err := s.callLLMEndpointWithJSON("/ask", request)
	if err != nil {
		log.Printf("GetAnswer: LLM call failed: %v", err)
		return nil, fmt.Errorf("Q&A processing failed: %w", err)
	}

	log.Printf("GetAnswer: successfully processed question")
	return response, nil
}

// GetModel returns the current model being used
func (s *AIService) GetModel() string {
	return s.model
}
