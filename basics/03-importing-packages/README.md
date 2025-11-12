# Importing Packages

## Description
Learn how to import and use packages in Go. Understanding the import system is crucial for organizing code and using Go's extensive standard library.

## Key Concepts

### What is a Package?
- A **package** is a collection of Go source files in the same directory
- Every Go file belongs to exactly one package
- Packages help organize and reuse code
- Go has a rich **standard library** with built-in packages

### Import Syntax

#### 1. Single Import
```go
import "fmt"
```

#### 2. Multiple Imports (Recommended)
```go
import (
    "fmt"
    "math"
    "strings"
)
```

#### 3. Import with Alias
```go
import (
    f "fmt"              // Use f.Println()
    m "math"             // Use m.Sqrt()
)
```

#### 4. Dot Import (Not Recommended)
```go
import . "fmt"           // Use Println() directly
```
⚠️ **Avoid in production** - causes namespace pollution and confusion

#### 5. Blank Import (For Side Effects)
```go
import _ "image/png"     // Only runs init() function
```
Common use cases:
- Database drivers
- Image format decoders
- Plugin registration

## Common Standard Library Packages

### Input/Output
- **fmt** - Formatted I/O (print, scan)
- **io** - Basic I/O interfaces
- **bufio** - Buffered I/O

### Strings and Text
- **strings** - String manipulation
- **strconv** - String conversions
- **regexp** - Regular expressions
- **unicode** - Unicode utilities

### Math
- **math** - Mathematical functions
- **math/rand** - Random numbers (pseudo-random)
- **math/big** - Arbitrary-precision arithmetic

### Time
- **time** - Time and duration functions

### Collections
- **sort** - Sorting slices
- **container/list** - Doubly linked list
- **container/heap** - Heap operations

### File System
- **os** - Operating system functions
- **path/filepath** - File path manipulation
- **io/ioutil** - I/O utility functions

### Networking
- **net** - Network I/O
- **net/http** - HTTP client and server
- **net/url** - URL parsing

### Encoding
- **encoding/json** - JSON encoding/decoding
- **encoding/xml** - XML encoding/decoding
- **encoding/base64** - Base64 encoding

### Concurrency
- **sync** - Synchronization primitives
- **context** - Context for goroutines

### Error Handling
- **errors** - Error creation
- **fmt** - Error formatting

## Import Organization

### Good Practice
```go
import (
    // Standard library packages first
    "fmt"
    "strings"
    "time"
    
    // Third-party packages
    "github.com/gorilla/mux"
    
    // Your own packages
    "myproject/database"
    "myproject/models"
)
```

### Tools
- **goimports** - Automatically adds/removes imports
- **go fmt** - Formats code (including imports)

```bash
# Install goimports
go install golang.org/x/tools/cmd/goimports@latest

# Format code with auto-import
goimports -w main.go
```

## Package Naming Rules

1. **Lowercase** - Package names should be lowercase
2. **Short** - Keep names concise (`io` not `inputoutput`)
3. **No underscores** - Use `nethttp` not `net_http`
4. **Descriptive** - Name should describe what it does

## Running the Code
```bash
cd basics/03-importing-packages
go run main.go
```

## What the Code Demonstrates

1. **fmt** - Formatted output
2. **math** - Mathematical operations and constants
3. **strings** - String manipulation functions
4. **time** - Working with dates and times
5. **math/rand** - Generating random numbers
6. **Combining packages** - Using multiple packages together

## Exercises

1. Import the `os` package and print your username using `os.Getenv("USER")`
2. Use the `strconv` package to convert strings to numbers
3. Import `path/filepath` and work with file paths
4. Use `sort` package to sort a slice of numbers
5. Create a program that uses at least 5 different standard library packages

## Common Mistakes

❌ **Unused imports** - Go won't compile with unused imports
```go
import "fmt"  // If never used, compilation fails
```

❌ **Importing same package twice**
```go
import (
    "fmt"
    "fmt"  // Error!
)
```

❌ **Circular imports** - Package A imports B, B imports A
```go
// Not allowed in Go
```

## Notes

- Go automatically manages dependencies with `go.mod`
- Unused imports cause **compilation errors**
- Use `goimports` to automatically manage imports
- Standard library packages don't need `go get`
- Third-party packages need `go get` or `go mod download`
- Package names are typically the **last element** of import path
  - `import "net/http"` → use as `http.Get()`
  - `import "encoding/json"` → use as `json.Marshal()`

## Go Modules

### Initialize a module
```bash
go mod init github.com/username/project
```

### Add external dependencies
```bash
go get github.com/package/name
```

### Clean up dependencies
```bash
go mod tidy
```

## Related Topics
- Previous: fmt Package
- Next: Variables and Constants
- Related: Creating your own packages
- Related: Go modules and dependency management

