package structural

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

type DressFactory struct {
	dressMap map[string]Dress
}

var dressFactoryInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	switch dressType {
	case TerroristDressType:
		return d.dressMap[TerroristDressType], nil
	case CounterTerroristDressType:
		return d.dressMap[CounterTerroristDressType], nil
	default:
		return nil, fmt.Errorf("Wrong dressType passed")
	}
}

func getDressFactoryInstance() *DressFactory {
	return dressFactoryInstance
}
