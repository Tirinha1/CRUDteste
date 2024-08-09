package database

import "math/rand/v2"

type Database[T Identifiable] struct {
	data []T
}

type Identifiable interface {
	GetId() int
	SetId(int)
}

func (d *Database[T]) AddProduct(product T) {
	id := rand.IntN(9999999)
	product.SetId(id)
	if GetProduct(id) == nil {
		d.data = append(d.data, product)
	}
}

func  (d *Database[T]) DelProduct(id int) {
	for i, p := range d.data {
		if p.GetId() == id {
			d.data = append(d.data[:i], d.data[i+1:]...)
		}
	}
}

func  (d *Database[T]) UpdtProduct(product T) {
	for i, p := range d.data {
		if p.GetId() == product.GetId() {
			d.data[i] = product
		}
	}
}

func  (d *Database[T]) GetProduct(id int) (ret T){
	for _, p := range d.data {
		if p.GetId() == id {

			return p
		}
	}
	return ret
}

func  (d *Database[T]) GetAll() []T {
	return d.data
}
