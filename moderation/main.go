package main

import (
	//"fmt"

	"fmt"
	"os"

	"github.com/m4cd/aidevs/tasks"
)

func main() {
	var apikey string = os.Args[1]
	var taskname string = os.Args[2]

	TasksAPI := tasks.LoadAPI(apikey)
	token := tasks.Authorize(taskname, TasksAPI)
	task := tasks.GetTask(token, TasksAPI)

	fmt.Printf("Printing inputs....\n\n")
	for i := 0; i < len(task.Input); i++ {
		fmt.Println("\"" + task.Input[i] + "\"")
	}
	fmt.Printf("\n")

	//fmt.Printf("Quering API....\n\n")
	answer := tasks.OpenAiModerationFlagged("https://api.openai.com/v1/moderations", TasksAPI, task, "text-moderation-latest")
	//fmt.Printf("\n")

	tasks.SendAnswerBoolTable(token, TasksAPI, answer)

	//asfdsafasfsdafsadfsadfsfafsasafdfsa

}
