package structure_face

/*
	将多个细粒度的接口封装成一个接口对外服务
*/

type User struct {
	Name  string
	Phone string
}

type IUser interface {
	Login(phone string, code string) (*User, error)
	Register(phone string, code string) (*User, error)
}

type UserBiz struct {
}

var _userBiz = &UserBiz{}

func GetUserBiz() *UserBiz {
	return _userBiz
}

func (userBiz *UserBiz) Login(phone string, code string) (*User, error) {
	// 用户校验
	return &User{
		Phone: "000001",
	}, nil
}

func (userBiz *UserBiz) Register(phone string, code string) (*User, error) {

	// 注册逻辑
	// 新建用户

	return &User{
		Phone: "000002",
	}, nil
}

type UserFaceBiz interface {
	LoginOrRegister(phone string, code string) (*User, error)
}

// LoginOrRegister 登录或注册
func (userBiz *UserBiz) LoginOrRegister(phone string, code string) (*User, error) {
	var (
		user *User
		err  error
	)
	if user, err = userBiz.Login(phone, code); err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	return userBiz.Register(phone, code)
}
