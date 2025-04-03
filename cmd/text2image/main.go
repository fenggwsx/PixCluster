package main

import (
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/fenggwsx/PixCluster/internal/handlers/text2image"
	"github.com/fenggwsx/PixCluster/internal/pkg/router"
	"github.com/fenggwsx/PixCluster/pkg/fcrouter"
)

func main() {
	r := router.NewRouter()
	r.AddRoute("/api/text2image", map[string]fcrouter.RouteHandler{"POST": text2image.CreateTaskHandler})
	r.AddRoute("/api/text2image/:taskId", map[string]fcrouter.RouteHandler{"GET": text2image.GetTaskHandler})

	fc.Start(r.GetHandler())
}
