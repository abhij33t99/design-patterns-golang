package creational

// lets build a car
type Car struct {
	color      string
	engine     string
	hasSunroof bool
	hasStereo  bool
}

type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngine(engine string) CarBuilder
	SetSunroof(hasSunroof bool) CarBuilder
	SetStereo(hasStereo bool) CarBuilder
	Build() *Car
}

type carBuilder struct {
	car *Car
}

func NewCarBuilder() carBuilder {
	return carBuilder{
		car: &Car{},
	}
}

func (cb *carBuilder) SetColor(color string) *carBuilder {
	cb.car.color = color
	return cb
}

func (cb *carBuilder) SetEngine(engine string) *carBuilder {
	cb.car.engine = engine
	return cb
}

func (cb *carBuilder) SetSunroof(hasSunroof bool) *carBuilder {
	cb.car.hasSunroof = hasSunroof
	return cb
}

func (cb *carBuilder) SetStereo(hasStereo bool) *carBuilder {
	cb.car.hasStereo = hasStereo
	return cb
}

func (cb *carBuilder) Build() *Car {
	return cb.car
}
