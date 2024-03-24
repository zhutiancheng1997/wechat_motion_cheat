package manager

import (
	"crypto/tls"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/zhutiancheng1997/wechat_motion_cheat/model"
	"github.com/zhutiancheng1997/wechat_motion_cheat/util"
	"io"
	"net/http"
	"net/url"
)

type ProxyManager struct {
}

// todo add error return
func (p *ProxyManager) getProxy() *model.GetProxyResp {
	req, _ := http.NewRequest("GET", "http://localhost:5010/get?type=https", nil)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.WithError(err).Error("get proxy err")
		return nil
	}
	bs, _ := io.ReadAll(resp.Body)
	proxyResp := &model.GetProxyResp{}

	_ = json.Unmarshal(bs, proxyResp)
	log.WithField("proxyResp", proxyResp).Info("proxyResp")
	return proxyResp
}

func (p *ProxyManager) getProxyCount() {
	req, _ := http.NewRequest("GET", "127.0.0.1:5010/count", nil)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.WithError(err).Error("get proxy err")
		return
	}
	bs, _ := io.ReadAll(resp.Body)
	proxyResp := &model.GetProxyCountResp{}

	_ = json.Unmarshal(bs, proxyResp)
	log.WithField("proxyCntResp", proxyResp).Info("proxyCntResp")

}

type shuazuanResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (p *ProxyManager) DoGetDy1000Play(proxy *model.GetProxyResp) error {

	proxyURL, _ := url.Parse("http://" + proxy.Proxy)
	//proxyURL, _ := url.Parse(proxy.Proxy)
	//
	//// 创建一个自定义的 Transport，并设置代理
	transport := &http.Transport{

		//Proxy:           http.ProxyURL(proxyURL),
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	data := map[string]string{
		"tid":        "3210",
		"inputvalue": "https://v.douyin.com/iFCUTy4D/",
		"num":        "1",
	}
	req, _ := http.NewRequest("POST", "https://qwq1.qcw023.com/ajax.php?act=pay", util.GetFromData(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("Referer", "https://qwq1.qcw023.com/")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Host", "qwq1.qcw023.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	cli := &http.Client{Transport: transport}
	//cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.WithError(err).Error("doGetDy1000Play err")
		return err
	}
	defer resp.Body.Close()
	defer cli.CloseIdleConnections()
	bs, _ := io.ReadAll(resp.Body)
	respStruct := &shuazuanResp{}

	_ = json.Unmarshal(bs, respStruct)
	log.WithField("respStruct", respStruct).Info("respStruct")
	return nil
}
