package util

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// StopRedirect 用于停止重定向，返回错误信息。
func StopRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func GetFromData(data map[string]string) io.Reader {
	formData := url.Values{}
	for k, v := range data {
		formData.Add(k, v)
	}
	return strings.NewReader(formData.Encode())
}
