package tasks

type TasksAPI struct {
	Endpoint string
	Token    string
	Answer   string
	Task     string
	Apikey   string
}

type AuthResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

type TaskResponse struct {
	Cookie string `json:"cookie"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
}

type AnswerResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Note string `json:"note"`
}

type Answer struct {
	Answer string `json:"answer"`
}
