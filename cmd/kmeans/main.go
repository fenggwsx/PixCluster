package main

import (
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/fenggwsx/PixCluster/internal/handlers/kmeans"
	"github.com/fenggwsx/PixCluster/internal/pkg/router"
	"github.com/fenggwsx/PixCluster/pkg/fcrouter"
)

func main() {
	r := router.NewRouter()
	r.AddRoute("/api/kmeans", map[string]fcrouter.RouteHandler{"POST": kmeans.KMeansHandler})

	fc.Start(r.GetHandler())
}
