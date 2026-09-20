package ai

import (
	"strings"
	"testing"

	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/seller"
)

func TestBuildPrompt(t *testing.T) {
	s := seller.Seller{
		Name:         "Max",
		Persona:      "Vintage keyboard enthusiast",
		SystemPrompt: "Be friendly and knowledgeable.",
	}
	p := &product.Product{
		Name:        "IBM Model M",
		Description: "A classic mechanical keyboard.",
		Price:       12500,
		Category:    "Keyboards",
	}
	history := []Message{
		{
			Role:    string(RoleUser),
			Content: "hello",
		},
		{
			Role:    string(RoleAssistant),
			Content: "Absolutely!",
		},
	}
	result := BuildPrompt(s, p, history)

	expected := []string{
		"Max",
		"Vintage keyboard enthusiast",
		"Be friendly and knowledgeable.",
		"IBM Model M",
		"A classic mechanical keyboard.",
		"12500",
		"Keyboards",
		"hello",
		"Absolutely!",
	}

	for _, value := range expected {
		if !strings.Contains(result, value) {
			t.Errorf("expected prompt to contain %q, got : \n %s", value, result)
		}
	}
}

func TestBuildPromptWithoutProduct(t *testing.T) {
	s := seller.Seller{
		Name:         "Max",
		Persona:      "Vintage keyboard enthusiast",
		SystemPrompt: "Be friendly and knowledgeable",
	}

	history := []Message{
		{
			Role:    string(RoleUser),
			Content: "What keyboards do you recommend?",
		},
	}

	result := BuildPrompt(s, nil, history)

	if !strings.Contains(result, "Max") {
		t.Error("expected prompt to contain seller name")
	}

	if !strings.Contains(result, "What keyboards do you recommend?") {
		t.Error("expected prompt to contain conversation history")
	}

	if strings.Contains(result, "CURRENT PRODUCT:") {
		t.Error("prompt should not contain CURRENT PRODUCT when product is nil")
	}
}
