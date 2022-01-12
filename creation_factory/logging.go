package creation_factory

import "time"

/*
工厂模式
	通过 NewXXX() 函数创建对象并返回
	单例也成为 静态工厂

*/

type Level int

const (
	DEBUG Level = iota
	INFO
	FATAL
	ERROR
	PANIC
)

type Writer interface {
	Write(t time.Time, level Level, msg []byte) error
	Close() error
}

