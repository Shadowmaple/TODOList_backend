package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"todolist_backend/util"
)

const (
	codeToSessionURL = "https://api.q.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type CodeToSessionPayload struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，一般没有权限获取
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

/*********************************************************************************
errcode 的合法值：
-1         => 系统繁忙，此时请开发者稍候再试
0          => 请求成功
40029      => code 无效
45011      => 频率限制，每个用户每分钟100次
-101222100 => 参数错误，请检查appid和appsecret是否正确,请检查ide上创建工程用的appid是否正确
-2100      => 参数错误，请检查appid和appsecret是否正确,请检查ide上创建工程用的appid是否正确
***********************************************************************************/

// 通过 code，获取 openId 和 sessionKey
func CodeToSession(code string) (string, string, error) {
	appID, appSecret := util.GetAppIDAndAppSecret()
	url := fmt.Sprintf(codeToSessionURL, appID, appSecret, code)

	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	payload := &CodeToSessionPayload{}
	if err := json.NewDecoder(resp.Body).Decode(payload); err != nil {
		return "", "", err
	}

	if payload.ErrCode != 0 {
		return "", "", errors.New(fmt.Sprintf("%d %s", payload.ErrCode, payload.ErrMsg))
	}

	return payload.OpenID, payload.SessionKey, nil
}
