package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aburg/native-message-bridge/models"
)

func SendResponse(resp models.Response) {
	DbusMsg(fmt.Sprintf("response.code: %d, content: %s", resp.Code, resp.Content))

	data, _ := json.Marshal(resp)
	// Write 4-byte length header
	binary.Write(os.Stdout, binary.LittleEndian, uint32(len(data)))
	// Write actual JSON
	os.Stdout.Write(data)
	os.Stdout.Sync()
}

func SendFailureResponse(msg string) {
	SendResponse(models.Response{Code: 1, Content: msg})
}
