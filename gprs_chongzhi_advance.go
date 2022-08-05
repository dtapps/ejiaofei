package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type GprsChOngZhiAdvanceParams struct {
	OrderID    string `json:"orderid"`    // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	Account    string `json:"account"`    // 充值手机号 需要充值的手机号
	Gprs       int    `json:"gprs"`       // 充值流量值 单位：MB（具体流量值请咨询商务）
	Area       int    `json:"area"`       // 充值流量范围 0 全国流量，1 省内流量
	EffectTime int    `json:"effecttime"` // 生效日期 0 即时生效，1次日生效，2 次月生效
	Validity   int    `json:"validity"`   // 流量有效期 传入月数，0为当月有效
	Times      string `json:"times"`      // 时间戳 格式：yyyyMMddhhmmss
}

type GprsChOngZhiAdvanceResponse struct {
	XMLName    xml.Name `xml:"response"`
	UserID     string   `xml:"userid"`     // 会员账号
	OrderID    string   `xml:"orderid"`    // 会员提交订单号
	PorderID   string   `xml:"Porderid"`   // 平台订单号
	Account    string   `xml:"account"`    // 充值手机号
	State      int      `xml:"state"`      // 订单状态
	StartTime  string   `xml:"starttime"`  // 开始时间
	EndTime    string   `xml:"endtime"`    // 结束时间
	Error      string   `xml:"error"`      // 错误提示
	UserPrice  float64  `xml:"userprice"`  // 会员购买价格
	Gprs       string   `xml:"gprs"`       // 充值流量值（单位MB）
	Area       string   `xml:"area"`       // 流量范围（0 全国流量，1省内流量）
	EffectTime string   `xml:"effecttime"` // 生效日期（0即时，1次日，2次月）
	Validity   string   `xml:"validity"`   // 流量有效期（显示月数，0为当月）
}

type GprsChOngZhiAdvanceResult struct {
	Result GprsChOngZhiAdvanceResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
	Err    error                       // 错误
}

func newGprsChOngZhiAdvanceResult(result GprsChOngZhiAdvanceResponse, body []byte, http gorequest.Response, err error) *GprsChOngZhiAdvanceResult {
	return &GprsChOngZhiAdvanceResult{Result: result, Body: body, Http: http, Err: err}
}

// GprsChOngZhiAdvance 流量充值接口
func (c *Client) GprsChOngZhiAdvance(notMustParams ...gorequest.Params) *GprsChOngZhiAdvanceResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 签名
	c.signStr = fmt.Sprintf("userid%vpwd%vorderid%vaccount%vgprs%varea%veffecttime%vvalidity%vtimes%v", c.getUserId(), c.getPwd(), params["orderid"], params["account"], params["gprs"], params["area"], params["effecttime"], params["validity"], params["times"])
	// 请求
	request, err := c.request(apiUrl+"/gprsChongzhiAdvance.do", params, http.MethodGet)
	// 定义
	var response GprsChOngZhiAdvanceResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newGprsChOngZhiAdvanceResult(response, request.ResponseBody, request, err)
}
