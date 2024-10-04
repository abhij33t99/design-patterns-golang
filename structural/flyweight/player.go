package structural

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType string, dressType string) *Player {
	dress, _ := getDressFactoryInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) changeLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
