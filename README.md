# learning-go-class
Repo used while learning golang in the ATX Learning Go Class

## Week 1 Homework
* Hello world with tests - based on https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

Test / Run
```bash
# Dev Cycle
cd 001-hello
go run hello.go
go test -v
# test from anywhere on your local system
go test -v github.com/jwalkerdev/learning-go-class/001-hello
```

## Week 2 Homework

* Inside a custom pkg walk through the following:
<https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces>
* In your top level folder create cmd directory with a main that calculates the area of a rectangle give two integers as parameters.
* Extra Credit: Use <https://github.com/spf13/cobra> to create sub command for each shape.

Test / Run
```bash
# Dev Cycle
cd 002-structs-methods-interfaces/shape
go test -v
# test from anywhere on your local system
go test -v github.com/jwalkerdev/learning-go-class/002-structs-methods-interfaces/shape
# Run main from cmd directory
cd 002-structs-methods-interfaces
go test ./cmd
go build ./...
go build ./cmd/main.go
./main
```

## Week 3 Homework - Tools

* Homework:
  * Write a palyndrome detection function in go, test it, and run all 3 tools.
	Check in your html coverage report along with your code.
    * Go lint (go get -u golang.org/x/lint/golint)
    * Go vet (built in) analysis tool
    * Go test coverage profile and the go tool cover HTML report generation.

```bash
# Dev
cd 003-tools/mypkg
go test -v
# Build
cd 003-tools
go build ./...
go run main.go
# golint
cd 003-tools
golint ./mypkg
# go vet
go vet ./mypkg
# go test coverage
go test -cover ./mypkg
go test -cover ./mypkg -coverprofile=./mypkg/mypkg-coverage.out
go tool cover -html=./mypkg/mypkg-coverage.out -o mypkg/coverage.html
```