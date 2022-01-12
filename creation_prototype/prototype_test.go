package creation_prototype

import (
	factory "github.com/penk110/design_pattern_go/creation_factory"
	"os"
	"path"
	"testing"
	"time"
)

var fw *fileWriter

func TestPrototype(t *testing.T) {
	var err error
	tempString := "TestPrototype"
	basePath, _ := os.Getwd()
	filepath := path.Join(basePath, "temp.log")
	t.Logf("[TestPrototype] base path: %s\n", basePath)
	fw, err = NewPrototypeManage(filepath)
	if err != nil {
		t.Errorf("[TestPrototype] create prototype manager failed, err: %v\n", err.Error())
		return
	}
	t.Logf("[TestPrototype] file path: %v\n", fw.filepath)

	err = fw.Write(time.Now(), factory.DEBUG, []byte(tempString+"\n"))
	if err != nil {
		t.Errorf("[TestPrototype] fw write failed, err: %v\n", err.Error())
		return
	}

	fwClone := fw.Clone()

	buffer := []byte{}
	n, err := fwClone.ReadToBuffer(buffer)
	if err != nil {
		t.Errorf("[TestPrototype] fw read failed, err: %v\n", err.Error())
	}
	t.Logf("[TestPrototype] n: %d, data: %v\n", n, string(buffer))

}
