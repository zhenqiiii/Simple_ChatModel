package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量:godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 空上下文
	ctx := context.Background()

	// 创建messages,但是起始内容是写死的
	log.Printf("===create messages===\n")
	messages := createMessagesFromTemplate()
	log.Printf("message: %+v\n\n", messages)

	// 创建llm
	log.Printf("===create llm===\n")
	chatModel := createArkChatModel(ctx)
	log.Printf("create llm success\n\n")

	// log.Printf("===llm generate===\n")
	// result := generate(ctx, chatModel, messages)
	// log.Printf("%+v\n\n", result)

	// 流式输出内容,没搞错的话模型只会执行一次
	log.Printf("===llm stream generate===\n")
	streamResult := stream(ctx, chatModel, messages)
	// 呈现流式输出效果
	reportStream(streamResult)

}
