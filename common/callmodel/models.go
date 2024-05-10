package callmodel

import (
	"context"

	"github.com/tmc/langchaingo/llms/ollama"
)

// IModel 定义一个接口, 支持泛类
// 1. 获取model client链接
// 2. 调用call model方法
type IModel interface {
	GetModelClient(ctx context.Context, inter, name string) (*ollama.LLM, error)
	CallModel(ctx context.Context, prompt string, args ...interface{}) (string, error)
}

// 链接ollama model, 执行 ollama 查看当前下载模型
func ListLLMModel(ctx context.Context) error {
	//ollClient := GetOllamaModel()
	//_, err = ollClient.Client.Call(ctx)
	return nil
}
