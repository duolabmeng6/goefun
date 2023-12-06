// Package eDingtalkBot Description: 钉钉流式机器人
package eDingtalkBot

import (
	"errors"
	"github.com/duolabmeng6/goefun/ecache"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/model/ejson"
)

func E钉钉取Token(appkey string, appsecret string) string {
	//url = f"https://oapi.dingtalk.com/gettoken?appkey={appkey}&appsecret={appsecret}"
	ec := ecache.NewFile缓存器("./cache/")
	get, err := ec.Get(appkey + appsecret)
	if err == nil {
		return ecore.E到文本(get)
	}
	h := ehttp.NewHttp()
	ret, err := h.Get("https://oapi.dingtalk.com/gettoken?appkey=" + appkey + "&appsecret=" + appsecret)
	if err != nil {
		return ""
	}
	// access_token
	access_token := ejson.Json解析文本(ret, "access_token")
	ec.Set(appkey+appsecret, access_token, 7100)

	return access_token
}

func E钉钉上传文件(token string, path string) (string, error) {
	//    url = f'https://oapi.dingtalk.com/media/upload?access_token={access_token}'
	//    files = {
	//        'media': open(file_path, 'rb'),
	//        'type': (None, 'file')
	//    }
	//    response = requests.post(url, files=files)
	//    return response.text
	h := ehttp.NewHttp()
	//datas := make(map[string]interface{})
	//datas["type"] = "file"
	//datas["media"] = "@file:" + path

	ret, err := h.E访问("https://oapi.dingtalk.com/media/upload?access_token="+token, "POST", "type=file&media="+"@file:"+path, "")
	if err != nil {
		return "", err
	}
	//print(string(ret))
	// {"errcode":0,"errmsg":"ok","media_id":"@lAjPDfmVdtXMN-DOfcm4p856y8EH","created_at":1700808802015,"type":"file"}
	// {"errcode":60020,"errmsg":"访问ip不在白名单之中，请参考FAQ：https://open.dingtalk.com/document/org-faq/app-faq,request ip=120.230.136.230 appKey\u0028dingmf6cuictoilijt0g\u0029"}
	media_id := ejson.Json解析文本(string(ret), "media_id")
	if media_id == "" {

		return "", errors.New(ejson.Json解析文本(string(ret), "errmsg"))
	}

	return media_id, nil
}
