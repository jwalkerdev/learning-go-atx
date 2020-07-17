# learning-go-class
Repo used while learning golang in the ATX Learning Go Class

## Setup Go

Install Go
...

Setup Go Environment:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Example .bash_profile
```bash
export GOPATH=~/code/go
export GOPATH=$GOPATH:~/Dropbox/edu/udemy/Mastering\ Go\ Programming\ -\ Packt/go-jrw
export GOPATH=$GOPATH:~/Dropbox/edu/udemy/Web\ Dev\ with\ Googles\ Go\ Programming\go-jrw
export PATH=$PATH:$GOPATH/bin
```

If using an independent project directory (outside of default go directory), create directories and updates path vars:
```bash
GO_DIR="~/code/somedir"
export GOPATH=$GO_DIR
export PATH=$PATH:$GOPATH/bin
mkdir -p $GOPATH/src $GOPATH/pkg $GOPATH/bin
```

### Go Modules (starting in 1.11, will become standard)

#### Short Version
For this project
```bash
go mod init github.com/jwalkerdev/learning-go-class2
```

#### Longer Version
Go 1.11 introduced Modules, enabling an alternative workflow. This new approach will gradually become the default mode, deprecating the use of GOPATH.

Modules aim to solve problems related to dependency management, version selection and reproducible builds; they also enable users to run Go code outside of GOPATH.

Using Modules is pretty straightforward. Select any directory outside GOPATH as the root of your project, and create a new module with the go mod init command.

A go.mod file will be generated, containing the module path, a Go version, and its dependency requirements, which are the other modules needed for a successful build.

If no <modulepath> is specified, go mod init will try to guess the module path from the directory structure, but it can also be overrided, by supplying an argument.

Run
```bash
mkdir my-project
cd my-project
go mod init <modulepath>
```

Sample go.mod file
```
module cmd

go 1.12

require (
        github.com/google/pprof v0.0.0-20190515194954-54271f7e092f
        golang.org/x/arch v0.0.0-20190815191158-8a70ba74b3a1
        golang.org/x/tools v0.0.0-20190611154301-25a4f137592f
)
```

go mod help
```bash
go help mod
go help mod init
```


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