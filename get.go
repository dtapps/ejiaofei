package ejiaofei

import "go.dtapp.net/golog"

func (c *Client) GetUserId() string {
	return c.config.userId
}

func (c *Client) GetPwd() string {
	return c.config.pwd
}

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) GetLogGorm() *golog.ApiClient {
	return c.log.logGormClient
}

func (c *Client) GetLogMongo() *golog.ApiClient {
	return c.log.logMongoClient
}
