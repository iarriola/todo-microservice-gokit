package service

import (
	"context"
	"fmt"
)

// TasksService describes the service.
type TasksService interface {
	// Add your methods here
	Create(ctx context.Context, task string) (rs string, err error)
}

type basicTasksService struct{}

func (b *basicTasksService) Create(ctx context.Context, task string) (rs string, err error) {
	fmt.Println("Creating task")

	return rs, err
}

// NewBasicTasksService returns a naive, stateless implementation of TasksService.
func NewBasicTasksService() TasksService {
	return &basicTasksService{}
}

// New returns a TasksService with all of the expected middleware wired in.
func New(middleware []Middleware) TasksService {
	var svc TasksService = NewBasicTasksService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
