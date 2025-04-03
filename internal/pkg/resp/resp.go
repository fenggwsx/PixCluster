package resp

import (
	"encoding/json"

	"github.com/aliyun/fc-runtime-go-sdk/events"
)

type ResponseBody struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewResponse(body ResponseBody, headers map[string]string) *events.HTTPTriggerResponse {
	bytes, err := json.Marshal(&body)
	if err != nil {
		panic(err)
	}
	return &events.HTTPTriggerResponse{
		StatusCode: body.Code,
		Headers:    headers,
		Body:       string(bytes),
	}
}
