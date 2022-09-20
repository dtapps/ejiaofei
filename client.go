package ejiaofei

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	UserId           string
	Pwd              string
	Key              string
	ApiGormClientFun golog.ApiClientFun // 日志配置
	Debug            bool               // 日志开关
	ZapLog           *golog.ZapLog      // 日志服务
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
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

	c.zapLog = config.ZapLog

	c.config.userId = config.UserId
	c.config.pwd = config.Pwd
	c.config.key = config.Key

	c.requestClient = gorequest.NewHttp()

	apiGormClient := config.ApiGormClientFun()
	if apiGormClient != nil {
		c.log.client = apiGormClient
		c.log.status = true
	}

	return c, nil
}
