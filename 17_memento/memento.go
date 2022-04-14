package memento

import "fmt"

type Memento interface {
}
type Game struct {
	hp, mp int
}

type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.hp += hpDelta
	g.mp += mpDelta

}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.hp = gm.hp
	g.mp = gm.mp

}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d ,MP:%d\n", g.hp, g.mp)

}
