package test

import (
	factory "github.com/zyphub/design_pattern_go/creation_factory"
	"github.com/zyphub/design_pattern_go/creation_factory/factory_method"
	"github.com/zyphub/design_pattern_go/creation_factory/simple"
	"testing"
	"time"
)

func TestFactory(t *testing.T) {

	fileFactory := simple.NewFileWriter()

	err := fileFactory.Write(time.Now(), factory.DEBUG, []byte("TestFactory"))
	if err != nil {
		t.Errorf("[TestFactory] err: %v\n", err)
		return
	}

	t.Logf("[TestFactory] write success.\n")

	_ = fileFactory.Close()
}

func TestFactoryMethod(t *testing.T) {
	var c factory.Writer
	consoleFactory := factory_method.ConsoleFactory{}
	c = consoleFactory.CreateConsoleWriter()

	err := c.Write(time.Now(), factory.DEBUG, []byte("TestFactoryMethod"))
	if err != nil {
		t.Errorf("[TestFactoryMethod] err: %v\n", err)
		return
	}

	t.Logf("[TestFactoryMethod] write success.\n")

	_ = c.Close()
}

func TestNoneWriter(t *testing.T) {
	none := simple.None{}
	err := none.Write(time.Now(), factory.DEBUG, []byte("TestFactoryMethod"))
	if err != nil {
		t.Errorf("[TestNoneWriter] err: %v\n", err)
		return
	}

	t.Logf("[TestNoneWriter] write success.\n")

	_ = none.Close()
}
