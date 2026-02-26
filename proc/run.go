package proc

import (
	"fmt"
	"regexp"

	"github.com/aburg/native-message-bridge/models"
	"github.com/aburg/native-message-bridge/run"
	"github.com/aburg/native-message-bridge/util"
)

func Run(commandLine string) models.Response {
	re := regexp.MustCompile(`^([a-z-]+)\s+(.*)$`)

	matches := re.FindStringSubmatch(commandLine)

	if len(matches) > 0 {
		command := matches[1]
		args := matches[2]
		switch command {
		case "bookmark":
			return run.Bookmark(args)
		default:
			util.DbusMsg(fmt.Sprintf("command: %s, args:%s", command, args))
		}
	} else {
		util.DbusMsg("no match")
	}
	return models.Response{Code: 0}
}
