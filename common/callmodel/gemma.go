package callmodel

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/zeromicro/go-zero/core/logc"
)

type OllamaModel struct {
	Client *ollama.LLM
}

func NewOllamaModel() *OllamaModel {
	return &OllamaModel{}
}

func (oll *OllamaModel) GetModelClient(ctx context.Context, inter, name string) error {
	url := ollama.WithServerURL(inter)
	lla, err := ollama.New(ollama.WithModel(name), url)
	if err != nil {
		logc.Error(ctx, "ollama new client failed", err)
		return err
	}
	oll.Client = lla
	return nil
}

func (oll *OllamaModel) CallModel(ctx context.Context, prompt string, args ...interface{}) (string, error) {
	// call model, use prompt and role
	logc.Infof(ctx, "ollama call model with prompt: %s", prompt)
	res, err := oll.Client.Call(context.Background(), prompt)
	if err != nil {
		logc.Error(ctx, "ollama generate content failed", err)
		return "", err
	}
	return res, nil
}
