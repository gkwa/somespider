package somespider

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

func GenPath(configRelPath string) (string, error) {
	configFilePath, err := xdg.ConfigFile(configRelPath)
	if err != nil {
		return "", err
	}

	dirPerm := os.FileMode(0o700)

	d := filepath.Dir(configFilePath)

	if err := os.MkdirAll(d, dirPerm); err != nil {
		slog.Error("cache", "mkdir", "error", err)
		return "", err
	}

	return configFilePath, nil
}
