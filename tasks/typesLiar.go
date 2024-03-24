package tasks

type QuestionAnswer struct {
	Code   int  `json:"code"`
	Msg    string `json:"msg"`
	Answer string `json:"answer"`
}