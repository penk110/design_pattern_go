package simple

import (
	factory "github.com/penk110/design_pattern_go/creation_factory"
	"time"
)

type None struct {
}

func (n *None) Write(t time.Time, level factory.Level, msg []byte) error {

	return nil
}

func (n *None) Close() error {

	return nil
}
