package prototype

import "fmt"

// 【创建型】
// 原型模式
// 		意图：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象
// 		解决：利用已有的一个原型对象，快速的生产和原型对象一样的实例
// 		实例：克隆一个人，拥有人类吃饭、睡觉、打豆豆的功能

type IPerson interface {
	Eat()
	Sleep()
	PlayPeas()
}

type People struct {
	name string
}

func (p *People) Eat() {
	fmt.Println(p.name, "在吃饭")
}

func (p *People) Sleep() {
	fmt.Println(p.name, "在睡觉")
}

func (p *People) PlayPeas() {
	fmt.Println(p.name, "在打豆豆")
}

func (p *People) setName(name string) {
	p.name = name
}

func (p *People) Clone() *People {
	if p == nil {
		return nil
	}

	new_people := (*p)
	return &new_people
}
