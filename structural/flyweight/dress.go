package structural

type Dress interface{
	getColor() string
}

type TerroristDress struct{
	color string
}

func (t *TerroristDress) getColor() string {
    return t.color
}

func newTerroristDress() *TerroristDress {
    return &TerroristDress{color: "red"}
}

type CounterTerroristDress struct {
    color string
}

func (c *CounterTerroristDress) getColor() string {
    return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
    return &CounterTerroristDress{color: "green"}
}