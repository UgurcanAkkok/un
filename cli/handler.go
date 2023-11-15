package main

import (
	un "uakkok.dev/un/common"
)

type BackendHandler interface {
	Init() error
	GetTasks() (un.Tasks, error)
	Close() error
}
