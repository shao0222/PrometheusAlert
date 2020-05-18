package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type YZJMessage struct {
	Content string `json:"content"` // 消息体内容
}

// 发送消息至云之家
func PostToYunZhiJia(text, YzjUrl, logsign string) string {
	open := beego.AppConfig.String("open-yunzhijia")
	if open == "0" {
		logs.Info(logsign, "[yunzhijia]", "云之家接口未配置未开启状态,请先配置open-yunzhijia")
		return "云之家接口未配置未开启状态,请先配置open-yunzhijia"
	}
	u := YZJMessage{
		Content: text,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[yunzhijia]", b)
	var tr *http.Transport
	if proxyUrl := beego.AppConfig.String("proxy"); proxyUrl != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxy,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport: tr}
	res, err := client.Post(YzjUrl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[yunzhijia]", err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[yunzhijia]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("yunzhijia", text, "").Add(1)
	logs.Info(logsign, "[yunzhijia]", string(result))
	return string(result)
}
