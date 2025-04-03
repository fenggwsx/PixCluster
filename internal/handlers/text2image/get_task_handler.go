package text2image

import (
	"context"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/fenggwsx/PixCluster/internal/pkg/resp"
	"github.com/fenggwsx/PixCluster/internal/services/text2image"
)

type GetTaskResponseBody struct {
	Status string  `json:"status"`
	Prefix *string `json:"prefix"`
	Data   *string `json:"data"`
}

func GetTaskHandler(ctx context.Context, event *events.HTTPTriggerEvent, params map[string]string) *events.HTTPTriggerResponse {
	taskId := params["taskId"]
	taskOutput := text2image.GetTaskService(taskId).Output
	resBody := GetTaskResponseBody{Status: taskOutput.TaskStatus}
	if taskOutput.Results != nil && len(*taskOutput.Results) > 0 {
		url := (*taskOutput.Results)[0].Url
		dataUrl := text2image.GetImageService(url)
		resBody.Prefix = &dataUrl.Prefix
		resBody.Data = &dataUrl.Data
	}

	return resp.NewResponse(resp.ResponseBody{
		Success: true,
		Code:    http.StatusOK,
		Message: "OK",
		Data:    &resBody,
	}, nil)
}
