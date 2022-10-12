package ejiaofei

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	UserId string
	Pwd    string
	Key    string
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		userId  string
		pwd     string
		key     string
		signStr string // 需要签名的字符串
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.userId = config.UserId
	c.config.pwd = config.Pwd
	c.config.key = config.Key

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
