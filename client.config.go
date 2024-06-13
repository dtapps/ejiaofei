package ejiaofei

import "go.dtapp.net/gorequest"

func (c *Client) SetUserId(userId string) *Client {
	c.config.userId = userId
	return c
}

func (c *Client) SetPwd(pwd string) *Client {
	c.config.pwd = pwd
	return c
}

func (c *Client) SetKey(key string) *Client {
	c.config.key = key
	return c
}

// SetClientIP 配置
func (c *Client) SetClientIP(clientIP string) *Client {
	c.clientIP = clientIP
	if c.httpClient != nil {
		c.httpClient.SetClientIP(clientIP)
	}
	return c
}

// SetLogFun 设置日志记录函数
func (c *Client) SetLogFun(logFun gorequest.LogFunc) {
	if c.httpClient != nil {
		c.httpClient.SetLogFunc(logFun)
	}
}
