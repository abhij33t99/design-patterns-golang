package structural

type IPizza interface {
	GetPrice() int
}

type VegPizza struct{}

func (v *VegPizza) GetPrice() int {
	return 15
}

type NonVegPizza struct{}

func (nv *NonVegPizza) GetPrice() int {
	return 20
}
