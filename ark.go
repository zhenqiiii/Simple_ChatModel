package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

// 连接大模型并返回模型实例
func createArkChatModel(ctx context.Context) model.ChatModel {
	// 获取环境变量
	apiKey := os.Getenv("MY_API_KEY")
	modelName := os.Getenv("MODEL_NAME")
	baseURL := os.Getenv("BASE_URL")
	// 获取模型实例
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: baseURL,
		Model:   modelName,
		APIKey:  apiKey,
	})
	if err != nil {
		log.Fatalf("create ark chat model failed, err=%v\n", err)
	}

	return chatModel

}
