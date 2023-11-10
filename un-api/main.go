package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getIsAlive(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type TaskStatus string

const (
	TaskStatusOpen       TaskStatus = "open"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID      int        `json:"id"`
	Status  TaskStatus `json:"status"`
	Message string     `json:"message"`
}

type Tasks struct {
  Items []Task
}

func getTasks(c echo.Context) error {
  c.Logger().Debug("Getting tasks..")
  tasks := Tasks {
    Items: []Task { 
      {ID: 1, Status: TaskStatusOpen, Message: "create an awesome task manager app" },
    },
  }
  return c.JSON(http.StatusOK, tasks)
}

func main() {
	e := echo.New()
	e.GET("/isAlive", getIsAlive)
	e.GET("/tasks", getTasks)
	e.Logger.Fatal(e.Start(":8080"))
}
