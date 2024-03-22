package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadAPI(apikey string) TasksAPI {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error while loading .env file.")
	}
	TasksAPI := TasksAPI{
		Endpoint:     os.Getenv("ENDPOINT"),
		Token:        os.Getenv("TOKEN"),
		Answer:       os.Getenv("ANSWER"),
		Task:         os.Getenv("TASK"),
		Apikey:       apikey,
		OpenAiApikey: os.Getenv("OPENAI_API_KEY"),
	}
	return TasksAPI
}

func Authorize(taskname string, TasksAPI TasksAPI) string {

	requestURL := TasksAPI.Endpoint + TasksAPI.Token + taskname

	jsonbody := fmt.Sprintf(`{"apikey": "%s"}`, TasksAPI.Apikey)

	jsonbytes := []byte(jsonbody)
	bodyReader := bytes.NewReader(jsonbytes)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("Cannot create authorization request: %s\n", err)
		os.Exit(1)
	}

	httpClient := http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("Client error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("UnAuthorized with code: %d\n", res.StatusCode)
		os.Exit(1)
	}

	var AuthResponse AuthResponse
	bodyBytes, _ := io.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &AuthResponse)

	return AuthResponse.Token
}

func GetTask(token string, TasksAPI TasksAPI) TaskResponse {
	requestURL := TasksAPI.Endpoint + TasksAPI.Task + token

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("Cannot create task request: %s\n", err)
		os.Exit(1)
	}

	httpClient := http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("Cannot fetch the task: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("Task response error with code: %d\n", res.StatusCode)
		os.Exit(1)
	}

	var TaskResponse TaskResponse
	bodyBytes, _ := io.ReadAll(res.Body)

	fmt.Printf("\nPrinting task...\n")
	fmt.Printf("TOKEN:\n%v\n\n", token)
	fmt.Println(string(bodyBytes))
	fmt.Printf("\n")

	json.Unmarshal(bodyBytes, &TaskResponse)

	return TaskResponse
}

func SendAnswer(token string, TasksAPI TasksAPI, answer string) AnswerResponse {
	requestURL := TasksAPI.Endpoint + TasksAPI.Answer + token
	//fmt.Println(requestURL)

	jsonbody := fmt.Sprintf(`{"answer": "%s"}`, answer)

	fmt.Println(jsonbody)

	jsonbytes := []byte(jsonbody)
	bodyReader := bytes.NewReader(jsonbytes)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("Cannot create answer request: %s\n", err)
		os.Exit(1)
	}

	httpClient := http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("Client error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("Answer response error with code: %d\n", res.StatusCode)
		os.Exit(1)
	}

	var AnswerResponse AnswerResponse
	bodyBytes, _ := io.ReadAll(res.Body)

	fmt.Printf("\nPrinting answer response...\n")
	fmt.Println(string(bodyBytes))
	fmt.Printf("\n")

	json.Unmarshal(bodyBytes, &AnswerResponse)

	return AnswerResponse
}

func SendAnswerBoolTable(token string, TasksAPI TasksAPI, answer []bool) AnswerResponse {
	requestURL := TasksAPI.Endpoint + TasksAPI.Answer + token

	answerMarshaled, _ := json.Marshal(answer)
	jsonbody := fmt.Sprintf(`{"answer": %v}`, string(answerMarshaled))

	fmt.Println("Answer body...")
	fmt.Println(jsonbody)

	jsonbytes := []byte(jsonbody)
	bodyReader := bytes.NewReader(jsonbytes)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("Cannot create answer request: %s\n", err)
		os.Exit(1)
	}

	httpClient := http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("Client error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	/*
	if res.StatusCode != 200 {
		fmt.Printf("Answer response error with code: %d\n", res.StatusCode)
		os.Exit(1)
	}
*/
	var AnswerResponse AnswerResponse
	bodyBytes, _ := io.ReadAll(res.Body)

	fmt.Printf("\nPrinting answer response...\n")
	fmt.Println(string(bodyBytes))
	fmt.Printf("\n")

	json.Unmarshal(bodyBytes, &AnswerResponse)

	return AnswerResponse
}

func OpenAiModerationFlagged(endpoint string, TasksAPI TasksAPI, task TaskResponse, model string) []bool {
	requestURL := endpoint

	inputBytes, _ := json.Marshal(task.Input)
		
	jsonbody := fmt.Sprintf(`{"model": "%s", "input": %s}`, model, string(inputBytes))
	fmt.Println(jsonbody)
	
	jsonbytes := []byte(jsonbody)
	bodyReader := bytes.NewReader(jsonbytes)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("Cannot create answer request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+TasksAPI.OpenAiApikey)

	httpClient := http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("Client error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	
	if res.StatusCode != 200 {
		fmt.Printf("Answer response error with code: %d\n", res.StatusCode)
		os.Exit(1)
	}

	var ModerationResponse ModerationResponse
	bodyBytes, _ := io.ReadAll(res.Body)

	// fmt.Printf("\nPrinting answer response...\n")
	// fmt.Println(string(bodyBytes))
	// fmt.Printf("\n")

	json.Unmarshal(bodyBytes, &ModerationResponse)

	//fmt.Println(ModerationResponse.Results[0].Flagged)

	var result []bool
	
	for i := 0; i < len(ModerationResponse.Results); i++ {
		//fmt.Println(ModerationResponse.Results[i].Flagged)
		result = append(result, ModerationResponse.Results[i].Flagged)
	}
	
	return result

}
