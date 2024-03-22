package tasks

type ModerationResult struct {
	//Categories     ModerationResultCategories     `json:"categories"`
	//CategoryScores ModerationResultCategoryScores `json:"category_scores"`
	Flagged bool `json:"flagged"`
}

type ModerationResponse struct {
	ID      string             `json:"id"`
	Model   string             `json:"model"`
	Results []ModerationResult `json:"results"`
}

type ModerationInput struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}
