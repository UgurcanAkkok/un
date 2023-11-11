package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"uakkok.dev/un/api/tasks"
)

func getIsAlive(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getTasks(c echo.Context) error {
	c.Logger().Debug("Getting tasks..")
	tasks := tasks.Tasks{
		Items: []tasks.Task{
			{ID: 1, Status: tasks.TaskStatusOpen, Message: "create an awesome task manager app"},
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
