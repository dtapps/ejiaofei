package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func (c *Client) requestXml(ctx context.Context, url string, param gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 签名
	param.Set("userkey", c.xmlSign(url, param))

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(url)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeForm()

	// 设置参数
	c.httpClient.SetParams(param)

	// OpenTelemetry链路追踪
	c.TraceSetAttributes(attribute.String("http.url", uri))
	c.TraceSetAttributes(attribute.String("http.method", method))
	c.TraceSetAttributes(attribute.String("http.params", gojson.JsonEncodeNoError(param)))

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return gorequest.Response{}, err
	}

	// 解析响应
	err = xml.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}

	return request, err
}

func (c *Client) requestJson(ctx context.Context, url string, param gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 签名
	param.Set("sign", c.jsonSign(param))

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeForm()

	// 设置参数
	c.httpClient.SetParams(param)

	// OpenTelemetry链路追踪
	c.TraceSetAttributes(attribute.String("http.url", uri))
	c.TraceSetAttributes(attribute.String("http.method", method))
	c.TraceSetAttributes(attribute.String("http.params", gojson.JsonEncodeNoError(param)))

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return gorequest.Response{}, err
	}

	// 解析响应
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}

	return request, err
}
