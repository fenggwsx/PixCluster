package router

import (
	"context"
	"net/http"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	"github.com/fenggwsx/PixCluster/internal/pkg/resp"
	"github.com/fenggwsx/PixCluster/pkg/fcrouter"
)

func NewRouter() *fcrouter.Router {
	return fcrouter.NewRouter(fcrouter.ErrorHandlers{
		InternalServerError: func(ctx context.Context, event *events.HTTPTriggerEvent, err any) *events.HTTPTriggerResponse {
			fctx, _ := fccontext.FromContext(ctx)
			fctx.GetLogger().Error(err)

			return resp.NewResponse(resp.ResponseBody{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "Something went wrong in the server",
			}, nil)
		},
		MethodNotAllowed: func(ctx context.Context, event *events.HTTPTriggerEvent, allowedMethods []string) *events.HTTPTriggerResponse {
			return resp.NewResponse(resp.ResponseBody{
				Success: false,
				Code:    http.StatusMethodNotAllowed,
				Message: "The request method is not allowed in this route",
			}, map[string]string{"Allow": strings.Join(allowedMethods, ", ")})
		},
		NotFound: func(ctx context.Context, event *events.HTTPTriggerEvent) *events.HTTPTriggerResponse {
			return resp.NewResponse(resp.ResponseBody{
				Success: false,
				Code:    http.StatusNotFound,
				Message: "No route matches the request path",
			}, nil)
		},
		NotHTTPTrigger: func(ctx context.Context, event *events.HTTPTriggerEvent) *events.HTTPTriggerResponse {
			return resp.NewResponse(resp.ResponseBody{
				Success: false,
				Code:    http.StatusTeapot,
				Message: "The request did not come from an HTTP Trigger",
			}, nil)
		},
	})
}
