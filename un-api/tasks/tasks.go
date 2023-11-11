package tasks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Tasks struct {
	Items []Task `json:"items"`
}

func GetTasks(c echo.Context) error {
	c.Logger().Debug("Getting tasks..")
	tasks := Tasks{
		Items: []Task{
			{ID: 1, Status: TaskStatusOpen,       Message: "create an awesome task manager app"},
			{ID: 2, Status: TaskStatusInProgress, Message: "create a beautiful task manager app"},
			{ID: 3, Status: TaskStatusCompleted,  Message: "create the best task manager app"},
			{ID: 4, Status: TaskStatusOpen,       Message: "create an intuitive task manager app"},
			{ID: 5, Status: TaskStatusInProgress, Message: "create a flexible task manager app"},
			{ID: 6, Status: TaskStatusOpen,       Message: "create a powerful task manager app"},
		},
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostTasks(c echo.Context) error {
	c.Logger().Debug("Creating tasks..")
	tasks := &Tasks{}
	if err := c.Bind(tasks); err != nil {
		c.Logger().Warn("Cant understand the post data for PostTasks")
		return c.String(http.StatusBadRequest, "Cant understand post data")
	}
	return c.JSON(http.StatusCreated, tasks)
}
