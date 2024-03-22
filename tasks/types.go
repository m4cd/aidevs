package tasks

type TasksAPI struct {
	Endpoint     string
	Token        string
	Answer       string
	Task         string
	Apikey       string
	OpenAiApikey string
}

type AuthResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

//===========================================
// moderation
type TaskResponse struct {
	//Cookie string `json:"cookie"`
	Input []string `json:"input"`
	Msg   string   `json:"msg"`
	//Token string `json:"token"`
	Code int `json:"code"`
}

type ModerationResult struct {
	//Categories     ModerationResultCategories     `json:"categories"`
	//CategoryScores ModerationResultCategoryScores `json:"category_scores"`
	Flagged        bool                 `json:"flagged"`
}

type ModerationResponse struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Results []ModerationResult `json:"results"`
}

type ModerationInput struct {
	Input []string `json:"input"`
	Model string `json:"model"`
}


//===========================================
// helloapi
/*
type TaskResponse struct {
	Cookie string `json:"cookie"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
}
*/
type AnswerResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Note string `json:"note"`
}

type Answer struct {
	//Answer string `json:"answer"`
	Answer []bool `json:"answer"`
}
