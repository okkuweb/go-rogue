package main

import "codeberg.org/anaseto/gruid"

type model struct {
	grid   gruid.Grid // drawing grid
	game   *game      // game state
	action action     // UI action
}

type game struct {
	PlayerPos gruid.Point // tracks player position
	ECS       *ECS        // entities present on the map
	Map       *Map        // the game map, made of tiles
}

func (m *model) Update(msg gruid.Msg) gruid.Effect {
	m.action = action{} // reset last action information
	switch msg := msg.(type) {
	case gruid.MsgInit:
		m.game = &game{}
		// Initialize map
		size := m.grid.Size() // map size: for now the whole window
		m.game.Map = NewMap(size)
		// Initialize entities
		m.game.ECS = NewECS()
		m.game.ECS.PlayerID = m.game.ECS.AddEntity(&Player{}, m.game.Map.RandomFloor())
		m.game.ECS.PlayerID = m.game.ECS.AddEntity(NewPlayer(), m.game.Map.RandomFloor())
		m.game.UpdateFOV()
		// Initialization: create a player entity centered on the map.
		m.game.ECS.PlayerID = m.game.ECS.AddEntity(&Player{}, size.Div(2))
	case gruid.MsgKeyDown:
		// Update action information on key down.
		m.updateMsgKeyDown(msg)
	}
	// Handle action (if any).
	return m.handleAction()
}

const (
	ColorFOV gruid.Color = iota + 1
)

func (m *model) updateMsgKeyDown(msg gruid.MsgKeyDown) {
	pdelta := gruid.Point{}
	switch msg.Key {

	case "h":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(-1, 0)}
	case "j":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(0, 1)}
	case "k":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(0, -1)}
	case "l":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(1, 0)}

	case "y":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(-1, -1)}
	case "u":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(1, -1)}
	case "b":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(-1, 1)}
	case "n":
		m.action = action{Type: ActionMovement, Delta: pdelta.Shift(1, 1)}

	case gruid.KeyEscape, "Q":
		m.action = action{Type: ActionQuit}
	}
}

func (m *model) Draw() gruid.Grid {
	m.grid.Fill(gruid.Cell{Rune: ' '})
	g := m.game
	// We draw the map tiles.
	it := g.Map.Grid.Iterator()
	for it.Next() {
		if !g.Map.Explored[it.P()] {
			continue
		}
		c := gruid.Cell{Rune: g.Map.Rune(it.Cell())}
		if g.InFOV(it.P()) {
			c.Style.Bg = ColorFOV
		}
		m.grid.Set(it.P(), c)
	}
	// We draw the entities.
	for i, e := range g.ECS.Entities {
		p := g.ECS.Positions[i]
		if !g.Map.Explored[p] || !g.InFOV(p) {
			continue
		}
		c := m.grid.At(p)
		c.Rune = e.Rune()
		c.Style.Fg = e.Color()
		m.grid.Set(p, c)
	}
	return m.grid
}
