package proxy

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestGenerate(t *testing.T) {
	want := `package proxy

type UserProxy struct {
	child *User
}

func NewUserProxy(child *User) *UserProxy {
	return &UserProxy{child: child}
}

func (p *UserProxy) Login(username string, password string) (r0 error) {
	start := time.Now()
	r0 = p.child.Login(username, password)
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return r0
}
`
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("[TestGenerate] get work directory failed, err: %s\n", err.Error())
		return
	}
	filePath := path.Join(pwd, "static_proxy.go")
	t.Logf("[TestGenerate] get work directory: %s", filePath)
	out, err := generate(filePath)
	require.Nil(t, err)
	assert.Equal(t, want, out)
}
