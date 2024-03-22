package main

import (
	//"fmt"

	"encoding/json"
	"fmt"
	"os"

	"github.com/m4cd/aidevs/tasks"
)

func main() {
	var apikey string = os.Args[1]
	var taskname string = os.Args[2]
	var answer tasks.Answer

	TasksAPI := tasks.LoadAPI(apikey)
	token := tasks.Authorize(taskname, TasksAPI)
	task := tasks.GetTask(token, TasksAPI)

	systemMessage := tasks.CompletionMessage{
		Role: "system",
		Content: "Jesteś bloggerem piszącym o pizzy.",
	}

	for i := 0; i < len(task.Blog); i++ {
		content := fmt.Sprintf("Processing... \"%v\"",task.Blog[i])
		fmt.Println(content)
		userMessage := tasks.CompletionMessage{
			Role: "user",
			Content: content,
		}
		res := tasks.OpenAiCompletionRequest("https://api.openai.com/v1/chat/completions", TasksAPI, task, "gpt-4", systemMessage, userMessage)
		paragraph := res.Choices[0].Message.Content
		answer.Answer = append(answer.Answer, fmt.Sprintf("%v", paragraph))
	}

	jsonAnswer,_ := json.Marshal(answer)
    fmt.Println(string(jsonAnswer))
	

	tasks.SendAnswerBytes(token, TasksAPI, jsonAnswer)

}
