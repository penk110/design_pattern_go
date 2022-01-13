package structure_face

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	var (
		user *User
		err  error
	)
	if user, err = GetUserBiz().Login("000001", "123321"); err != nil {
		t.Error(err.Error())
		return
	}
	assert.Equal(t, &User{Phone: "000001"}, user)
}

func TestRegister(t *testing.T) {
	var (
		user *User
		err  error
	)
	if user, err = GetUserBiz().Register("000002", "321123"); err != nil {
		t.Error(err.Error())
		return
	}
	assert.Equal(t, &User{Phone: "000002"}, user)
}

func TestLoginOrRegister(t *testing.T) {
	var (
		user *User
		err  error
	)
	if user, err = GetUserBiz().LoginOrRegister("000001", "321123"); err != nil {
		t.Error(err.Error())
		return
	}
	assert.Equal(t, &User{Phone: "000001"}, user)
}
