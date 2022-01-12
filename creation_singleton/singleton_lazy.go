package singleton

import "sync"

/*
懒汉式
	1.调用时检查是否已存在实例，无则先实例化；
	2.支持延迟加载；
	3.需要处理线程安全问题；
		双重锁，先检查是否为nil创建时加锁确保同一时刻只有一个线程在操作
*/

var Once = &sync.Once{}
var singletonLazy *Singleton

func GetLazy() *Singleton {
	if singletonLazy == nil {
		Once.Do(func() {
			singletonLazy = new(Singleton)
		})
	}

	return singletonLazy
}
