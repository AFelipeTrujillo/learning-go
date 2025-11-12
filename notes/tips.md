# Go Tips and Best Practices

## General Tips

### 1. Go Formatting
- Always use `go fmt` to format your code
- Go has strict formatting rules that are enforced by tools
```bash
go fmt ./...
```

### 2. Naming Conventions
- **Exported** names start with uppercase: `PublicFunction()`
- **Unexported** names start with lowercase: `privateFunction()`
- Use **camelCase** for multi-word names
- Keep names concise but meaningful

### 3. Error Handling
```go
// Good: Check errors immediately
result, err := someFunction()
if err != nil {
    return err
}

// Avoid: Ignoring errors
result, _ := someFunction() // Bad practice
```

### 4. Use Go Modules
```bash
# Initialize a new module
go mod init github.com/username/projectname

# Add dependencies
go get package-name

# Clean up dependencies
go mod tidy
```

### 5. Defer for Cleanup
```go
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close() // Ensures file is closed when function returns
```

## Common Mistakes to Avoid

1. **Not checking errors**: Always handle errors returned by functions
2. **Ignoring goroutine leaks**: Make sure goroutines can exit properly
3. **Mutating loop variables**: Be careful with closures in loops
4. **Using pointers unnecessarily**: Go passes structs efficiently
5. **Not using context**: Use context for cancellation and timeouts

## Useful Commands

```bash
# Run code
go run main.go

# Build executable
go build

# Run tests
go test ./...

# Get dependencies
go get package

# Format code
go fmt ./...

# Lint code
go vet ./...

# View documentation
go doc package.Function

# Initialize module
go mod init module-name

# Download dependencies
go mod download

# Clean up unused dependencies
go mod tidy
```

## Code Organization

```
project/
├── cmd/           # Command-line applications
├── internal/      # Private application code
├── pkg/           # Public library code
├── api/           # API definitions
├── web/           # Web assets
├── configs/       # Configuration files
├── scripts/       # Build and installation scripts
├── test/          # Additional test data
└── docs/          # Documentation
```

## Performance Tips

1. **Use string builder** for concatenating many strings
2. **Preallocate slices** when size is known
3. **Use sync.Pool** for frequently allocated objects
4. **Profile your code** with pprof
5. **Benchmark critical sections** with testing.B

## Testing Tips

```go
// Table-driven tests are idiomatic
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

