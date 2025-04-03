package main

import (
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/fenggwsx/PixCluster/internal/handlers/summarize"
	"github.com/fenggwsx/PixCluster/internal/pkg/router"
	"github.com/fenggwsx/PixCluster/pkg/fcrouter"
)

func main() {
	r := router.NewRouter()
	r.AddRoute("/api/summarize", map[string]fcrouter.RouteHandler{"POST": summarize.SummarizeKMeansHandler})

	fc.Start(r.GetHandler())
}
