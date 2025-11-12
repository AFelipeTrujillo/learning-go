# Hello World

## Description
The classic first program in any programming language. This example demonstrates the basic structure of a Go program.

## Key Concepts
- **Package declaration**: Every Go file must belong to a package
- **Import statements**: Importing the `fmt` package for formatted I/O
- **main function**: The entry point of executable programs
- **fmt.Println**: Function to print text with a newline
- **fmt.Printf**: Function for formatted printing

## Running the Code
```bash
go run main.go
```

## Building an Executable
```bash
go build
./01-hello-world  # On Unix/Linux/macOS
```

## Expected Output
```
Hello, World!
Hello, Go Developer! Welcome to Go programming.
```

## Notes
- Go programs must have a `package main` declaration for executables
- The `main()` function is automatically called when the program runs
- Go uses automatic semicolon insertion, so semicolons are usually not needed

