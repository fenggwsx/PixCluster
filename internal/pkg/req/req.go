package req

import (
	"encoding/base64"

	"github.com/aliyun/fc-runtime-go-sdk/events"
)

func MustDecodeBodyString(event *events.HTTPTriggerEvent) []byte {
	if event.IsBase64Encoded != nil && *event.IsBase64Encoded {
		decodedByte, err := base64.StdEncoding.DecodeString(*event.Body)
		if err != nil {
			panic(err)
		}
		return decodedByte
	}
	return []byte(*event.Body)
}
