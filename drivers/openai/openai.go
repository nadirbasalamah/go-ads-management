package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-ads-management/utils"
	"net/http"
	"regexp"
)

type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenAIResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

type Choice struct {
	Index        int64       `json:"index"`
	Message      Message     `json:"message"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Message struct {
	Role    string      `json:"role"`
	Content string      `json:"content"`
	Refusal interface{} `json:"refusal"`
}

type Usage struct {
	PromptTokens            int64                   `json:"prompt_tokens"`
	CompletionTokens        int64                   `json:"completion_tokens"`
	TotalTokens             int64                   `json:"total_tokens"`
	PromptTokensDetails     PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens"`
}

type PromptTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens"`
}

type GenerateAdRequest struct {
	TargetAudience string `json:"target_audience" validate:"required"`
	ProductName    string `json:"product_name" validate:"required"`
	Platform       string `json:"platform" validate:"required"`
}

type GenerateAdResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// FOR DEMO ONLY
type ChannelResponse[T any] struct {
	Response T
	Error    error
}

type ReviewAdRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ReviewAdResponse struct {
	Review   string `json:"review"`
	Score    string `json:"score"`
	Feedback string `json:"feedback"`
}

func GenerateAd(ctx context.Context, request GenerateAdRequest) (GenerateAdResponse, error) {
	// request
	systemContent := fmt.Sprintf("You are marketing specialist for %s product", request.ProductName)

	adStructure := `{"title":"ad title","description":"ad description"}`

	userContent := fmt.Sprintf(
		"Write an %s advertisement in JSON format for a %s targeting %s, structured as %s",
		request.Platform,
		request.ProductName,
		request.TargetAudience,
		adStructure,
	)

	requestBody := OpenAIRequest{
		Model: utils.GetConfig("OPENAI_MODEL"),
		Messages: []Message{
			{
				Role:    "system",
				Content: systemContent,
			},
			{
				Role:    "user",
				Content: userContent,
			},
		},
	}

	response, err := sendRequest(ctx, requestBody)

	if err != nil {
		return GenerateAdResponse{}, err
	}

	responseContent := response.Choices[0].Message.Content

	re := regexp.MustCompile("(?s)```json\\s*(.*?)\\s*```")
	cleanedResponse := re.ReplaceAllString(responseContent, "$1")

	var adResponse GenerateAdResponse

	if err := json.Unmarshal([]byte(cleanedResponse), &adResponse); err != nil {
		return GenerateAdResponse{}, err
	}

	return adResponse, nil
}

func GenerateAdTrial(request GenerateAdRequest, ch chan ChannelResponse[GenerateAdResponse]) {
	// request
	systemContent := fmt.Sprintf("You are marketing specialist for %s product", request.ProductName)

	adStructure := `{"title":"ad title","description":"ad description"}`

	userContent := fmt.Sprintf(
		"Write an %s advertisement in JSON format for a %s targeting %s, structured as %s",
		request.Platform,
		request.ProductName,
		request.TargetAudience,
		adStructure,
	)

	requestBody := OpenAIRequest{
		Model: utils.GetConfig("OPENAI_MODEL"),
		Messages: []Message{
			{
				Role:    "system",
				Content: systemContent,
			},
			{
				Role:    "user",
				Content: userContent,
			},
		},
	}

	responsech := make(chan ChannelResponse[OpenAIResponse], 1)

	go sendRequestTrial(requestBody, responsech)

	response := <-responsech

	if response.Error != nil {
		ch <- ChannelResponse[GenerateAdResponse]{
			Error: response.Error,
		}
	}

	responseContent := response.Response.Choices[0].Message.Content

	re := regexp.MustCompile("(?s)```json\\s*(.*?)\\s*```")
	cleanedResponse := re.ReplaceAllString(responseContent, "$1")

	var adResponse GenerateAdResponse

	if err := json.Unmarshal([]byte(cleanedResponse), &adResponse); err != nil {
		ch <- ChannelResponse[GenerateAdResponse]{
			Error: err,
		}
	}

	ch <- ChannelResponse[GenerateAdResponse]{
		Response: adResponse,
		Error:    nil,
	}
}

func ReviewAd(ctx context.Context, request ReviewAdRequest) (ReviewAdResponse, error) {
	// request
	systemContent := "You are marketing specialist working with advertisement"

	reviewStructure := `{"review":"ad review","score":"ad score","feedback":"ad feedback"}`

	userContent := fmt.Sprintf(
		`Write a review of the given advertisement structured as %s. The score is from 1 up to 10.
		The advertisement is delimited with """.
		
		"""
		Title: %s
		Description: %s
		"""
		`,
		reviewStructure,
		request.Title,
		request.Description,
	)

	requestBody := OpenAIRequest{
		Model: utils.GetConfig("OPENAI_MODEL"),
		Messages: []Message{
			{
				Role:    "system",
				Content: systemContent,
			},
			{
				Role:    "user",
				Content: userContent,
			},
		},
	}

	response, err := sendRequest(ctx, requestBody)

	if err != nil {
		return ReviewAdResponse{}, err
	}

	responseContent := response.Choices[0].Message.Content

	var reviewResponse ReviewAdResponse

	if err := json.Unmarshal([]byte(responseContent), &reviewResponse); err != nil {
		return ReviewAdResponse{}, err
	}

	return reviewResponse, nil
}

func sendRequest(ctx context.Context, requestBody OpenAIRequest) (OpenAIResponse, error) {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return OpenAIResponse{}, errors.New("error marshalling request body")
	}

	// send the request
	apiKey := utils.GetConfig("OPENAI_API_KEY")
	endpoint := "https://api.openai.com/v1/chat/completions"
	body := bytes.NewBuffer(jsonData)

	// Create a new HTTP POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return OpenAIResponse{}, errors.New("error creating request")
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return OpenAIResponse{}, errors.New("error sending request")
	}
	defer resp.Body.Close()

	var data OpenAIResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return OpenAIResponse{}, errors.New("error parsing response")
	}

	return data, nil
}

func sendRequestTrial(requestBody OpenAIRequest, responsech chan ChannelResponse[OpenAIResponse]) {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		responsech <- ChannelResponse[OpenAIResponse]{
			Error: errors.New("error marshalling request body"),
		}
	}

	// send the request
	apiKey := utils.GetConfig("OPENAI_API_KEY")
	endpoint := "https://api.openai.com/v1/chat/completions"
	body := bytes.NewBuffer(jsonData)

	// Create a new HTTP POST request
	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		responsech <- ChannelResponse[OpenAIResponse]{
			Error: errors.New("error creating request"),
		}
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		responsech <- ChannelResponse[OpenAIResponse]{
			Error: errors.New("error sending request"),
		}
	}
	defer resp.Body.Close()

	var data OpenAIResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		responsech <- ChannelResponse[OpenAIResponse]{
			Error: errors.New("error parsing response"),
		}
	}

	responsech <- ChannelResponse[OpenAIResponse]{
		Error:    nil,
		Response: data,
	}
}
