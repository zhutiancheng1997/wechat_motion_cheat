package util

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
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

func GetRandInt64(n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(n)

}

func FakeIP() string {
	rand.Seed(time.Now().UnixNano())

	// 随便找的国内IP段：223.64.0.0 - 223.117.255.255
	ip := fmt.Sprintf("223.%d.%d.%d", rand.Intn(54)+64, rand.Intn(254), rand.Intn(254))
	return ip
}
