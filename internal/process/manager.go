package process

import (
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/Jordanlopez546/jitreloadgo/internal/config"
	"github.com/Jordanlopez546/jitreloadgo/internal/logger"
	"github.com/Jordanlopez546/jitreloadgo/pkg/utils"
)

type Manager struct {
	config *config.Config
	logger *logger.Logger
	cmd    *exec.Cmd
}

func New(config *config.Config, logger *logger.Logger) *Manager {
	return &Manager{
		config: config,
		logger: logger,
	}
}

func (m *Manager) StartProcess() error {
	if m.cmd != nil && m.cmd.Process != nil {
		m.StopProcess()
	}

	m.cmd = exec.Command(utils.GetTempBinaryPath())
	m.cmd.Stdout = os.Stdout
	m.cmd.Stderr = os.Stderr

	return m.cmd.Start()
}

func (m *Manager) StopProcess() {
	if m.cmd == nil || m.cmd.Process == nil {
		return
	}

	// Send interrupt signal first
	m.cmd.Process.Signal(syscall.SIGINT)

	// Wait a bit and then force kill if needed
	done := make(chan error)
	go func() {
		done <- m.cmd.Wait()
	}()

	select {
	case <-done:
		return
	case <-time.After(500 * time.Millisecond):
		m.cmd.Process.Kill()
	}
}
