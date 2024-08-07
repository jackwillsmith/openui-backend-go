// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/openui-backend-go/service/chat-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/ollama/api/chat",
				Handler: ChatHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ollama/api/tags",
				Handler: OllTagsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ollama/api/version",
				Handler: VersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ollama/v1/chat/completions",
				Handler: CompleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/changelog",
				Handler: ChangelogHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/config",
				Handler: ConfigHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/configs/default/models",
				Handler: GetDefaultModelsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/prompts",
				Handler: GetPromptHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/chats",
				Handler: ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/chats/:id",
				Handler: UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/chats/:id",
				Handler: RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/chats/:id",
				Handler: DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/chats/new",
				Handler: CreateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
