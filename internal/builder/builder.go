package builder

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Jordanlopez546/jitreloadgo/internal/config"
	"github.com/Jordanlopez546/jitreloadgo/internal/logger"
	"github.com/Jordanlopez546/jitreloadgo/pkg/utils"
)

type Builder struct {
	config *config.Config
	logger *logger.Logger
}

func New(config *config.Config, logger *logger.Logger) *Builder {
	return &Builder{
		config: config,
		logger: logger,
	}
}

func (b *Builder) Build() error {
	buildArgs := []string{"build"}
	buildArgs = append(buildArgs, b.config.BuildFlags...)
	buildArgs = append(buildArgs, "-o", utils.GetTempBinaryPath())
	buildArgs = append(buildArgs, b.config.EntryPoint)

	cmd := exec.Command("go", buildArgs...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	b.logger.Debug("Running build command: go %s", strings.Join(buildArgs, " "))

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build error: %s\n%s", err, stderr.String())
	}

	return nil
}
