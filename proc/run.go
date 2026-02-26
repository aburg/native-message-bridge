package proc

import (
	"fmt"
	"regexp"

	"github.com/aburg/native-message-bridge/models"
	"github.com/aburg/native-message-bridge/util"
)

func Run(commandLine string) models.Response {
	re := regexp.MustCompile(`^([a-z-]+)\s+(.*)$`)

	matches := re.FindStringSubmatch(commandLine)

	if len(matches) > 0 {
		util.DbusMsg("no match")
	} else {
		util.DbusMsg(fmt.Sprintf("command: %s, args:%s", matches[1], matches[2]))
	}
	return models.Response{Code: 0}
}
