package structural

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (g *game) addTerrorists() {
	player := NewPlayer("T", TerroristDressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *game) addCounterTerrorists() {
	player := NewPlayer("CT", CounterTerroristDressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}
