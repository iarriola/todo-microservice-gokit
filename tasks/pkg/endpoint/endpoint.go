package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/iarriola/todo-microservice-gokit/tasks/pkg/service"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Task string `json:"task"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.TasksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		rs, err := s.Create(ctx, req.Task)
		return CreateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, task string) (rs string, err error) {
	request := CreateRequest{Task: task}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).Rs, response.(CreateResponse).Err
}
