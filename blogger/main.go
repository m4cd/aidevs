package main

import (
	//"fmt"

	"os"

	"github.com/m4cd/aidevs/tasks"
)

func main() {
	var apikey string = os.Args[1]
	var taskname string = os.Args[2]

	TasksAPI := tasks.LoadAPI(apikey)
	token := tasks.Authorize(taskname, TasksAPI)
	tasks.GetTask(token, TasksAPI)

	//var answer string = task.Cookie
	//tasks.SendAnswer(token, TasksAPI, answer)

}
