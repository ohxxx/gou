package prototype

import (
	"fmt"
	"testing"
)

var manager *ProtypeManager

type Name struct {
	name string
}

func (n *Name) Clone() Cloneable {
	tc := *n
	return &tc
}

func TestCloneFormManager(t *testing.T) {
	r, err := manager.Get("n1")

	if err != nil {
		t.Fatal("n1 does not exist")
	}

	fmt.Println(r)

	c := r.Clone()
	t1 := c.(*Name)
	if t1.name != "xxx" {
		t.Fatal("error")
	}
}

func init() {
	manager = NewProtypeManager()
	n := &Name{
		name: "xxx",
	}
	manager.Set("n1", n)
}
