package util

import (
	"os"
	"path/filepath"

	"github.com/aburg/native-message-bridge/settings"
)

func getConfigPath(segments ...string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("could not get homedir")
	}
	return filepath.Join(homeDir, settings.HomeRelativeConfigPath, filepath.Join(segments...))
}

func GetRcPath() string {
	return getConfigPath(settings.RcFile)
}
