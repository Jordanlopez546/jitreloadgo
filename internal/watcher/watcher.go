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
			return err
		}

		// Skip if it's not a directory
		if !info.IsDir() {
			return nil
		}

		// Skip ignored directories
		for _, ignoredDir := range w.config.IgnoreDirs {
			if filepath.Base(path) == ignoredDir {
				return filepath.SkipDir
			}
		}

		// Add directory to watcher
		if err := w.fsWatcher.Add(path); err != nil {
			return err
		}

		w.logger.Debug("Watching directory: %s", path)
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
	// Check if enough time has passed since last event
	if time.Since(w.lastEvent) < time.Duration(w.config.DelayMS)*time.Millisecond {
		return false
	}

	// Only trigger on write and rename events
	if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) == 0 {
		return false
	}

	// Check file extension
	ext := filepath.Ext(event.Name)
	if !utils.Contains(w.config.IncludeExt, ext) {
		return false
	}

	// Check ignored directories
	for _, dir := range w.config.IgnoreDirs {
		if strings.Contains(event.Name, dir) {
			return false
		}
	}

	w.logger.Debug("File changed: %s", event.Name)
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
