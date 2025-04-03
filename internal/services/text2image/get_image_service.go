package text2image

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/fenggwsx/PixCluster/internal/types/urltype"
)

func GetImageService(url string) urltype.DataURL {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(errors.New("Failed to get image data"))
	}

	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		panic(errors.New("Content-Type is empty"))
	}

	return urltype.DataURL{
		Prefix: fmt.Sprintf("data:%s;base64",
			contentType),
		Data: base64.StdEncoding.EncodeToString(imageData),
	}
}
