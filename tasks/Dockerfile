FROM golang

WORKDIR /go/src/github.com/iarriola/microservice-todo-gokit/tasks

RUN mkdir -p /go/src/github.com/iarriola/todo-microservice-gokit

ADD . /go/src/github.com/iarriola/todo-microservice-gokit

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/iarriola/todo-microservice-gokit/tasks/cmd  -watch github.com/iarriola/todo-microservice-gokit/tasks
