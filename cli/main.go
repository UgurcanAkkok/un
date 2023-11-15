package main

import (
	"fmt"

	apihandler "uakkok.dev/un/cli/api-handler"
	localhandler "uakkok.dev/un/cli/local-handler"
	"uakkok.dev/un/common/tasks"
)

func main() {
	var tasksResp tasks.Tasks
	api := apihandler.ApiHandler{
		ApiEndpoint: "http://localhost:8080",
	}
	tasksResp, err := api.GetTasks()
	if err != nil {
		fmt.Println("Got error while getting tasks using api:", err)
		return
	}
	fmt.Printf("Tasks: %v\n", tasksResp)
	local := localhandler.LocalHandler{
		LocalDBFile: "./un.db",
	}
  local.Init()
	tasksResp, err = local.GetTasks()
	if err != nil {
		fmt.Println("Got error while getting tasks using local:", err)
		return
	}
	fmt.Printf("Tasks: %v\n", tasksResp)
	return
}
