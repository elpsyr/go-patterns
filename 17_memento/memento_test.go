package memento

import "testing"

func TestGame_Play(t *testing.T) {

	game := Game{
		hp: 10,
		mp: 10,
	}
	game.Status()

	game.Play(-2, -3)
	save := game.Save()
	game.Load(save)

	game.Status()

}
