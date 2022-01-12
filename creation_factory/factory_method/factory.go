package factory_method

import (
	factory "github.com/penk110/design_pattern_go/creation_factory"
	"time"
)

/*
当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，
推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂
*/

type ConsoleWriter struct {
}

func (c *ConsoleWriter) Write(t time.Time, level factory.Level, msg []byte) error {

	return nil
}

func (c *ConsoleWriter) Close() error {
	return nil
}

/*
	定义承载工厂方法的 接口 和 方法接收者 ConsoleFactory

	实例化出方法接收者后通过 create 方法创建出对应对象

*/

type ConsoleInterface interface {
	CreateConsoleWriter() factory.Writer
}

type ConsoleFactory struct {
}

func (cf ConsoleFactory) Write(t time.Time, level factory.Level, msg []byte) error {
	return nil
}

func (cf ConsoleFactory) Close() error {

	return nil
}

func (cf ConsoleFactory) CreateConsoleWriter() *ConsoleWriter {

	return &ConsoleWriter{}
}

/*
	writer 公共的工厂方法
*/

type WriterInterface interface {
	CreateWriter(wt string) factory.Writer
}

type WriterFactory struct {
}

func (w *WriterFactory) CreateWriter(wt string) factory.Writer {
	switch wt {
	case "none":
		// TODO: 初始化逻辑
		return nil
	default:
		// TODO: 初始化逻辑
		return &ConsoleWriter{}
	}
}
