package creation_prototype

import (
	factory "github.com/zyphub/design_pattern_go/creation_factory"
	"os"
	"sync"
	"time"
)

/*
	原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象

	原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置
*/

type Writer interface {
	Write(t time.Time, level factory.Level, msg []byte) error
	ReadToBuffer(buffer []byte) (n int, err error)
	Clone() Writer
	Close() error
}

type fileWriter struct {
	filepath string
	file     *os.File
	mutex    *sync.Mutex
}

func NewPrototypeManage(filepath string) (*fileWriter, error) {
	var err error
	f := &fileWriter{
		filepath: filepath,
		file:     nil,
		mutex:    nil,
	}
	obj := &fileWriter{
		filepath: f.filepath,
		mutex:    &sync.Mutex{},
	}
	// TODO: check file exist or not
	obj.file, err = os.Create(f.filepath)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (f *fileWriter) Clone() Writer {
	var obj = *f
	return &obj
}

func (f *fileWriter) Write(t time.Time, level factory.Level, msg []byte) error {
	// logics
	f.mutex.Lock()
	defer f.mutex.Unlock()
	_, err := f.file.Write(msg)
	if err != nil {
		return err
	}
	return nil
}

func (f *fileWriter) ReadToBuffer(buffer []byte) (n int, err error) {
	// logics
	return f.file.Read(buffer)

}

func (f *fileWriter) Close() error {
	// logics
	return nil
}
