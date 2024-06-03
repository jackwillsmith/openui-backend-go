package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/openui-backend-go/common/consts"

	"github.com/openui-backend-go/common/callmodel"
	"github.com/openui-backend-go/common/utils"
	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewChatEntity
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 设置HTTP头部，告诉客户端我们将发送的是流式数据
		// 设置HTTP状态码和头部，指定内容类型
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// 使用 Flush 来通知客户端HTTP头已经发送
		flusher, _ := w.(http.Flusher)
		defer flusher.Flush()

		var msgChan = make(chan string)

		// new Ollama client
		oll := callmodel.NewOllamaModel()

		oll.GetModelClient(r.Context(), utils.GetOllUrl(), req.Model)

		var prompt string = ""
		for _, message := range req.Messages {
			if message.Role == "user" {
				prompt = message.Content
			}
		}

		fmt.Println("prompt:", prompt)

		go oll.Generate(r.Context(), prompt, msgChan)

		// 启动一个协程来发送数据
		for {
			select {
			case msg, ok := <-msgChan:
				if !ok {
					// 如果msgChan被关闭，则结束流式传输
					return
				}
				// fmt.Print(msg)
				msgResp := consts.MessageEntity{
					Content: msg,
					Role:    consts.ASSISTANT,
				}
				steamRes := consts.MessageResponse{
					Model:     req.Model,
					CreatedAt: utils.GetNowTime().String(),
					Message:   msgResp,
					Done:      false,
				}
				mStreamRes, _ := json.Marshal(steamRes)
				// 将 /n 添加到 mStreamRes 末尾
				mStreamRes = append(mStreamRes, '\n')
				if _, err := w.Write(mStreamRes); err != nil {
					return
				}
				//if _, err := w.Write([]byte(msg)); err != nil {
				//	fmt.Println("write error:", err)
				//	return
				//}
				flusher.Flush()
			case <-r.Context().Done():
				// 如果客户端连接关闭，则结束流式传输
				return
			}
		}
	}
}
