package database

import (
	"fmt"
	"math/rand/v2"
)

type Database[T Identifiable] struct {
	data []T
}

type Identifiable interface {
	GetId() int
	SetId(int)
}


func (d *Database[T]) Add(data T) {
	id := rand.IntN(9999999)
	data.SetId(id)
	if _, err := d.Get(id); err != nil  {
		d.data = append(d.data, data)
	}
}

func  (d *Database[T]) Delete(id int) {
	for i, p := range d.data {
		if p.GetId() == id {
			d.data = append(d.data[:i], d.data[i+1:]...)
		}
	}
}

func  (d *Database[T]) Update(data T) {
	for i, p := range d.data {
		if p.GetId() == data.GetId() {
			d.data[i] = data
		}
	}
}

func  (d *Database[T]) Get(id int) (ret T, err error){
	for _, p := range d.data {
		if p.GetId() == id {
			return p, nil
		}
	}	
	return ret, fmt.Errorf("erro")
}

func  (d *Database[T]) GetAll() []T {
	return d.data
}
