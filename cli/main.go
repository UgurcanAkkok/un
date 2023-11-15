package main

import (
	"fmt"

	apihandler "uakkok.dev/un/cli/api-handler"
	localhandler "uakkok.dev/un/cli/local-handler"
	"uakkok.dev/un/common/tasks"
)

func main() {
	var tasksResp tasks.Tasks
	var backend BackendHandler
	backendType := "api"

	switch backendType {
	case "local":
		backend = &localhandler.LocalHandler{LocalDBFile: "./un.db"}
	case "api":
		backend = &apihandler.ApiHandler{ApiEndpoint: "http://localhost:8080"}
	default:
		fmt.Println("Error getting the backend type")
	}

	var err error
	backend.Init()

	if err != nil {
		fmt.Println("Failed to initalize backend: ", backendType, err)
		return
	}

	tasksResp, err = backend.GetTasks()
	if err != nil {
		fmt.Println("Failed to get tasks:", err)
		return
	}
	fmt.Printf("Tasks: %v\n", tasksResp)
	return
}
