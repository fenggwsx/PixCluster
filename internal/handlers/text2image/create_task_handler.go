package text2image

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/fenggwsx/PixCluster/internal/pkg/req"
	"github.com/fenggwsx/PixCluster/internal/pkg/resp"
	"github.com/fenggwsx/PixCluster/internal/services/text2image"
)

type CreateTaskRequestBody struct {
	Positive string `json:"positive"`
	Negative string `json:"negative"`
}

type CreateTaskResponseBody struct {
	TaskID string `json:"task_id"`
}

func CreateTaskHandler(ctx context.Context, event *events.HTTPTriggerEvent, params map[string]string) *events.HTTPTriggerResponse {
	bodyByte := req.MustDecodeBodyString(event)

	reqBody := CreateTaskRequestBody{}
	err := json.Unmarshal(bodyByte, &reqBody)
	if err != nil {
		return resp.NewResponse(resp.ResponseBody{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "The request body is not in a valid format",
		}, nil)
	}

	resBody := CreateTaskResponseBody{
		TaskID: text2image.CreateTaskService(reqBody.Positive, reqBody.Negative),
	}

	return resp.NewResponse(resp.ResponseBody{
		Success: true,
		Code:    http.StatusOK,
		Message: "OK",
		Data:    &resBody,
	}, nil)
}
