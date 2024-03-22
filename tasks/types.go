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

type AnswerResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Note string `json:"note"`
}

type Answer struct {
	//helloapi
	//Answer string `json:"answer"`

	//moderation
	//Answer []bool `json:"answer"`

	//blogger
	Answer []string `json:"answer"`
}

// Type Alias
type TaskResponse = TaskResponseBlogger

// blogger
type TaskResponseBlogger struct {
	//Cookie string `json:"cookie"`
	Blog []string `json:"blog"`
	Msg  string   `json:"msg"`
	//Token string `json:"token"`
	Code int `json:"code"`
}

// moderation
type TaskResponseModeration struct {
	//Cookie string `json:"cookie"`
	Input []string `json:"input"`
	Msg   string   `json:"msg"`
	//Token string `json:"token"`
	Code int `json:"code"`
}


// helloapi
type TaskResponseHelloapi struct {
	Cookie string `json:"cookie"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
}