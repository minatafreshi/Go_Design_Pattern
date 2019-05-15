package structural


type Flyweight struct {
	Name string
}

// NewFlyweight creates a new Flyweight object.
func NewFlyweight(name string) *Flyweight {
	return &Flyweight{name}
}

// FlyweightFactory is a factory for creating and storing flyweights.
type FlyweightFactory struct {
	pool map[string]*Flyweight
}

// NewFlyweightFactory creates a new FlyweightFactory.
func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{pool: make(map[string]*Flyweight)}
}


func (f *FlyweightFactory) GetFlyweight(name string) *Flyweight {
	flyweight, okay := f.pool[name]
	if !okay {
		flyweight = NewFlyweight(name)
		f.pool[name] = flyweight
	}
	return flyweight
}