package creation_builder

import (
	"fmt"
	"time"
)

type HttpConfigOption struct {
	MaxSize    int
	MaxIdle    time.Duration
	MinIdle    time.Duration
	MaxTimeout time.Duration
}

// HttpConfigOptionFunc to set option
type HttpConfigOptionFunc func(option *HttpConfigOption)

func NewHttpClientBuilder(optFuncList ...HttpConfigOptionFunc) (client *HttpClient, err error) {
	// 可直接设置默认值
	option := &HttpConfigOption{}

	// 通过参数构建方法装载 对象参数，方法参数必须是引用类型
	for _, opt := range optFuncList {
		opt(option)
	}

	if option.MaxSize == 0 {
		option.MaxSize = defaultMaxSize
	}

	if option.MaxIdle == 0 {
		option.MaxIdle = defaultMaxIdle
	}

	if option.MinIdle == 0 {
		option.MinIdle = defaultMinIdle
	}

	if option.MaxTimeout == 0 {
		option.MaxTimeout = defaultMaxTimeout
	}

	// TODO: 参数校验
	if option.MaxSize < 0 {
		return nil, fmt.Errorf("MaxSize must gt zero")
	}

	if option.MinIdle > option.MaxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", option.MaxIdle, option.MinIdle)
	}

	client = &HttpClient{
		MaxSize:    option.MaxSize,
		MaxIdle:    option.MaxIdle,
		MinIdle:    option.MinIdle,
		MaxTimeout: option.MaxTimeout,
	}

	return client, err
}
