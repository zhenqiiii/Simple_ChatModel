package main

import (
	"context"
	"log"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// 给大模型的消息的模板
func createTemplate() prompt.ChatTemplate {
	// 创建模板，这里用GoTemplate格式
	template := prompt.FromMessages(schema.GoTemplate,
		// 系统消息模板
		schema.SystemMessage("你是一个{{.role}}。你需要用{{.style}}的语气回答问题。你的目标是帮助新手程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),

		// 插入对话历史（新对话没有）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题：{{.question}}"),
	)

	return template

}

// 使用模板生成消息
func createMessagesFromTemplate() []*schema.Message {
	// 先创建一个模板实例
	template := createTemplate()
	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "从业20余年的老程序员，同时也是一个爱说脏话的贴吧老哥，在一个论坛中担任程序员鼓励师的职位",
		"style":    "专业，一针见血并且直白",
		"question": "我的代码一直报错，感觉好沮丧，该怎么办？",
		// 对话历史（这个例子里模拟两轮对话历史）
		"chat_history": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("嘿！我是你的程序员鼓励师！记住，每个优秀的程序员都是从 Debug 中成长起来的。有什么我可以帮你的吗？", nil),
			schema.UserMessage("我觉得自己写的代码太烂了"),
			schema.AssistantMessage("每个程序员都经历过这个阶段！重要的是你在不断学习和进步。让我们一起看看代码，我相信通过重构和优化，它会变得更好。记住，Rome wasn't built in a day，代码质量是通过持续改进来提升的。", nil),
		},
	})
	if err != nil {
		log.Fatalf("failed to format template: %v/n", err)
	}

	return messages
}
