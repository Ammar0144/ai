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

// GetChatCompletion generates a chat completion based on conversation history
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

// GetComplete sends a completion request to the LLM server
func (s *AIService) GetComplete(prompt string, maxTokens int, temperature float64) (string, error) {
	log.Printf("GetComplete: processing prompt")

	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}

	// Set defaults
	if maxTokens == 0 {
		maxTokens = 150
	}
	if temperature == 0 {
		temperature = 0.7
	}

	// Prepare request
	request := LLMRequest{
		Prompt:      prompt,
		MaxLength:   maxTokens,
		Temperature: temperature,
		DoSample:    true,
	}

	// Call the complete endpoint which returns a different format
	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Printf("GetComplete: failed to marshal request: %v", err)
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	url := s.llmBaseURL + "/complete"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("GetComplete: failed to create request: %v", err)
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("GetComplete: request failed: %v", err)
		return "", fmt.Errorf("failed to connect to LLM server: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("GetComplete: failed to read response: %v", err)
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("GetComplete: API request failed with status %d: %s", resp.StatusCode, string(body))
		return "", fmt.Errorf("LLM server responded with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the response - /complete endpoint returns "completion" field
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("GetComplete: failed to unmarshal response: %v", err)
		return "", fmt.Errorf("failed to parse LLM response: %w", err)
	}

	// Extract completion field
	if completion, ok := response["completion"].(string); ok && completion != "" {
		log.Printf("GetComplete: successfully received completion")
		return completion, nil
	}

	log.Printf("GetComplete: no completion field in response")
	return "", fmt.Errorf("LLM server returned empty or invalid completion")
}

// GetGenerate sends a generation request to the LLM server
func (s *AIService) GetGenerate(prompt string, maxTokens int, temperature float64) (string, error) {
	log.Printf("GetGenerate: processing prompt")

	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}

	// Set defaults
	if maxTokens == 0 {
		maxTokens = 150
	}
	if temperature == 0 {
		temperature = 0.7
	}

	// Prepare request
	request := LLMRequest{
		Prompt:      prompt,
		MaxLength:   maxTokens,
		Temperature: temperature,
		DoSample:    true,
	}

	response, err := s.callLLMEndpoint("/generate", request)
	if err != nil {
		log.Printf("GetGenerate: LLM call failed: %v", err)
		return "", fmt.Errorf("generation failed: %w", err)
	}

	return response, nil
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

// GetModel returns the current model being used
func (s *AIService) GetModel() string {
	return s.model
}

// callLLMEndpointWithJSON makes a generic JSON request to LLM endpoints
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
