package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	api "uakkok.dev/un/cli/api-handler"
	backendhandler "uakkok.dev/un/cli/backend-handler"
	embedded "uakkok.dev/un/cli/embedded-handler"
	un "uakkok.dev/un/common"
)

// TODO: Better error and error message handling

type ContextKey string

var TaskArgs []string

func taskRun(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		fmt.Println("Error printing help on command task")
	}
}

func taskListRun(cmd *cobra.Command, args []string) {
	// Get the backend value from context and type assert to BackendHandler
	backend := cmd.Context().Value(ContextKey("backend")).(backendhandler.BackendHandler)
	if tasksResp, err := backend.GetTasks(); err != nil {
		fmt.Println("Failed to get tasks:", err)
		return
	} else {
		fmt.Printf("Tasks: %v\n", tasksResp)
	}
}

func taskAddRun(cmd *cobra.Command, args []string) {
	// Get the backend value from context and type assert to BackendHandler
	backend := cmd.Context().Value(ContextKey("backend")).(backendhandler.BackendHandler)
	var tasksData un.Tasks = un.Tasks{Items: []un.Task{}}
	if messageFlag, err := cmd.Flags().GetStringArray("message"); err != nil {
		fmt.Println("Error getting task messages from cli flag:", err)
	} else {
		for _, m := range messageFlag {
			if m == "" || m == "\n" {
				fmt.Println("Cant add empty task, skipping it")
			} else {
				task := un.Task{
					Status:  un.TaskStatusOpen,
					Message: m,
				}
				tasksData.Items = append(tasksData.Items, task)
			}
		}
		backend.PostTasks(tasksData)
	}
}

var taskCmd = &cobra.Command{
	Use:     "task",
	Run:     taskRun,
	Aliases: []string{"t"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var backend backendhandler.BackendHandler
		backendType := "embedded"

		switch backendType {
		case "embedded":
			backend = &embedded.EmbeddedHandler{LocalDBFile: "./un.db"}
		case "api":
			backend = &api.ApiHandler{ApiEndpoint: "http://localhost:8080"}
		default:
			fmt.Println("Error getting the backend type")
			return
		}
		if err := backend.Init(); err != nil {
			fmt.Println("Failed to initalize backend: ", err)
			return
		}
		ctx := context.WithValue(cmd.Context(), ContextKey("backend"), backend)
		cmd.SetContext(ctx)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		backendValue := ctx.Value(ContextKey("backend"))
		// panic on failure to type assert
		backend := backendValue.(backendhandler.BackendHandler)
		backend.Close()
	},
}

func init() {
	taskCmd.AddCommand(taskAddCmd)
	taskCmd.AddCommand(taskListCmd)
	taskAddCmd.Flags().StringArrayP("message", "m", []string{""}, "Task message")
	taskAddCmd.MarkFlagRequired("message")
}

var taskAddCmd = &cobra.Command{
	Use:     "add",
	Run:     taskAddRun,
	Aliases: []string{"a"},
}

var taskListCmd = &cobra.Command{
	Use:     "list",
	Run:     taskListRun,
	Aliases: []string{"l"},
}
