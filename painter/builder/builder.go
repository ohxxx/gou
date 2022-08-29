package builder

//【创建型】
// 建造者模式
// 		意图：将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。
// 		解决：复杂对象的各个部分经常面临着不同的变化，但将它们组合在一起的算法却相对稳定。
// 		实例：创建一个房子都需要：材料、颜色、面积，但可以通过不同的建造者来创建不同的房子。

type IHouse interface {
	Material(string) // 房子材料
	Color(string)    // 房子颜色
	Area(float32)    // 房子面积
}

type House struct {
	material string
	color    string
	area     float32
}

func (h *House) Material(m string) {
	h.material = m
}

func (h *House) Color(c string) {
	h.color = c
}

func (h *House) Area(a float32) {
	h.area = a
}

type IBuilder interface {
	Material(string) IBuilder
	Color(string) IBuilder
	Area(float32) IBuilder
	Build() IHouse
}

type Builder struct {
	house IHouse
}

func (b Builder) Material(m string) IBuilder {
	b.house.Material(m)
	return b
}

func (b Builder) Color(c string) IBuilder {
	b.house.Color(c)
	return b
}

func (b Builder) Area(a float32) IBuilder {
	b.house.Area(a)
	return b
}

func (b Builder) Build() IHouse {
	return b.house
}
