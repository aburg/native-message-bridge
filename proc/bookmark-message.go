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

	path := bookmarkCommand.Title
	if bookmarkCommand.Folder != "" {
		path = fmt.Sprintf("%s/%s", bookmarkCommand.Folder, bookmarkCommand.Title)
	}

	out, err := exec.Command("bash", "/home/abu/.tx-store/bookmark/record-with-hypr-popup.sh", bookmarkCommand.URL, path).CombinedOutput()
	if err != nil {
		return models.Response{Code: 1, Content: fmt.Sprintf("bookmark error: %s, output: %s", err.Error(), out)}
	}

	return models.Response{Code: 0, Content: "a bookmark might have been created. i am blind for now. or forever?"}
}
