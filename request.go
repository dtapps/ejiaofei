package ejiaofei

import (
	"context"
	"fmt"
	"go.dtapp.net/gomd5"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, method string) (gorequest.Response, error) {

	// 公共参数
	params["userid"] = c.getUserId()
	params["pwd"] = c.getPwd()

	// 签名
	params["userkey"] = gomd5.ToUpper(fmt.Sprintf("%s%s", c.signStr, c.getKey()))

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Request()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.PgsqlDb != nil {
		go c.log.GormMiddlewareXml(ctx, request, Version)
	}
	if c.config.MongoDb != nil {
		go c.log.MongoMiddlewareXml(request)
	}

	return request, err
}
