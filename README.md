# JITReloadGo

A fast, efficient, and developer-friendly hot-reload tool for Go applications. JITReloadGo watches your Go files and automatically rebuilds and restarts your application when changes are detected.

## Features

- 🚀 Fast reload times
- 🔄 Automatic build and restart
- 🎯 Intelligent file watching
- 🛠 Multiple entry point support
- 📝 Detailed logging
- ⚡ Efficient process management
- 🛑 Clean shutdown handling
- ⚙️ Highly configurable

## Installation

You can install JITReloadGo in several ways:

### 1. Direct Installation from GitHub

```bash
go install github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest
```

You're right. Let me update that section to include instructions for all operating systems:

### 2. Building From Source

```bash
# Clone the repository
git clone https://github.com/Jordanlopez546/jitreloadgo.git

# Navigate to the project directory
cd jitreloadgo

# For Mac/Linux:
go build -o jitreloadgo cmd/jitreloadgo/main.go
sudo mv jitreloadgo $(go env GOPATH)/bin/

# For Windows:
go build -o jitreloadgo.exe cmd/jitreloadgo/main.go
move jitreloadgo.exe %GOPATH%\bin
# Or if using PowerShell:
move jitreloadgo.exe $env:GOPATH\bin
```

Note for Windows users:

- You may need to run the Command Prompt or PowerShell as Administrator
- If `%GOPATH%` isn't set, the default is usually `%USERPROFILE%\go`
- Make sure your `%GOPATH%\bin` is in your PATH environment variable

For Windows PATH setup:

1. Open System Properties (Win + X > System)
2. Click on "Advanced system settings"
3. Click on "Environment Variables"
4. Under "User variables", find PATH and add `%GOPATH%\bin`

## Usage

### Basic Usage

```bash
# In your Go project directory
jitreloadgo

# View help and all available options
jitreloadgo -help
```

### Advanced Options

```bash
# Watch a specific directory
jitreloadgo -dir=./cmd

# Specify a different entry point
jitreloadgo -entry=app.go

# Enable debug logging
jitreloadgo -debug

# Combine options
jitreloadgo -dir=./cmd -entry=main.go -debug
```

### Alternative Usage (without installation)

You can use `go run` directly from GitHub without installing:

```bash
# Basic usage
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest

# All flags together
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest -debug -dir=./src -entry=main.go

# Different combinations:
# Debug mode only
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest -debug

# Specific directory with debug
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest -dir=./api -debug

# Different entry point with debug
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest -entry=server.go -debug

# All options with different values
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest -debug -dir=./services/auth -entry=auth.go
```

## Quick Start Example

Create a simple Go web server and see hot-reload in action:

```go
// main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Hot Reload!")
    })

    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}
```

Run with JITReloadGo:

```bash
jitreloadgo
# or
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest
```

Now edit your main.go file and save - the server will automatically rebuild and restart!

## Configuration Options

| Flag    | Description              | Default    | Example Usage |
|---------|-------------------------|------------|---------------|
| -dir    | Directory to watch      | .          | -dir=./src    |
| -entry  | Entry point file        | main.go    | -entry=app.go |
| -debug  | Enable debug logging    | false      | -debug        |

## Common Use Cases

### Web Development

```bash
jitreloadgo -entry=server.go
```

### API Development

```bash
jitreloadgo -dir=./api
```

### Microservices

```bash
jitreloadgo -dir=./services/auth -entry=auth.go
```

## Troubleshooting

### Common Issues

1. **Command not found**

   ```bash
   # Add Go bin to your PATH
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. **Permission denied**

   ```bash
   # Install with sudo
   sudo mv jitreloadgo $(go env GOPATH)/bin/
   ```

3. **Build errors**

   ```bash
   # Enable debug mode for more information
   jitreloadgo -debug
   ```

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Author

Jordan Nwabuike. (JeeTix)

- GitHub: [@Jordanlopez546](https://github.com/Jordanlopez546)

## Support

If you find any bugs or have feature requests, please create an issue in the GitHub repository.
