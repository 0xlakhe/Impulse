package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/seller"
)

type LLMProvider struct {
	client  *http.Client
	apiKey  string
	model   string
	baseURL string
}

func NewLLMProvider(client *http.Client, apiKey string, model string, baseURL string) *LLMProvider {
	return &LLMProvider{
		client:  client,
		apiKey:  apiKey,
		model:   model,
		baseURL: baseURL,
	}
}

func (p *LLMProvider) GenerateResponse(
	ctx context.Context, s seller.Seller, product *product.Product, history []Message,
) (*string, error) {

	messages := BuildRequestMessages(s, product, history)

	payload := responseRequest{
		Model:    p.model,
		Messages: messages,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/chat/completions", bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Printf("%+v", resp)
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("read AI API error response: %w", err)
		}
		return nil, &apiError{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}
	}

	// data,err:=io.ReadAll(resp.Body)
	// if err!=nil{
	// 	return nil,err
	// }

	var result responseResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if len(result.Choices) == 0 {
		return nil, errors.New("empty message from api response")
	}

	responseText := result.Choices[0].Message.Content

	return &responseText, nil
}
