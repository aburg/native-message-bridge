package run

import (
	"fmt"
	"os/exec"

	"github.com/aburg/native-message-bridge/models"
	"github.com/aburg/native-message-bridge/util"
)

func Bookmark(args string) models.Response {
	out, err := exec.Command("bookmark", args).CombinedOutput()
	if err != nil {
		util.DbusMsg(fmt.Sprintf("bookmark error: %s", err.Error()))
		return models.Response{Code: 1, Content: "bookmark error"}
	} else {
		util.DbusMsg("bookmark done")
		return models.Response{Code: 0, Content: string(out)}
	}
}
