package main

import (
	"codeberg.org/anaseto/gruid"
)

type model struct {
	gd gruid.Grid // user interface grid
	g *game
	state string
	// other fields with the state of the application
}
func (md *model) setState (state string) {
	md.state = state
}
func (md *model) getState (state string) {
	md.state = state
}


func (md *model) Update(msg gruid.Msg) gruid.Effect {
    switch msg.(type) {
    case gruid.MsgInit:
        // Initialize your grid here if needed
		Print("hey")
		md.setState("menu")
        return nil
    case gruid.MsgKeyDown:
		Print("bye")
        return gruid.End()
    default:
    }
    return nil
}
func (md *model) Draw() gruid.Grid {
	Print("bye")
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
