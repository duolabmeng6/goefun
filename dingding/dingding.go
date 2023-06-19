// Package dingding 钉钉机器人
package dingding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type E钉钉机器人 struct {
	WebHookUrl string
	Secret     string
}

type response struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func New钉钉机器人(WebHookUrl, secret string) *E钉钉机器人 {
	return &E钉钉机器人{WebHookUrl: WebHookUrl, Secret: secret}
}

// E发送文本消息
// text 发送内容
// atIdType 1 atMobiles 2 atUserIds 3 all
// at 传入的手机号或者用户id
func (t *E钉钉机器人) E发送文本消息(text string, atIdType int, at ...string) error {

	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": text,
		},
	}
	if atIdType == 1 {
		msg["at"] = map[string]interface{}{
			"atMobiles": at,
			"isAtAll":   false,
		}
	}
	if atIdType == 2 {
		msg["at"] = map[string]interface{}{
			"atUserIds": at,
			"isAtAll":   false,
		}
	}
	if atIdType == 3 {
		msg["at"] = map[string]interface{}{
			"isAtAll": true,
		}
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	//fmt.Println("发送数据", bytes.NewBuffer(b).String())
	resp, err := http.Post(t.getURL(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}
	if r.Code != 0 {
		return errors.New(fmt.Sprintf("response error: %s", string(body)))
	}
	return err
}

// E发送markdown消息
// text 发送内容
// atIdType 1 atMobiles 2 atUserIds 3 all
// at 传入的手机号或者用户id
func (t *E钉钉机器人) E发送markdown消息(title, text string, atIdType int, at ...string) error {
	msg := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": title,
			"text":  text,
		},
	}

	if atIdType == 1 {
		msg["at"] = map[string]interface{}{
			"atMobiles": at,
			"isAtAll":   false,
		}
	}
	if atIdType == 2 {
		msg["at"] = map[string]interface{}{
			"atUserIds": at,
			"isAtAll":   false,
		}
	}
	if atIdType == 3 {
		msg["at"] = map[string]interface{}{
			"isAtAll": true,
		}
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(t.getURL(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return err
}

func (t *E钉钉机器人) hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (t *E钉钉机器人) getURL() string {
	wh := t.WebHookUrl
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, t.Secret)
	sign := t.hmacSha256(stringToSign, t.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", wh, timestamp, sign)
	return url
}
