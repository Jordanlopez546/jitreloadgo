package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jordanlopez546/jitreloadgo/internal/builder"
	"github.com/Jordanlopez546/jitreloadgo/internal/config"
	"github.com/Jordanlopez546/jitreloadgo/internal/logger"
	"github.com/Jordanlopez546/jitreloadgo/internal/process"
	"github.com/Jordanlopez546/jitreloadgo/internal/watcher"
)

func main() {
	// Parse command line flags
	watchDir := flag.String("dir", ".", "Directory to watch")
	entryPoint := flag.String("entry", "main.go", "Entry point file")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	// Initialize config
	cfg := config.NewDefaultConfig()
	cfg.WatchDir = *watchDir
	cfg.EntryPoint = *entryPoint
	cfg.DebugMode = *debug

	// Initialize components
	logger := logger.New(cfg.DebugMode)
	builder := builder.New(cfg, logger)
	procManager := process.New(cfg, logger)

	watcher, err := watcher.New(cfg, logger)
	if err != nil {
		logger.Error("Failed to initialize watcher: %v", err)
		os.Exit(1)
	}

	// Handle rebuilds
	go func() {
		for range watcher.Watch() {
			logger.Info("Changes detected, rebuilding...")

			if err := builder.Build(); err != nil {
				logger.Error("Build failed: %v", err)
				continue
			}

			if err := procManager.StartProcess(); err != nil {
				logger.Error("Failed to start process: %v", err)
			}
		}
	}()

	// Initialize build
	logger.Info("Building...")
	if err := builder.Build(); err != nil {
		logger.Error("Initial build failed: %v", err)
		os.Exit(1)
	}

	if err := procManager.StartProcess(); err != nil {
		logger.Error("Failed to start process: %v", err)
		os.Exit(1)
	}

	// Handle interrupts
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logger.Info("Shutting down...")
	procManager.StopProcess()
}
