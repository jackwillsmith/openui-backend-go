package callmodel

import (
	"context"
	"github.com/tmc/langchaingo/llms"
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
	var msgChan = make(chan string)

	go oll.Generate(ctx, prompt, msgChan)

	return "", nil

	//c.Stream(func(w io.Writer) bool {
	//	select {
	//	case msg, ok := <-msgChan:
	//		if !ok {
	//			// 如果msgChan被关闭，则结束流式传输
	//			return false
	//		}
	//		fmt.Print(msg)
	//		c.SSEvent("message", msg)
	//		return true
	//	case <-c.Done():
	//		// 如果客户端连接关闭，则结束流式传输
	//		return false
	//	}
	//})
}

func (oll *OllamaModel) Generate(ctx context.Context, prompt string, msgChan chan string) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 设置超时
	// defer cancel()                                                          // 确保在函数结束时取消上下文

	callOp := llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		select {
		case msgChan <- string(chunk):
		case <-ctx.Done():
			return ctx.Err() // 返回上下文的错误
		}
		return nil
	})

	_, err := oll.Client.Call(ctx, prompt, callOp)
	if err != nil {
		logc.Errorf(ctx, "Call failed: %v", err) // 处理错误，而不是 panic
	}

	// 确保在所有数据处理完毕后关闭 msgChan
	close(msgChan)
}
