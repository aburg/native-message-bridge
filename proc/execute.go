package proc

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aburg/native-message-bridge/models"
	"github.com/aburg/native-message-bridge/settings"
	"github.com/aburg/native-message-bridge/util"
)

func Excecute() {
	// util.DbusMsg(conn, fmt.Sprintf("pwd: %s", util.GetPwd()))
	reader := bufio.NewReader(os.Stdin)
	util.DbusMsg("reader is connected")
	for {
		// 1. Read the 4-byte length header (Little Endian)
		var length uint32
		err := binary.Read(reader, binary.LittleEndian, &length)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("Error reading length")
		}

		// 2. Read the JSON payload
		payload := make([]byte, length)
		_, err = io.ReadFull(reader, payload)
		if err != nil {
			panic("error reading payload")
		}

		// 3. Log the raw message
		util.DbusMsg(string(payload))

		var msg models.Message
		json.Unmarshal(payload, &msg)

		switch msg.Cmd {
		case "version":
			response := models.Response{Version: settings.TridactylNativeMessengerEmulationVersion, Code: 0}
			util.SendResponse(response)
		case "getconfig":
			config, err := util.ReadConfig()
			if err != nil {
				util.DbusMsg(fmt.Sprintf("getconfig error: %s", err))
			} else {
				response := models.Response{Content: config, Code: 0}
				util.SendResponse(response)
			}
		case "run":
			var command models.Command
			json.Unmarshal([]byte(msg.Command), &command)
			switch command.Cmd {
			case "bookmark":
				util.SendResponse(ProcessBookmarkMessage(msg))
			case "jellify":
				util.SendResponse(ProcessJellifyMessage(msg))
			default:
				util.SendFailureResponse(fmt.Sprintf("unknown cmd: %s", command.Cmd))
			}
		default:
			util.SendFailureResponse(fmt.Sprintf("unknown command: %s", msg.Cmd))
		}

		// 4. Send a JSON response back (Required by Firefox)
		// TODO: check if this is really needed or just an ai hallucination again
		util.SendResponse(models.Response{Status: "received", Time: time.Now().Format(time.Kitchen)})
	}
}
