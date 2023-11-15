package apihandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"uakkok.dev/un/common/tasks"
)

type ApiHandler struct {
	ApiEndpoint string
}

func (a *ApiHandler) Init() error {
  // Authentication..
  return nil
}

func (a *ApiHandler) GetTasks() (tasks.Tasks, error) {
	tasksResponse := &tasks.Tasks{}

	resp, err := http.Get(a.ApiEndpoint + "/tasks")
	if err != nil {
		return *tasksResponse, fmt.Errorf("Unable to get tasks: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *tasksResponse, fmt.Errorf("Unable to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return *tasksResponse,
			fmt.Errorf("Get tasks not successful with status code: %d, body: %v", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, &tasksResponse); err != nil {
		return *tasksResponse,
			fmt.Errorf("Failed to unmarshal body into json, err: %w, body: %v", err, body)
	}

	return *tasksResponse, nil
}
