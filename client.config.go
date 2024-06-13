package ejiaofei

func (c *Client) SetStoreId(storeId int64) *Client {
	c.config.storeId = storeId
	return c
}

func (c *Client) SetAppKey(appKey int64) *Client {
	c.config.appKey = appKey
	return c
}

func (c *Client) SetAppSecret(appSecret string) *Client {
	c.config.appSecret = appSecret
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
