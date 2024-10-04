package creational

import "fmt"

// gun interface
type IGun interface {
	setName(name string)
	setAmmo(ammo string)
	getName() string
	getAmmo() string
}

// gun struct
type Gun struct {
	name string
	ammo string
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setAmmo(ammo string) {
	g.ammo = ammo
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getAmmo() string {
	return g.ammo
}

// concrete gun
type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name: "AK47",
			ammo: "7.62mm",
		},
	}
}

type M416 struct {
	Gun
}

func newM416() IGun {
	return &M416{
		Gun: Gun{
			name: "M416",
			ammo: "5.56mm",
		},
	}
}

// gun factory
func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "AK47":
		return newAK47(), nil
	case "M416":
		return newM416(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}
