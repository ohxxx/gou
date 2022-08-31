package proxy

import (
	"fmt"
	"time"
)

// 【结构型】
// 代理模式
// 		意图：为其他对象提供一种代理以控制对这个对象的访问
// 		解决：在直接访问对象时带来的问题（不安全等因素）
// 		实例：在用户登录的时候进行统计、监控
//  	注意事项：
// 			1、和适配器模式区别：适配器模式主要改变所考虑对象的接口，代理模式不能改变所代理的接口
//			2、和装饰器模式区别：装饰器模式为了增强功能，代理模式是为了加以控制
//			3、和中介者模式区别：中介者模式为了减少对象之间的相互耦合

type IUser interface {
	Login(username, password string) error
}

type User struct{}

func (u *User) Login(username, password string) error {
	// ...
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{user: user}
}

func (u *UserProxy) Login(username, password string) error {
	// 统计...
	start := time.Now()

	if err := u.user.Login(username, password); err != nil {
		return err
	}

	// 监控...

	fmt.Printf("用户登录时间: %s\n", time.Now().Sub(start))

	return nil
}
