# Learn Go with Tests - by github.com/quii

https://quii.gitbook.io/learn-go-with-tests
https://github.com/quii/learn-go-with-tests

# Install Go

Install options
* Get install from golang site. 
* OSX: use brew  
    `brew install go --cross-compile-common`
* etc (use package managers)

## Setup env vars

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
mkdir -p $GOPATH/src $GOPATH/pkg $GOPATH/bin
```

## Setup VSCode extension
`code --install-extension ms-vscode.go`

## Go Debugger
Delve is a good VS Code debugger. Install using go get:
```bash
go get -u github.com/derekparker/delve/cmd/dlv
```

## Improved linting
```sh
go get -u github.com/alecthomas/gometalinter
gometalinter --install
```

## Using godoc
`godoc -http=:6060`
navigate to http://localhost:6060/pkg/

### Example funcs
Example func name format: `"Example"+<func name>`
e.g. Example of func "Add" would be named "ExampleAdd"
Add example funcs to the *_test.go file. They are compiled during the build.
Example funcs show up in the godoc under the target func as an example of that func.

### Benchmarking
Benchmark func name format: `"Benchmark"+<func name>`.  
e.g. Benchmark for func "Repeat" would be named "BenchmarkRepeat"  

The code slightly different than a test.  
func signature:  `func Benchmark+<func name>(b *testing.B)`  

The testing.B gives you access to `b.N`. The var `b.N` is the number of times to repeat the test.
When the benchmark code is executed, it runs b.N times and measures how long it takes.  
The framework determines the number of times to run the code

Running the benchmark:
`go test -bench=.`

### Test Coverage
`go test -cover`

## Slices
https://blog.golang.org/go-slices-usage-and-internals

## Error linting (finding unchecked errors)
`go get -u github.com/kisielk/errcheck`  
Then, inside the directory with your code run `errcheck .`
