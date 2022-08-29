package prototype

import "errors"

type Cloneable interface {
	Clone() Cloneable
}

type ProtypeManager struct {
	prototypes map[string]Cloneable
}

func NewProtypeManager() *ProtypeManager {
	return &ProtypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *ProtypeManager) Get(name string) (Cloneable, error) {
	if p.prototypes[name] == nil {
		return nil, errors.New("name does not exist")
	}
	return p.prototypes[name].Clone(), nil
}

func (p *ProtypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
