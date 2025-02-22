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

## Building From Source

1. Clone the repository
git clone <https://github.com/Jordanlopez546/jitreloadgo.git>

2. Navigate to the project directory
cd jitreloadgo

3. Build and install

go build -o jitreloadgo cmd/jitreloadgo/main.go

sudo mv jitreloadgo $(go env GOPATH)/bin/

## Basic Usage

<!-- In your Go project directory, run: -->
jitreloadgo

## Advanced Options

<!-- Watch a specific directory -->
jitreloadgo -dir=./cmd

<!--  Specify a different entry point -->
jitreloadgo -entry=app.go

<!-- Enable debug logging -->
jitreloadgo -debug

<!-- Combine options -->
jitreloadgo -dir=./cmd -entry=main.go -debug

## Alternative Usage (without even installation)

<!-- You can also use go run directory -->
<!-- From your project directory -->
go run github.com/Jordanlopez546/jitreloadgo/cmd/jitreloadgo@latest
