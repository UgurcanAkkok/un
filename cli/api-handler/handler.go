package apihandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	un "uakkok.dev/un/common"
)

type ApiHandler struct {
	ApiEndpoint string
}

func (a *ApiHandler) Init() error {
  // Authentication..
  return nil
}

func (a *ApiHandler) GetTasks() (un.Tasks, error) {
	tasksResponse := &un.Tasks{}

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

func (a *ApiHandler) Close() error {
  // Necessary cleanup potentially..
  return nil
}
