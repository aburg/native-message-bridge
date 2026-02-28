package proc

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aburg/native-message-bridge/models"
	"github.com/aburg/native-message-bridge/util"
)

func ProcessReadMessage(payload []byte) models.Response {
	var readMessage models.ReadMessage
	err := json.Unmarshal(payload, &readMessage)
	if err != nil {
		return models.Response{Code: 1, Content: fmt.Sprintf("json decode failed with error: %s", err.Error())}
	}

	if !util.IsSafePath(util.GetConfigPath(), readMessage.File) {
		return models.Response{Code: 1, Content: fmt.Sprintf("not allowed to read file: %s", readMessage.File)}
	}

	content, err := os.ReadFile(readMessage.File)
	if err != nil {
		return models.Response{Code: 1, Content: fmt.Sprintf("file read failed with error: %s", err.Error())}
	}

	return models.Response{Code: 0, Content: string(content)}
}
