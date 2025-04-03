package summarize

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/fenggwsx/PixCluster/internal/pkg/req"
	"github.com/fenggwsx/PixCluster/internal/pkg/resp"
	"github.com/fenggwsx/PixCluster/internal/services/summarize"
	"github.com/fenggwsx/PixCluster/internal/types/kmeanstype"
	"github.com/fenggwsx/PixCluster/internal/types/urltype"
)

type SummarizeKMeansRequestBody struct {
	ImageUrl     urltype.DataURL             `json:"image_url"`
	KMeansResult []kmeanstype.CentroidResult `json:"kmeans_result"`
}

type SummarizeKMeansResponseBody struct {
	Summary string `json:"summary"`
}

func SummarizeKMeansHandler(ctx context.Context, event *events.HTTPTriggerEvent, params map[string]string) *events.HTTPTriggerResponse {
	bodyByte := req.MustDecodeBodyString(event)

	reqBody := SummarizeKMeansRequestBody{}
	err := json.Unmarshal(bodyByte, &reqBody)
	if err != nil {
		return resp.NewResponse(resp.ResponseBody{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "The request body is not in a valid format",
		}, nil)
	}

	resBody := SummarizeKMeansResponseBody{
		Summary: summarize.SummarizeKMeansService(reqBody.ImageUrl, reqBody.KMeansResult),
	}

	return resp.NewResponse(resp.ResponseBody{
		Success: true,
		Code:    http.StatusOK,
		Message: "OK",
		Data:    &resBody,
	}, nil)
}
