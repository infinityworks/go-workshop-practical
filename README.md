# Go Workshop

## Part 1 - Initialize the module

### Task
Create your Go module.

### Hints
- The [`go mod`](https://pkg.go.dev/cmd/go#hdr-Module_maintenance) command can help with initializing a Go module.
- Use `go help <command>` for more information about a command.

### Solution
[Part 1](/part1)

## Part 2 - Create an executable

### Task
Create and run an executable that prints "Hello, world!" to the terminal.

### Hints
- Executables are commonly stored in the cmd folder. ("./cmd/foo/main.go")
- Executable packages must be named `main` and have a `main` function.
- fmt can help print things to the terminal output.
- The [`go run`](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program) command can execute a package.

### Solution
[Part 2](/part2)

## Part 3 - Create a http server

### Task
Amend your executable to listen for incoming http requests, and route those requests to a handler which writes "Hello, world!" to the response.

### Hints
- The [`http.ServeMux`](https://pkg.go.dev/net/http#ServeMux) type can be used to route incoming requests to a handler.
- Your handler will need to implement the [`http.Handler`](https://pkg.go.dev/net/http#Handler) interface.
- [`http.ListenAndServe`](https://pkg.go.dev/net/http#ListenAndServe) can be used to listen for incoming http requests and direct them to your router.

### Solution
[Part 3](/part3)

## Part 4 - Basic request and response

### Task
Read in a name from the request body and use this to customise the response to "Hello, <name>".

### Hints
- The request and response types can be defined as structs.
- The [`json`](https://pkg.go.dev/encoding/json) package will be useful for dealing with request and response bodies.
- By default, the response status will be set to 200, but this can be overwritten using the `WriteHeader` func on [`ResponseWriter`](https://pkg.go.dev/net/http@go1.20.6#ResponseWriter) e.g. in the event of an error.

### Solution
[Part 4](/part4)

## Part 5 - Stub our real endpoint

### Task
Change the request & response to accept a transcript and return a summary, and return a placeholder summary for now.

### Solution
[Part 5](/part5)

## Part 6 - Install go-openai client (and create your free account)

### Task
Add the `go-openai` client as a dependency.

### Hints
- The [`go get`](https://pkg.go.dev/cmd/go#hdr-Add_dependencies_to_current_module_and_install_them) command is used to add dependencies.
- The package information can be found at https://pkg.go.dev/github.com/sashabaranov/go-openai

### Solution
[Part 6](/part6)

## Part 7 - Create a wrapper for the openai client

### Task
Create a new struct to wrap the openai client.  Add a `summarise` func that takes in a transcript, constructs and sends a suitable openai chat request, and returns the response.  Call this from your handler.

### Hints
- Given a type `Thing` it is convention to call the 'constructor' func `NewThing`.
- It's common in Go to have multiple type definitions in one file, if they are used together.
- You'll need to pass in a token for the openai service - this can be stored as an environment variable and retrieved using [`os.Getenv`](https://pkg.go.dev/os@go1.20.6#Getenv).

### Solution
[Part 7](/part7)

## Part 8 - Concurrently request summaries

### Task
Amend your code to accept multiple transcripts and return multiple summaries.

### Hints
- [`errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup) can be used to set up and wait for concurrent requests.

### Solution
[Part 8](/part8)
