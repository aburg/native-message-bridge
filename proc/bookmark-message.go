package proc

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/aburg/native-message-bridge/models"
)

func ProcessBookmarkMessage(msg models.Message) models.Response {
	var bookmarkCommand models.BookmarkCommand
	err := json.Unmarshal([]byte(msg.Command), &bookmarkCommand)
	if err != nil {
		return models.Response{Code: 1, Content: fmt.Sprintf("json decode failed with error: %s", err.Error())}
	}

	if bookmarkCommand.SubCmd != "add" {
		return models.Response{Code: 1, Content: fmt.Sprintf("invalid subcmd: %s", bookmarkCommand.SubCmd)}
	}

	out, err := exec.Command("bookmark", bookmarkCommand.SubCmd, bookmarkCommand.Path, bookmarkCommand.URL).CombinedOutput()
	if err != nil {
		return models.Response{Code: 1, Content: fmt.Sprintf("bookmark error: %s, output: %s", err.Error(), out)}
	}

	return models.Response{Code: 0, Content: fmt.Sprintf("bookmark created: %s", bookmarkCommand.Path)}
}
