package simple

import (
	factory "github.com/penk110/design_pattern_go/creation_factory"
	"time"
)

func NewFileWriter() *fileWriter {

	return &fileWriter{}
}

type fileWriter struct {
}

func (f *fileWriter) Write(t time.Time, level factory.Level, msg []byte) error {

	return nil
}

func (f *fileWriter) Close() error {

	return nil
}
