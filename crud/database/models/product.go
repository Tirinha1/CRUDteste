package models

type Product struct {
	ID                 int
	Name               string
	ManufactureCountry string
	Price              int
}

func (p *Product) GetId() int{
	return p.ID
}

func (p *Product) SetId(x int){
	p.ID = x
}


