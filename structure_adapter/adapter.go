package structure_adapter

import "fmt"

type LogImp interface {
	Info(f string, args ...interface{})
}

type Console struct {
}

func (c *Console) Info(f string, args ...interface{}) {
	fmt.Printf(f, args...)
	return
}

type Non struct {
}

func (c *Non) Info(f string, args ...interface{}) {
	fmt.Printf(f, args...)
	return
}

type ConsoleAdapter struct {
	C *Console
}

func (ca *ConsoleAdapter) Info(f string, args ...interface{}) {
	// TODO: 新增debug、trace等逻辑
	fmt.Println("debug .... trace ...")
	fmt.Printf(f, args...)
	return
}
