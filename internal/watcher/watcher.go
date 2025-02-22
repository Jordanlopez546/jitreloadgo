package watcher

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Jordanlopez546/jitreloadgo/internal/config"
	"github.com/Jordanlopez546/jitreloadgo/internal/logger"
	"github.com/Jordanlopez546/jitreloadgo/pkg/utils"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	config    *config.Config
	logger    *logger.Logger
	fsWatcher *fsnotify.Watcher
	events    chan bool
	lastEvent time.Time
}

func New(config *config.Config, logger *logger.Logger) (*Watcher, error) {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		config:    config,
		logger:    logger,
		fsWatcher: fsWatcher,
		events:    make(chan bool),
		lastEvent: time.Now(),
	}, nil
}

func (w *Watcher) Watch() chan bool {
	// Start watching the directory
	if err := w.addDirsToWatch(w.config.WatchDir); err != nil {
		w.logger.Error("Failed to start watching directories: %v", err)
		return w.events
	}

	go w.watchLoop()
	return w.events
}

func (w *Watcher) addDirsToWatch(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}

		// Only watch directories
		if !info.IsDir() {
			return nil
		}

		// Skip ignored directories
		for _, ignoredDir := range w.config.IgnoreDirs {
			if strings.Contains(path, ignoredDir) {
				return filepath.SkipDir
			}
		}

		// Add directory to watcher
		if err := w.fsWatcher.Add(path); err != nil {
			if !os.IsNotExist(err) {
				w.logger.Error("Failed to watch directory: %s - %v", path, err)
			}
			return nil // Continue even if there's an error
		}

		w.logger.Debug("Now watching directory: %s", path)
		return nil
	})
}

func (w *Watcher) watchLoop() {
	for {
		select {
		case event, ok := <-w.fsWatcher.Events:
			if !ok {
				return
			}

			// Important: Handle directory creation immediately
			if event.Op&fsnotify.Create == fsnotify.Create {
				// Check if it's a directory
				if info, err := os.Stat(event.Name); err == nil && info.IsDir() {
					w.logger.Debug("New directory created: %s", event.Name)
					// Add the new directory to the watcher
					if err := w.addDirsToWatch(event.Name); err != nil {
						w.logger.Error("Failed to watch new directory: %v", err)
					}
					// Trigger rebuild for new directory
					w.handleEvent()
				}
			}

			// For file creation/modification
			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
				// If it's a new file, check if we should watch it
				if info, err := os.Stat(event.Name); err == nil && !info.IsDir() {
					ext := filepath.Ext(event.Name)
					if utils.Contains(w.config.IncludeExt, ext) {
						w.handleEvent()
					}
				}
			}

			// For any other changes that should trigger rebuild
			if w.shouldTriggerRebuild(event) {
				w.handleEvent()
			}

		case err, ok := <-w.fsWatcher.Errors:
			if !ok {
				return
			}
			w.logger.Error("Watcher error: %v", err)
		}
	}
}

func (w *Watcher) shouldTriggerRebuild(event fsnotify.Event) bool {
	// Skip if not enough time has passed since last event
	if time.Since(w.lastEvent) < time.Duration(w.config.DelayMS)*time.Millisecond {
		return false
	}

	// Skip dot files and directories
	if strings.HasPrefix(filepath.Base(event.Name), ".") {
		return false
	}

	// Skip ignored directories
	for _, dir := range w.config.IgnoreDirs {
		if strings.Contains(event.Name, dir) {
			return false
		}
	}

	// For files, check extension
	if !strings.HasSuffix(event.Name, string(os.PathSeparator)) {
		ext := filepath.Ext(event.Name)
		if !utils.Contains(w.config.IncludeExt, ext) {
			return false
		}
	}

	w.logger.Debug("Change detected: %s", event.Name)
	return true
}

func (w *Watcher) handleEvent() {
	w.lastEvent = time.Now()
	w.events <- true
}

func (w *Watcher) Close() error {
	if w.fsWatcher != nil {
		return w.fsWatcher.Close()
	}
	return nil
}
