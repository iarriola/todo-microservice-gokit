# todo-microservice-gokit
A simple microservice with Go and Go Kit. This project has been constructed making used of 
[GoKit CLI](https://github.com/kujtimiihoxha/kit).

# Project structure
```
.
├── README.md
├── docker-compose.yml
└── tasks (module)
    ├── Dockerfile
    ├── cmd
    ├── go.mod
    ├── go.sum
    ├── pkg
    └── vendor
```
# How to run
- Execute `go run tasks/cmd/main.go`
- Execute `curl -XPOST localhost:8081/create -d '{"task":"test"}'`
- Optionally you can use the included `Dockerfile` file, so you can execute `docker-compose up -d`

# References
1. [Go Modules Reference](https://golang.org/ref/mod)
   * [A Proposal for Package Versioning in Go](https://blog.golang.org/versioning-proposal)
   * [Using Go Modules](https://blog.golang.org/using-go-modules)
2. [Microservices with go-kit. Part 1 (video)](https://www.youtube.com/watch?v=1ScP5DyS1_g&ab_channel=packagemain)