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

func Run() {
	conn := util.DbusConnect()
	defer conn.Close()

	// util.DbusMsg(conn, fmt.Sprintf("pwd: %s", util.GetPwd()))
	reader := bufio.NewReader(os.Stdin)
	// util.DbusMsg(conn, "reader is connected")
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
		util.DbusMsg(conn, string(payload))

		var msg models.Message
		json.Unmarshal(payload, &msg)

		switch msg.Cmd {
		case "version":
			response := models.Response{Version: settings.TridactylNativeMessengerEmulationVersion, Code: 0}
			util.SendResponse(response)
		case "getconfig":
			config, err := util.ReadConfig()
			if err != nil {
				util.DbusMsg(conn, fmt.Sprintf("getconfig error: %s", err))
			} else {
				response := models.Response{Content: config, Code: 0}
				util.SendResponse(response)
			}
		case "run":
			resp := models.Response{Content: msg.Command, Code: 0}
			util.SendResponse(resp)
		default:
			util.DbusMsg(conn, fmt.Sprintf("unknown command: %s", msg.Cmd))
		}

		// 4. Send a JSON response back (Required by Firefox)
		util.SendResponse(models.Response{Status: "received", Time: time.Now().Format(time.Kitchen)})
	}
}
