package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

type GprsChOngZhiAdvanceParams struct {
	OrderID    string `json:"orderid"`    // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	Account    string `json:"account"`    // 充值手机号 需要充值的手机号
	Gprs       int64  `json:"gprs"`       // 充值流量值 单位：MB（具体流量值请咨询商务）
	Area       int64  `json:"area"`       // 充值流量范围 0 全国流量，1 省内流量
	EffectTime int64  `json:"effecttime"` // 生效日期 0 即时生效，1次日生效，2 次月生效
	Validity   int64  `json:"validity"`   // 流量有效期 传入月数，0为当月有效
	Times      string `json:"times"`      // 时间戳 格式：yyyyMMddhhmmss
}

type GprsChOngZhiAdvanceResponse struct {
	XMLName    xml.Name `xml:"response"`
	UserID     string   `xml:"userid"`     // 会员账号
	OrderID    string   `xml:"orderid"`    // 会员提交订单号
	PorderID   string   `xml:"Porderid"`   // 平台订单号
	Account    string   `xml:"account"`    // 充值手机号
	State      int64    `xml:"state"`      // 订单状态
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
}

func newGprsChOngZhiAdvanceResult(result GprsChOngZhiAdvanceResponse, body []byte, http gorequest.Response) *GprsChOngZhiAdvanceResult {
	return &GprsChOngZhiAdvanceResult{Result: result, Body: body, Http: http}
}

// GprsChOngZhiAdvance 流量充值接口
// orderid = 用户提交的订单号	用户提交的订单号，最长32位（用户保证其唯一性）
// account = 充值手机号	需要充值的手机号
// gprs = 充值流量值	单位：MB（具体流量值请咨询商务）
// area = 充值流量范围	0 全国流量，1 省内流量
// effecttime = 生效日期	0 即时生效，1次日生效，2 次月生效
// validity = 流量有效期	传入月数，0为当月有效
func (c *Client) GprsChOngZhiAdvance(ctx context.Context, orderid string, account string, gprs int64, area int64, effectTime int64, validity int64, notMustParams ...gorequest.Params) (*GprsChOngZhiAdvanceResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "gprsChongzhiAdvance")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId())  // 用户编号
	params.Set("pwd", c.GetPwd())        // 加密密码
	params.Set("orderid", orderid)       // 用户提交的订单号	用户提交的订单号，最长32位（用户保证其唯一性）
	params.Set("account", account)       // 充值手机号	需要充值的手机号
	params.Set("gprs", gprs)             // 充值流量值	单位：MB（具体流量值请咨询商务）
	params.Set("area", area)             // 充值流量范围	0 全国流量，1 省内流量
	params.Set("effecttime", effectTime) // 生效日期	0 即时生效，1次日生效，2 次月生效
	params.Set("validity", validity)     // 流量有效期	传入月数，0为当月有效

	// 请求
	request, err := c.requestXml(ctx, "gprsChongzhiAdvance.do", params, http.MethodGet)
	if err != nil {
		return newGprsChOngZhiAdvanceResult(GprsChOngZhiAdvanceResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response GprsChOngZhiAdvanceResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newGprsChOngZhiAdvanceResult(response, request.ResponseBody, request), err
}
