package imageutil

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/fenggwsx/PixCluster/internal/types/imagetype"
)

func DecodeBase64Image(base64Data string) ([]imagetype.Pixel, error) {
	byteData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(byteData))
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	pixels := make([]imagetype.Pixel, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels[y*width+x] = imagetype.Pixel{
				Red:   float64(r>>8) / 255.0,
				Green: float64(g>>8) / 255.0,
				Blue:  float64(b>>8) / 255.0,
			}
		}
	}

	return pixels, nil
}
