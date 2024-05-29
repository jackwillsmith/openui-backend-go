package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openui-backend-go/common/curlhttp"
	"github.com/openui-backend-go/common/utils"
	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OllTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOllTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OllTagsLogic {
	return &OllTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OllTagsLogic) OllTags() (resp *types.ModelReponse, err error) {
	// 构建 http requset结构体
	uri := utils.GetOllUrl()
	req := curlhttp.HeaderRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/api/tags", uri),
	}
	// 调用 Ollama url
	res, err := req.SendRequest()
	if err != nil {
		logc.Error(l.ctx, "request => GetWithParam: ", err)
		return nil, err
	}
	// 解析返回数据
	resp = &types.ModelReponse{}
	json.Unmarshal(res, resp)

	// 遍历 resp models, 将 id 为空的 model, id 设置为 name
	for i, model := range resp.Models {
		if model.Id == "" {
			resp.Models[i].Id = model.Name
		}
	}
	return
}
