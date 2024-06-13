package ejiaofei

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

type FindBalanceResponse struct {
	Code    int64   `json:"code"`    // 返回状态编码
	Balance float64 `json:"balance"` // 用户余额
}

type FindBalanceResult struct {
	Result FindBalanceResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newFindBalanceResult(result FindBalanceResponse, body []byte, http gorequest.Response) *FindBalanceResult {
	return &FindBalanceResult{Result: result, Body: body, Http: http}
}

// FindBalance 余额查询接口
func (c *Client) FindBalance(ctx context.Context, notMustParams ...gorequest.Params) (*FindBalanceResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "findBalance")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appId", c.GetUserId())  // 用户编号 由鼎信商务提供
	params.Set("appSecret", c.GetPwd()) // 加密密码 由鼎信商务提供

	// 请求
	request, err := c.requestJson(ctx, "findBalance.do", params, http.MethodGet)
	if err != nil {
		return newFindBalanceResult(FindBalanceResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response FindBalanceResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newFindBalanceResult(response, request.ResponseBody, request), err
}
