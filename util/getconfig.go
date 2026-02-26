package util

import (
	"os"
	"path/filepath"

	"github.com/aburg/native-message-bridge/settings"
)

func ReadConfig() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, settings.HomeRelativeConfigPath)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
