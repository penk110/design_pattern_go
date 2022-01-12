package singleton

/*
饿汉式：
	1.预先创建；
	2.不支持延迟加载；
*/

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = new(Singleton)
}

func GetHungary() *Singleton {
	return singleton
}
