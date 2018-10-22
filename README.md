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
go run shapes.go
# test from anywhere on your local system
go test -v github.com/jwalkerdev/learning-go-class/002-structs-methods-interfaces/shape
```