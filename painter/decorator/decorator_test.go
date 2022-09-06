package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	p := &Packing{name: "普通蛋糕", price: 12.00}
	pp := &PlasticBoxPack{pack: p, price: 5.00}
	gp := &GlassBoxPack{pack: pp, price: 10.00}

	fmt.Printf("名称: %s, 价格: %f\n", gp.Desc(), gp.Price())
	// log：名称: 普通蛋糕, 价格: 27.000000
	// 普通蛋糕 + 一个塑料盒子 + 一个玻璃盒子 = 27
	//

	t.Log("DONE")
}
