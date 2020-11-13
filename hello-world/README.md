# Hello World

A simple and conventional hello world program in Go.

## Some Common Questions

#### Q: How do we run the code in our project?

Let's a look at Go CLI and look at the available commands.

| Command | Functionality |
|---------|---------------|
| `go build` | Compile go source code files |
| `go run` | Compiles and executes one or more files |
| `go fmt` | Formats all the code in each file of the current directory |
| `go install` | Compiles and "installs" a package |
| `go get` | Downloads the raw source code of someone's else package |
| `go test` | Runs any tests associated with the current project |

#### Q: What does a 'package main' mean?
A package is nothing but a directory inside your Go workspace containing one or more Go source files, or other Go packages. There are two types of packages
1. **Executable**: Generates a file that we can run
2. **Reusable**: Code used as 'helpers'. Holds reusable logic.

#### Q: What does 'import fmt' mean?
`fmt` is part of Golang's standard package library. Package `fmt` implements formatted I/O with functions analogous to C's printf and scanf. `import` simply import this package for our use.

#### Q: What is that 'func' thing?
`func` is short for function and is simply a function decleration term similar to for example `def` in Python or `function` in JavaScript.

#### Q: How is the main.go file organized?
General structure is like
1. Package decleration
2. Import other packages
3. Declare functions