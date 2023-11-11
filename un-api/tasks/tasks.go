package tasks

import (
	"net/http"
	"github.com/labstack/echo/v4"
  un "uakkok.dev/un/common"
)

func GetTasks(c echo.Context) error {
	c.Logger().Debug("Getting tasks..")
	tasks := un.Tasks{
		Items: []un.Task{
			{ID: 1, Status: un.TaskStatusOpen,       Message: "create an awesome task manager app"},
			{ID: 2, Status: un.TaskStatusInProgress, Message: "create a beautiful task manager app"},
			{ID: 3, Status: un.TaskStatusCompleted,  Message: "create the best task manager app"},
			{ID: 4, Status: un.TaskStatusOpen,       Message: "create an intuitive task manager app"},
			{ID: 5, Status: un.TaskStatusInProgress, Message: "create a flexible task manager app"},
			{ID: 6, Status: un.TaskStatusOpen,       Message: "create a powerful task manager app"},
		},
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostTasks(c echo.Context) error {
	c.Logger().Debug("Creating tasks..")
	tasks := &un.Tasks{}
	if err := c.Bind(tasks); err != nil {
		c.Logger().Warn("Cant understand the post data for PostTasks")
		return c.String(http.StatusBadRequest, "Cant understand post data")
	}
	return c.JSON(http.StatusCreated, tasks)
}
