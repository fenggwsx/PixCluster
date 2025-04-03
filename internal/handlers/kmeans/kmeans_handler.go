package kmeans

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/fenggwsx/PixCluster/internal/pkg/req"
	"github.com/fenggwsx/PixCluster/internal/pkg/resp"
	"github.com/fenggwsx/PixCluster/internal/services/kmeans"
	"github.com/fenggwsx/PixCluster/internal/types/kmeanstype"
	"github.com/fenggwsx/PixCluster/internal/types/urltype"
	"github.com/fenggwsx/PixCluster/internal/utils/imageutil"
)

type KMeansRequestBody struct {
	K        uint            `json:"k"`
	ImageUrl urltype.DataURL `json:"image_url"`
}

type KMeansResponseBody struct {
	Result []kmeanstype.CentroidResult `json:"result"`
}

func KMeansHandler(ctx context.Context, event *events.HTTPTriggerEvent, params map[string]string) *events.HTTPTriggerResponse {
	bodyByte := req.MustDecodeBodyString(event)

	reqBody := KMeansRequestBody{}
	err := json.Unmarshal(bodyByte, &reqBody)
	if err != nil {
		return resp.NewResponse(resp.ResponseBody{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "The request body is not in a valid format",
		}, nil)
	}

	if reqBody.K < 2 {
		return resp.NewResponse(resp.ResponseBody{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "K should not be less than two",
		}, nil)
	}

	pixels, err := imageutil.DecodeBase64Image(reqBody.ImageUrl.Data)
	if err != nil {
		return resp.NewResponse(resp.ResponseBody{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "The base64 image is not in a valid format",
		}, nil)
	}

	resBody := KMeansResponseBody{
		Result: kmeans.KMeansService(pixels, kmeans.Config{
			NumCentroids: reqBody.K,
		}),
	}

	return resp.NewResponse(resp.ResponseBody{
		Success: true,
		Code:    http.StatusOK,
		Message: "OK",
		Data:    &resBody,
	}, nil)
}
