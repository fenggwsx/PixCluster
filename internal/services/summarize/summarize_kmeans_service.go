package summarize

import (
	"context"
	"encoding/json"
	"os"

	"github.com/fenggwsx/PixCluster/internal/types/kmeanstype"
	"github.com/fenggwsx/PixCluster/internal/types/urltype"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func SummarizeKMeansService(imageUrl urltype.DataURL, kmeansResult []kmeanstype.CentroidResult) string {
	data, err := json.MarshalIndent(kmeansResult, "", "\t")
	if err != nil {
		panic(err)
	}

	client := openai.NewClient(
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
		option.WithAPIKey(os.Getenv("DASHSCOPE_API_KEY")),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(os.Getenv("SYSTEM_PROMPT")),
			openai.UserMessage([]openai.ChatCompletionContentPartUnionParam{
				{OfText: &openai.ChatCompletionContentPartTextParam{Text: string(data)}},
				{OfImageURL: &openai.ChatCompletionContentPartImageParam{
					ImageURL: openai.ChatCompletionContentPartImageImageURLParam{
						URL: imageUrl.Url(),
					}}},
			}),
		},
		Model: os.Getenv("MODEL_NAME"),
	})
	if err != nil {
		panic(err)
	}

	return chatCompletion.Choices[0].Message.Content
}
