package builder

import "testing"

func TestBuild(t *testing.T) {
	house := &House{}
	builderHouse := &Builder{house: house}
	h := builderHouse.Material("marble").Color("lightblue").Area(120.21)
	t.Logf("%+v\n", h.Build())
	// log：&{material:marble color:lightblue area:120.21}
}
