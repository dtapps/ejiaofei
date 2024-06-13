package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

type QueryTxProductResponse struct {
	XMLName xml.Name `xml:"response"`
	Error   string   `xml:"error"` // 错误提示
}

type QueryTxProductResult struct {
	Result QueryTxProductResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newQueryTxProductResult(result QueryTxProductResponse, body []byte, http gorequest.Response) *QueryTxProductResult {
	return &QueryTxProductResult{Result: result, Body: body, Http: http}
}

// QueryTxProduct 可充值腾讯产品查询
func (c *Client) QueryTxProduct(ctx context.Context, notMustParams ...gorequest.Params) (*QueryTxProductResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "queryTXproduct")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId()) // 用户编号
	params.Set("pwd", c.GetPwd())       // 加密密码

	// 请求
	request, err := c.requestXml(ctx, "queryTXproduct.do", params, http.MethodGet)
	if err != nil {
		return newQueryTxProductResult(QueryTxProductResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response QueryTxProductResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newQueryTxProductResult(response, request.ResponseBody, request), err
}
