package kmeans

import (
	"errors"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"sync"

	"github.com/fenggwsx/PixCluster/internal/types/imagetype"
	"github.com/fenggwsx/PixCluster/internal/types/kmeanstype"
	"github.com/fenggwsx/PixCluster/internal/utils/colorutil"
)

type Config struct {
	NumCentroids  uint    `json:"num_centroids"`
	MaxIterations uint    `json:"max_iterations"`
	NumBlocks     uint    `json:"num_blocks"`
	StopThreshold float64 `json:"stop_threshold"`
}

type Centroid struct {
	imagetype.Pixel
	Count uint64 `json:"count"`
}

type ClusterSummary struct {
	Sums   [][3]float64
	Counts []uint64
}

func initCentroids(pixels []imagetype.Pixel, numCentroids uint) []Centroid {
	centroids := make([]Centroid, numCentroids)
	first := rand.Intn(len(pixels))
	centroids[0].Pixel = pixels[first]

	n := len(pixels)
	distances := make([]float64, n)
	for i := range distances {
		distances[i] = math.Inf(1)
	}

	for i := uint(1); i < numCentroids; i++ {
		total := 0.0
		for j := 0; j < n; j++ {
			dx := pixels[j].Red - centroids[i-1].Red
			dy := pixels[j].Green - centroids[i-1].Green
			dz := pixels[j].Blue - centroids[i-1].Blue
			dist := dx*dx + dy*dy + dz*dz
			if dist < distances[j] {
				distances[j] = dist
			}
			total += distances[j]
		}
		r := rand.Float64() * total
		for j := 0; j < n; j++ {
			r -= distances[j]
			if r <= 0.0 {
				centroids[i].Pixel = pixels[j]
				break
			}
		}
	}

	return centroids
}

func findClosest(pixel imagetype.Pixel, centroids []Centroid) int {
	minDist := math.MaxFloat64
	closest := 0
	for idx, centroid := range centroids {
		dx := pixel.Red - centroid.Red
		dy := pixel.Green - centroid.Green
		dz := pixel.Blue - centroid.Blue
		dist := dx*dx + dy*dy + dz*dz
		if dist < minDist {
			minDist = dist
			closest = idx
		}
	}

	return closest
}

func KMeansService(pixels []imagetype.Pixel, config Config) []kmeanstype.CentroidResult {
	numCentroids := config.NumCentroids
	if numCentroids < 2 {
		panic(errors.New("invalid centroids number"))
	}

	maxIterations := config.MaxIterations
	if maxIterations == 0 {
		maxIterations = 100
	}

	numBlocks := config.NumBlocks
	if numBlocks == 0 {
		numBlocks = uint(runtime.NumCPU())
	}

	stopThreshold := config.StopThreshold
	if stopThreshold == 0 {
		stopThreshold = 1e-5
	}

	numPixels := uint(len(pixels))
	centroids := initCentroids(pixels, numCentroids)
	blockSize := numPixels / numBlocks
	for iteration := uint(0); iteration < maxIterations; iteration++ {
		results := make(chan ClusterSummary, numBlocks)
		var wg sync.WaitGroup

		for i := uint(0); i < numBlocks; i++ {
			wg.Add(1)
			start := i * blockSize
			end := (i + 1) * blockSize
			if i == numBlocks-1 {
				end = numPixels
			}
			go func(start, end uint) {
				defer wg.Done()
				sums := make([][3]float64, numCentroids)
				counts := make([]uint64, numCentroids)
				for i := start; i < end; i++ {
					pixel := pixels[i]
					closest := findClosest(pixel, centroids)
					sums[closest][0] += pixel.Red
					sums[closest][1] += pixel.Green
					sums[closest][2] += pixel.Blue
					counts[closest]++
				}
				results <- ClusterSummary{Sums: sums, Counts: counts}
			}(start, end)
		}

		wg.Wait()
		close(results)

		totalSums := make([][3]float64, numCentroids)
		totalCounts := make([]uint64, numCentroids)
		for res := range results {
			for i := uint(0); i < numCentroids; i++ {
				totalSums[i][0] += res.Sums[i][0]
				totalSums[i][1] += res.Sums[i][1]
				totalSums[i][2] += res.Sums[i][2]
				totalCounts[i] += res.Counts[i]
			}
		}

		oldCentroids := make([]Centroid, numCentroids)
		copy(oldCentroids, centroids)
		for i := uint(0); i < numCentroids; i++ {
			if totalCounts[i] == 0 {
				idx := rand.Intn(len(pixels))
				centroids[i].Pixel = pixels[idx]
			} else {
				centroids[i].Red = totalSums[i][0] / float64(totalCounts[i])
				centroids[i].Green = totalSums[i][1] / float64(totalCounts[i])
				centroids[i].Blue = totalSums[i][2] / float64(totalCounts[i])
			}
			centroids[i].Count = totalCounts[i]
		}

		converged := true
		for i := uint(0); i < numCentroids; i++ {
			dx := centroids[i].Red - oldCentroids[i].Red
			dy := centroids[i].Green - oldCentroids[i].Green
			dz := centroids[i].Blue - oldCentroids[i].Blue
			dist := dx*dx + dy*dy + dz*dz
			if dist > stopThreshold {
				converged = false
				break
			}
		}
		if converged {
			break
		}
	}

	sort.Slice(centroids, func(i, j int) bool {
		return centroids[i].Count > centroids[j].Count
	})

	result := make([]kmeanstype.CentroidResult, numCentroids)
	for i := uint(0); i < numCentroids; i++ {
		result[i].Color = colorutil.NormalizedRGB2Hex(centroids[i].Red, centroids[i].Green, centroids[i].Blue)
		result[i].Count = centroids[i].Count
	}

	return result
}
