package text2image

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ImageSynthesisRequestBody struct {
	Model string `json:"model"`
	Input struct {
		Prompt         string  `json:"prompt"`
		NegativePrompt string `json:"negative_prompt,omitempty"`
	} `json:"input"`
	Parameters struct {
		Size         *string `json:"size,omitempty"`
		N            *int    `json:"n,omitempty"`
		Seed         *int    `json:"seed,omitempty"`
		PromptExtend *bool   `json:"prompt_extend,omitempty"`
		Watermark    *bool   `json:"watermark,omitempty"`
	} `json:"parameters"`
}

type ImageSynthesisSuccessResponseBody struct {
	Output struct {
		TaskStatus string `json:"task_status"`
		TaskID     string `json:"task_id"`
	} `json:"output"`
	RequestID string `json:"request_id"`
}

type ImageSynthesisFailResponseBody struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func CreateTaskService(positivePrompt string, negativePrompt string) string {
	targetUrl := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text2image/image-synthesis"
	pictureCount := 1
	pictureSize := os.Getenv("PICTURE_SIZE")
	reqHeader := http.Header{
		"Authorization":     {fmt.Sprintf("Bearer %s", os.Getenv("DASHSCOPE_API_KEY"))},
		"Content-Type":      {"application/json"},
		"X-DashScope-Async": {"enable"},
	}
	reqBody := ImageSynthesisRequestBody{Model: os.Getenv("MODEL_NAME")}
	reqBody.Input.Prompt = positivePrompt
	reqBody.Input.NegativePrompt = negativePrompt
	reqBody.Parameters.N = &pictureCount
	reqBody.Parameters.Size = &pictureSize
	reqPayload, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", targetUrl, bytes.NewReader(reqPayload))
	if err != nil {
		panic(err)
	}
	request.Header = reqHeader

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	resPayload, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if response.StatusCode == http.StatusOK {
		resBody := ImageSynthesisSuccessResponseBody{}
		err = json.Unmarshal(resPayload, &resBody)
		if err != nil {
			panic(err)
		}
		return resBody.Output.TaskID
	} else {
		resBody := ImageSynthesisFailResponseBody{}
		err = json.Unmarshal(resPayload, &resBody)
		if err != nil {
			panic(err)
		}
		panic(errors.New(resBody.Message))
	}
}
