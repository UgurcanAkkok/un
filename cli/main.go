package main

import (
	"fmt"

	api "uakkok.dev/un/cli/api-handler"
	embedded "uakkok.dev/un/cli/embedded-handler"
	un "uakkok.dev/un/common"
)

func main() {
	var tasksResp un.Tasks
	var backend BackendHandler
	backendType := "embedded"

	switch backendType {
	case "embedded":
		backend = &embedded.EmbeddedHandler{LocalDBFile: "./un.db"}
	case "api":
		backend = &api.ApiHandler{ApiEndpoint: "http://localhost:8080"}
	default:
		fmt.Println("Error getting the backend type")
	}

	var err error
	backend.Init()
	defer backend.Close()

	if err != nil {
		fmt.Println("Failed to initalize backend: ", backendType, err)
		return
	}

	tasksData := un.Tasks{
		Items: []un.Task{
			{ID: 1, Status: un.TaskStatusOpen, Message: "create an awesome task manager app"},
			{ID: 2, Status: un.TaskStatusInProgress, Message: "create a beautiful task manager app"},
			{ID: 3, Status: un.TaskStatusCompleted, Message: "create the best task manager app"},
			{ID: 4, Status: un.TaskStatusOpen, Message: "create an intuitive task manager app"},
			{ID: 5, Status: un.TaskStatusInProgress, Message: "create a flexible task manager app"},
			{ID: 6, Status: un.TaskStatusOpen, Message: "create a powerful task manager app"},
		},
	}

	if err = backend.PostTasks(tasksData); err != nil {
		fmt.Println("Failed to post tasks:", tasksData)
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
