package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/gorequest"
	"net/http"
)

type MoneyJkUserResponse struct {
	XMLName   xml.Name `xml:"response"`
	LastMoney float64  `xml:"lastMoney"` // 用户余额
	Tag       int64    `xml:"tag"`       // 用户状态（0正常 1暂停）
	Error     int64    `xml:"error"`     // 错误提示
}

type MoneyJkUserResult struct {
	Result MoneyJkUserResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newMoneyJkUserResult(result MoneyJkUserResponse, body []byte, http gorequest.Response) *MoneyJkUserResult {
	return &MoneyJkUserResult{Result: result, Body: body, Http: http}
}

// MoneyJkUser 用户余额查询
func (c *Client) MoneyJkUser(ctx context.Context, notMustParams ...gorequest.Params) (*MoneyJkUserResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId()) // 用户编号
	params.Set("pwd", c.GetPwd())       // 加密密码
	// 请求
	request, err := c.requestXml(ctx, apiUrl+"/money_jkuser.do", params, http.MethodGet)
	if err != nil {
		return newMoneyJkUserResult(MoneyJkUserResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MoneyJkUserResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newMoneyJkUserResult(response, request.ResponseBody, request), err
}
