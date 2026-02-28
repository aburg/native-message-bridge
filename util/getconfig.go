package util

import (
	"os"
)

func ReadConfig() (string, error) {
	content, err := os.ReadFile(GetRcPath())
	if err != nil {
		return "", err
	}
	return string(content), nil
}
