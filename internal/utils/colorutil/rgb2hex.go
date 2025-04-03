package colorutil

import (
	"fmt"
	"math"

	"github.com/fenggwsx/PixCluster/internal/utils/mathutil"
)

func NormalizedRGB2Hex(red, green, blue float64) string {
	redInt := uint8(math.Round(mathutil.Clamp(red*255, 0, 255)))
	greenInt := uint8(math.Round(mathutil.Clamp(green*255, 0, 255)))
	blueInt := uint8(math.Round(mathutil.Clamp(blue*255, 0, 255)))

	return fmt.Sprintf("#%02x%02x%02x", redInt, greenInt, blueInt)
}
