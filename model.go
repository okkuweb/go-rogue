package main

import (
	"fmt"

	"codeberg.org/anaseto/gruid"
)

type model struct {
	gd gruid.Grid // user interface grid
	g *game
	opt *Options
	State string
}

func (md *model) setState (state string) error {
	if md.opt.States[state] {
		md.State = state
		return nil
	}
	return fmt.Errorf("No such state: %s", state)
}

func (md *model) Update(msg gruid.Msg) gruid.Effect {
	var effect gruid.Effect
    switch msg.(type) {
    case gruid.MsgInit:
        // Initialize your grid here if needed
		if _, ok := msg.(gruid.MsgInit); ok {
			md.Print("Init game")
			err := md.setState("menu")
			if err != nil {
				md.Print(err)
				return gruid.End()
			}
			return nil
		}
    case gruid.MsgKeyDown:
	 	effect = md.KeyDown(msg)
    default:
    }
	return effect
}
func (md *model) Draw() gruid.Grid {
	max := md.gd.Size()
	for x := range max.X {
		// Top and bottom borders
		md.gd.Set(gruid.Point{X: x, Y: 0}, gruid.Cell{Rune: '-'})
		md.gd.Set(gruid.Point{X: x, Y: max.Y - 1}, gruid.Cell{Rune: '-'})
	}
	
	for y := range max.Y {
		// Left and right borders
		md.gd.Set(gruid.Point{X: 0, Y: y}, gruid.Cell{Rune: '|'})
		md.gd.Set(gruid.Point{X: max.X - 1, Y: y}, gruid.Cell{Rune: '|'})
	}
	return md.gd
}
