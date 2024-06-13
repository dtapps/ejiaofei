package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/gorequest"
	"net/http"
)

type QueryJkOrdersResponse struct {
	XMLName   xml.Name `xml:"response"`
	UserID    string   `xml:"userid"`    // 会员账号
	POrderID  string   `xml:"Porderid"`  // 鼎信平台订单号
	OrderID   string   `xml:"orderid"`   // 用户订单号
	Account   string   `xml:"account"`   // 需要充值的手机号码
	Face      string   `xml:"face"`      // 充值面值
	Amount    string   `xml:"amount"`    // 购买数量
	StartTime string   `xml:"starttime"` // 开始时间
	State     string   `xml:"state"`     // 订单状态
	EndTime   string   `xml:"endtime"`   // 结束时间
	Error     string   `xml:"error"`     // 错误提示
}

type QueryJkOrdersResult struct {
	Result QueryJkOrdersResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newQueryJkOrdersResult(result QueryJkOrdersResponse, body []byte, http gorequest.Response) *QueryJkOrdersResult {
	return &QueryJkOrdersResult{Result: result, Body: body, Http: http}
}

// QueryJkOrders 通用查询接口
// orderid = 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
func (c *Client) QueryJkOrders(ctx context.Context, orderid string, notMustParams ...gorequest.Params) (*QueryJkOrdersResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "query_jkorders")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId()) // 用户编号
	params.Set("pwd", c.GetPwd())       // 加密密码
	params.Set("orderid", orderid)      // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）

	// 响应
	var response QueryJkOrdersResponse

	// 请求
	request, err := c.requestXml(ctx, "query_jkorders.do", params, http.MethodGet, &response)
	return newQueryJkOrdersResult(response, request.ResponseBody, request), err
}
