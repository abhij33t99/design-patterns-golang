package creational

import "fmt"

// lets create an abstract factory that return factory on basis of selected brand
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "Adidas":
		return &Adidas{}, nil // return adidas factory which creates adidas shoe and shirt
	case "Nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("Wrong brand selected!")
	}
}
