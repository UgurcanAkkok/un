package main

import (
	"fmt"

	apihandler "uakkok.dev/un/cli/api-handler"
)

func main() {
	api := apihandler.ApiHandler{
		ApiEndpoint: "http://localhost:8080",
	}
	tasksResp, err := api.GetTasks()
	if err != nil {
		fmt.Println("Got error while getting tasks using api:", err)
		return
	}
	fmt.Printf("First task in the list: %v", tasksResp.Items[0].Message)
	return
}
