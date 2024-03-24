package main

import (
	//"fmt"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/m4cd/aidevs/tasks"
)

func main() {
	var apikey string = os.Args[1]
	var taskname string = os.Args[2]

	TasksAPI := tasks.LoadAPI(apikey)
	token := tasks.Authorize(taskname, TasksAPI)
	tasks.GetTask(token, TasksAPI)

	questionURL := TasksAPI.Endpoint + TasksAPI.Task + token

	question := "\"What is neuroscience?\""

	var questionBody bytes.Buffer
	writer := multipart.NewWriter(&questionBody)
	writer.WriteField("question", question)
	writer.Close()

	res, err := http.Post(questionURL, writer.FormDataContentType(), &questionBody)
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("Response to the qestion request error with code: %d\n", res.StatusCode)
		os.Exit(1)
	}

	responseBytes, _ := io.ReadAll(res.Body)

	fmt.Printf("\nPrinting response...\n")

	var questionAnswer tasks.QuestionAnswer
	json.Unmarshal(responseBytes, &questionAnswer)

	fmt.Println(questionAnswer.Answer)

	systemMessage := tasks.CompletionMessage{
		Role: "system",
		Content: "I detect if the answer presented was indeed an answer to question given. I only give answer YES/NO.",
	}

	content := fmt.Sprintf("Question: %v\nAnswer: \"%v\"\n", question, questionAnswer.Answer)

	userMessage := tasks.CompletionMessage{
		Role: "user",
		Content: content,
	}

	resGuardrails := tasks.OpenAiCompletionRequestGuardrails("https://api.openai.com/v1/chat/completions", TasksAPI, questionAnswer, "gpt-4", systemMessage, userMessage)

	guardrailsAnswer := resGuardrails.Choices[0].Message.Content

	tasks.SendAnswer(token, TasksAPI, guardrailsAnswer)

}
