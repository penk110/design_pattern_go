package creation_builder

import (
	"errors"
	"fmt"
	"time"
)

/*

 	Golang 中对于创建类参数比较多的对象的时候，我们常见的做法是必填参数直接传递，可选参数通过传递可变的方法进行创建。
	本文会先实现课程中的建造者模式，然后再实现我们常用的方式

	建造者模式的代码其实会很长，这些是它的一个缺点，所以如果不是参数的校验逻辑很复杂的
	情况下一般我们在 Go 中不会采用这种方式，而会采用后面的另外一种方式
*/

const (
	defaultMaxSize    = 4096
	defaultMaxIdle    = time.Second * 2
	defaultMinIdle    = time.Second
	defaultMaxTimeout = time.Second
)

type HttpClient struct {
	MaxSize    int
	MaxIdle    time.Duration
	MinIdle    time.Duration
	MaxTimeout time.Duration
}

type HttpClientBuilder struct {
	MaxSize    int
	MaxIdle    time.Duration
	MinIdle    time.Duration
	MaxTimeout time.Duration
}

func (hcb *HttpClientBuilder) SetMaxSize(maxSize int) error {
	if maxSize < 0 {
		return errors.New("must gt zero")
	}

	hcb.MaxSize = maxSize
	return nil
}

func (hcb *HttpClientBuilder) SetMaxIdle(maxIdle time.Duration) error {
	if maxIdle < 0 {
		return errors.New("must gt zero")
	}

	hcb.MaxIdle = maxIdle
	return nil
}

func (hcb *HttpClientBuilder) SetMaxTimeout(maxTimeout time.Duration) error {
	if maxTimeout < 0 {
		return errors.New("must gt zero")
	}

	hcb.MaxTimeout = maxTimeout
	return nil
}

func (hcb *HttpClientBuilder) Builder() (client *HttpClient, err error) {
	if hcb.MaxSize == 0 {
		hcb.MaxSize = defaultMaxSize
	}

	if hcb.MaxIdle == 0 {
		hcb.MaxIdle = defaultMaxIdle
	}

	if hcb.MinIdle == 0 {
		hcb.MinIdle = defaultMinIdle
	}

	if hcb.MaxTimeout == 0 {
		hcb.MaxTimeout = defaultMaxTimeout
	}

	// TODO: 参数校验
	if hcb.MaxSize < 0 {
		return nil, fmt.Errorf("MaxSize must gt zero")
	}

	if hcb.MinIdle > hcb.MaxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", hcb.MaxIdle, hcb.MinIdle)
	}

	client = &HttpClient{
		MaxSize:    hcb.MaxSize,
		MaxIdle:    hcb.MaxIdle,
		MinIdle:    hcb.MinIdle,
		MaxTimeout: hcb.MaxTimeout,
	}

	return client, nil
}
