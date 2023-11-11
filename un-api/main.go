package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"uakkok.dev/un/api/tasks"
)

func getIsAlive(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/isAlive", getIsAlive)
	e.GET("/tasks", tasks.GetTasks)
	e.POST("/tasks", tasks.PostTasks)
	e.Logger.Fatal(e.Start(":8080"))
}
