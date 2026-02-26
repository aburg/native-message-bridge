package util

import (
	"encoding/binary"
	"encoding/json"
	"os"

	"github.com/aburg/native-message-bridge/models"
)

func SendResponse(resp models.Response) {
	data, _ := json.Marshal(resp)
	// Write 4-byte length header
	binary.Write(os.Stdout, binary.LittleEndian, uint32(len(data)))
	// Write actual JSON
	os.Stdout.Write(data)
	os.Stdout.Sync()
}
