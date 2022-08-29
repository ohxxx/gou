package prototype

import "testing"

func TestPrototype(t *testing.T) {
	p := new(People)
	p.setName("xxx")
	p.Eat()

	p1 := p.Clone()
	p1.setName("yyy")
	p1.Eat()
	p1.Sleep()
	p1.PlayPeas()
}
