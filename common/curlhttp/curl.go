package curlhttp

// 定义一个http get 方法

import (
	"bytes"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"io/ioutil"
	"net/http"
	"strings"
)

type HeaderRequest struct {
	Method      string
	Url         string
	Header      map[string]string
	Body        string
	ContentType string
}

type CommonResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func (headerRequest *HeaderRequest) SendRequest() ([]byte, error) {
	var req *http.Request
	var err error
	switch strings.ToUpper(headerRequest.Method) {
	case "GET":
		req, err = http.NewRequest("GET", headerRequest.Url, nil)
	case "POST":
		jsonBody := []byte(headerRequest.Body)
		req, err = http.NewRequest("POST", headerRequest.Url, bytes.NewBuffer(jsonBody))
	default:
		req = nil
		err = errors.New("wrong method")
	}

	if err != nil {
		return nil, err
	}
	if headerRequest.ContentType == "" {
		headerRequest.ContentType = "application/json"
	}
	req.Header.Set("Content-Type", headerRequest.ContentType)
	header := headerRequest.Header
	for i := range header {
		req.Header.Add(i, header[i])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logc.Error(context.Background(), "request => GetWithParam: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
