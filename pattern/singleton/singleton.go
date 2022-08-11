package singleton

import "sync"

// 【创建型】
// 单例模式：保证一个类只有一个实例，并提供一个全局访问点。
// 单例模式：饿汉式 和 懒汉式
// 饿汉式：在类加载的时候就创建一个实例，并且返回。
// 懒汉式：在第一次使用的时候创建一个实例，并且返回。

// Singleton 懒汉式
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
