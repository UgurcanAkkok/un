package main

import (
	tasks "uakkok.dev/un/common/tasks"
)

type BackendHandler interface {
	Init() error
	GetTasks() (tasks.Tasks, error)
	Close() error
}
