package models

type Player struct {
	Name string
	Amount int32
}

func NewPlayer(name string) *Player {
	return &Player{name, 100}
}

func (p *Player) Bet(sum int32) int32 {
	p.Amount -= sum

	return sum
}
