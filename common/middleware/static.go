package middleware

import (
	"github.com/zeromicro/go-zero/rest"
	"io/ioutil"
	"net/http"
)

// main直接调用
// 静太文件处理
func MiddelwareStatic(server *rest.Server) {
	staticFileHandler(server)
}

// 定义函数
func staticFileHandler(engine *rest.Server) {
	//这里注册
	patern := "web"
	dirpath := "web/static/"

	rd, _ := ioutil.ReadDir(dirpath)

	//添加进路由最后生成 /asset
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/index.html",
				Handler: dirhandler("index.html", patern),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: dirhandler("/", patern),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favicon.png",
				Handler: dirhandler("/favicon.png", patern),
			},
		})
	for _, f := range rd {
		filename := f.Name()
		path := "/static/" + filename
		//最后生成 /asset
		engine.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: dirhandler("/static/", dirpath),
			})
	}

}

func dirhandler(patern, filedir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))
		handler.ServeHTTP(w, req)
	}
}
