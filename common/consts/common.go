package consts

import (
	"google.golang.org/grpc/codes"
	"strings"
)

type CommonResponse struct {
	Code codes.Code  `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 状态短语
	Data interface{} `json:"data"` // 数据结果集
}

// StatusText 状态码对应信息
func StatusText(code codes.Code, v ...string) string {
	//默认中文
	wrongMsg := WrongMessageZh[code]
	if len(v) > 0 {
		wrongMsg = WrongMessageEn[code]
		lang := v[0] //第一个参数是语言
		if lang == "" || lang == "zh" || lang == "cn" {
			wrongMsg = WrongMessageZh[code]
		}

		if len(v) > 1 {
			wrongMsg += ":"
			for i := 1; i < len(v); i++ {
				wrongMsg += v[i] + ","
			}
			wrongMsg = strings.TrimRight(wrongMsg, ",")
		}
	}
	return wrongMsg
}

func ResponseNewCodeMsg(code codes.Code, v ...string) error {
	return &CommonResponse{
		Code: code,
		Msg:  StatusText(code, v...),
	}
}

func ResponseSuccess(v ...string) error {
	return ResponseNewCodeMsg(SUCCESS, v...)
}

func ResponseFail(v ...string) error {
	return ResponseNewCodeMsg(FAIL, v...)
}

func (err *CommonResponse) Error() string {
	return err.Msg
}
