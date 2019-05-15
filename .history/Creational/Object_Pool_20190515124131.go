package Creational

import (
	"sync"
)

type Pool struct {
	sync.Mutex
	inuse     []interface{}
	available []interface{}
	new       func() interface{}
}

func NewPool(new func() interface{}) *Pool {
	return &Pool{new: new}
}

func (p *Pool) Acquire() interface{} {
	p.Lock()
	var object interface{}
	if len(p.available) != 0 {
		object = p.available[0]
		p.available = append(p.available[:0], p.available[1:]...)
		p.inuse = append(p.inuse, object)
	} else {
		object = p.new()
		p.inuse = append(p.inuse, object)
	}
	p.Unlock()
	return object
}

func (p *Pool) Release(object interface{}) {
	p.Lock()
	p.available = append(p.available, object)
	for i, v := range p.inuse {
		if v == object {
			p.inuse = append(p.inuse[:i], p.inuse[i+1:]...)
			break
		}
	}
	p.Unlock()
}		