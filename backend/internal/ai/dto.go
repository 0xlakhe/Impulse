package ai

type responseRequest struct{
	Messages []Message `json:"messages"`
	Model string `json:"model"`

}

type responseResponse struct{
	ID string `json:"id"`
	Choices []struct{
		Message Message `json:"message"`
	}`json:"choices"`

}
