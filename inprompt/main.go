package main

import (
	//"fmt"

	"fmt"
	"os"
	"strings"

	"github.com/m4cd/aidevs/tasks"
)

func main() {
	var apikey string = os.Args[1]
	var taskname string = os.Args[2]

	TasksAPI := tasks.LoadAPI(apikey)
	token := tasks.Authorize(taskname, TasksAPI)
	task := tasks.GetTask(token, TasksAPI)

	//Asking model to figure out the name from the question
	systemMessage := tasks.CompletionMessage{
		Role: "system",
		Content: "I return only the person's name from a sentence given.",
	}

	content := fmt.Sprintf("\"%v\"", task.Question)

	userMessage := tasks.CompletionMessage{
		Role: "user",
		Content: content,
	}

	completionResponse := tasks.OpenAiCompletionRequest("https://api.openai.com/v1/chat/completions", TasksAPI, task, "gpt-4", systemMessage, userMessage)

	personName := completionResponse.Choices[0].Message.Content
	fmt.Printf("Person's name:\n%v\n\n",personName)

	// filter the input
	var relevantInput []string
	for _, input := range task.Input {
		if strings.Contains(input, personName) {
			relevantInput = append(relevantInput, "\"" + input + "\"")
		}
	}

	fmt.Printf("Relevant inputs:\n%v\n\n", relevantInput)
	
	//Asking to answer the question based on the context provided
	systemMessage = tasks.CompletionMessage{
		Role: "system",
		Content: "[",
	}

	for _, i := range relevantInput {
		systemMessage.Content = systemMessage.Content + i
	}

	systemMessage.Content = systemMessage.Content + "]"

	//fmt.Println(systemMessage.Content)
	//content := fmt.Sprintf("\"%v\"", task.Question)

	userMessage = tasks.CompletionMessage{
		Role: "user",
		Content: task.Question,
	}

	answer := tasks.OpenAiCompletionRequest("https://api.openai.com/v1/chat/completions", TasksAPI, task, "gpt-4", systemMessage, userMessage)
	finalAnswer := answer.Choices[0].Message.Content

	tasks.SendAnswer(token,TasksAPI,finalAnswer)
	
	

	//content = fmt.Sprintf("\"%v\"", task.Question)



	// fmt.Println(task.Input[0])
	// fmt.Println(task.Question)

	// var answer string = task.Cookie
	// tasks.SendAnswer(token, TasksAPI, answer)

}
