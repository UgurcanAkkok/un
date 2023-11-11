package main

import (
	"fmt"
	"io"
	"net/http"
  "encoding/json"
  un "uakkok.dev/un/common"
)


func main(){
  unApiHostName := "http://localhost:8080/"
  tasksEndpoint := "tasks"

  resp, err := http.Get(unApiHostName + tasksEndpoint)
  if err != nil {
    fmt.Println("Unable to get tasks:", err)
    return
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    fmt.Println("Unable to get tasks with http status code:", resp.StatusCode)
    return
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Unable to read response body:", err)
    return
  }

  var respTasks un.Tasks
  if err := json.Unmarshal(body, &respTasks); err != nil {
    fmt.Println("Unable to read body into json:", err)
  }


  fmt.Printf("Got the response: %v", respTasks )
  return
}
