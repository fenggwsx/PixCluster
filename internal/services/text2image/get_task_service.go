package text2image

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GetTaskResponseBody struct {
	RequestID string `json:"request_id"`
	Output    struct {
		TaskID        string `json:"task_id"`
		TaskStatus    string `json:"task_status"`
		SubmitTime    string `json:"subimit_time"`
		ScheduledTime string `json:"scheduled_time"`
		EndTime       string `json:"end_time"`
		Results       *[]struct {
			OriginalPrompt string `json:"orig_prompt"`
			AcctualPrompt  string `json:"actual_prompt"`
			Url            string `json:"url"`
		} `json:"results"`
		TaskMetrics struct {
			Total     int `json:"TOTAL"`
			Succeeded int `json:"SUCCEEDED"`
			Failed    int `json:"FAILED"`
		}
	} `json:"output"`
	Usage struct {
		ImageCount int `json:"image_count"`
	} `json:"usage"`
}

func GetTaskService(taskId string) *GetTaskResponseBody {
	targetUrl := fmt.Sprintf("https://dashscope.aliyuncs.com/api/v1/tasks/%s", taskId)
	reqHeader := http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", os.Getenv("DASHSCOPE_API_KEY"))},
	}

	request, err := http.NewRequest("GET", targetUrl, nil)
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

	resBody := GetTaskResponseBody{}
	err = json.Unmarshal(resPayload, &resBody)
	if err != nil {
		panic(err)
	}

	return &resBody
}
