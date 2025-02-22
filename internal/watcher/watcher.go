package watcher

import (
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
	go w.watchLoop()
	return w.events
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
	// Check if enough time has passed since the last event
	if time.Since(w.lastEvent) < time.Duration(w.config.DelayMS)*time.Millisecond {
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

	return true
}

func (w *Watcher) handleEvent() {
	w.lastEvent = time.Now()
	w.events <- true
}
