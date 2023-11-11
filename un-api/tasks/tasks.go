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
			{ID: 1, Status: TaskStatusOpen, Message: "create an awesome task manager app"},
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
	return c.JSON(http.StatusOK, tasks)
}
