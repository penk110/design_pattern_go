package abstract_factory

import (
	factory "github.com/penk110/design_pattern_go/creation_factory"
	"github.com/penk110/design_pattern_go/creation_factory/factory_method"
	"github.com/penk110/design_pattern_go/creation_factory/simple"
)

/*
抽象工厂
	一个工厂方法可以创建相关联的多个类的时候就是抽象工厂模式，这个不太常用

*/

type AbsWriterInterface interface {
	CreateConsoleWriter() factory.Writer
	CreateNoneWriter() factory.Writer
}

type WriterFactory struct {
}

func (wf *WriterFactory) CreateConsoleWriter() factory.Writer {
	return &factory_method.ConsoleWriter{}
}

func (wf *WriterFactory) CreateNoneWriter() factory.Writer {
	return &simple.None{}
}

/*----------------------------------------------------------------------*/

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct{}

// Parse Parse
func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

/*----------------------------------------------------------------------*/

// ISystemConfigParser ISystemConfigParser
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

// jsonSystemConfigParser jsonSystemConfigParser
type jsonSystemConfigParser struct{}

// ParseSystem parse
func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

/*----------------------------------------------------------------------*/

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
