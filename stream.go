package main

import (
	"fmt"
	"io"
	"log"

	"github.com/cloudwego/eino/schema"
)

// Eino官方的示例代码在流式输出上我感觉是有问题的
// 它每次都会打印一整行，而且打印整个messaage的内容（message是一个结构体）
// 然后又带上时间戳，在终端上看完全不是流式输出
// 我改成了fmt.Print(message.Content)，不带时间戳（log），不换行，每次只打印内容

// 用于呈现类似于打字机的流式输出效果：接收参数得是模型的流式输出(streamReader)
func reportStream(sr *schema.StreamReader[*schema.Message]) {
	defer sr.Close()

	fmt.Print("assitant:")
	for {
		message, err := sr.Recv()
		if err == io.EOF { // 流式输出结束
			return
		}
		if err != nil {
			log.Fatalf("recv failed: %v", err)
		}
		fmt.Print(message.Content)
	}
}
