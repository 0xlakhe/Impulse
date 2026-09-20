package ai

import "fmt"

type apiError struct {
	StatusCode int
	Body       string
}

func (e *apiError) Error() string {
	return fmt.Sprintf("AI API returned status %d: %s", e.StatusCode, e.Body)
}
