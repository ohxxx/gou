package decorator

// 【结构型】
// 装饰器模式
// 		意图：动态的给一个对象添加一些额外的职责
// 		解决：将具体功能职责划分，同时继承装饰器模式
// 		实例：对蛋糕进行包装，其中分为：普通盒包装、塑料盒包装、玻璃盒包装

// 包装
type IPack interface {
	Desc() string
	Price() float32
}

// 普通盒包装
type Packing struct {
	name  string
	price float32
}

func (p Packing) Desc() string {
	return p.name
}

func (p Packing) Price() float32 {
	return p.price
}

// 塑料盒包装
type PlasticBoxPack struct {
	pack  IPack
	name  string
	price float32
}

func (p PlasticBoxPack) Desc() string {
	return p.pack.Desc() + p.name
}

func (p PlasticBoxPack) Price() float32 {
	return p.pack.Price() + p.price
}

// 玻璃盒包装
type GlassBoxPack struct {
	pack  IPack
	name  string
	price float32
}

func (p GlassBoxPack) Desc() string {
	return p.pack.Desc() + p.name
}

func (p GlassBoxPack) Price() float32 {
	return p.pack.Price() + p.price
}
