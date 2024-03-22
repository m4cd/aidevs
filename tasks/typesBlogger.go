package tasks

type CompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	Model    string              `json:"model"`
	Messages []CompletionMessage `json:"messages"`
}

type CompletionChoice struct {
	Message CompletionMessage `json:"message"`
}

type CompletionResponse struct {
	Choices []CompletionChoice `json:"choices"`
}
