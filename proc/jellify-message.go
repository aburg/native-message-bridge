package proc

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/aburg/native-message-bridge/models"
	"github.com/aburg/native-message-bridge/util"
)

func ProcessJellifyMessage(msg models.Message) models.Response {
	var cmd models.JellifyCommand
	err := json.Unmarshal([]byte(msg.Command), &cmd)
	if err != nil {
		return models.Response{Code: 1, Content: fmt.Sprintf("json decode failed with error: %s", err.Error())}
	}

	util.DbusMsg("calling bash now")
	out, err := exec.Command("/home/abu/.script-store/jellyfin/imdb-url-to-folder.sh", cmd.ImdbURL, "clipNotify").CombinedOutput()
	if err != nil {
		return models.Response{Code: 1, Content: fmt.Sprintf("jellify error: %s, output: %s", err.Error(), out)}
	}
	util.DbusMsg("bashed...")

	return models.Response{Code: 0}
}
