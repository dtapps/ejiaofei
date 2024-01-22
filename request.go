package ejiaofei

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (c *Client) requestXml(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 签名
	param.Set("userkey", c.xmlSign(url, param))

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.gormLog.status {
		go c.gormLog.client.MiddlewareXml(ctx, request)
	}
	if c.mongoLog.status {
		go c.mongoLog.client.MiddlewareXml(ctx, request)
	}

	return request, err
}

func (c *Client) requestJson(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 签名
	param.Set("sign", c.jsonSign(param))

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}

	return request, err
}
