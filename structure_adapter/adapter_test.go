package structure_adapter

import (
	"testing"
)

func TestAdapterInfo(t *testing.T) {
	// adapter 必须实现目标接口
	var (
		console        *Console
		consoleAdapter *ConsoleAdapter
		f              string
	)
	console = &Console{}
	consoleAdapter = &ConsoleAdapter{}
	f = "struct: %s\n"

	console.Info(f, "console")

	consoleAdapter.Info(f, "console adapter")
}
