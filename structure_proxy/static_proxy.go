package proxy

import (
	"log"
	"time"
)

/*
	静态代理
*/

type UserImp interface {
	Login(username string, password string) error
}

// User user imp userImp
// @proxy UserImp
type User struct {
}

// Login login logics
func (user *User) Login(username string, password string) error {
	// logics

	return nil
}

// UserProxy user proxy
type UserProxy struct {
	user *User
}

// NewUserProxy new user proxy
func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

// Login 实现受代理对象的相关方法
func (p *UserProxy) Login(username, password string) error {
	// before: validate logic
	// eg: calculate cost time
	start := time.Now()

	// logics
	if err := p.user.Login(username, password); err != nil {
		return err
	}

	// after: auth logics
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
