package main

import "codeberg.org/anaseto/gruid"

type model struct {
	grid   gruid.Grid // drawing grid
	game   game       // game state
	action action     // UI action
}

type game struct {
	PlayerPos gruid.Point // tracks player position
	ECS *ECS // entities present on the map
	Map *Map // the game map, made of tiles
}

func (m *model) Update(msg gruid.Msg) gruid.Effect {
	m.action = action{} // reset last action information
	switch msg := msg.(type) {
	case gruid.MsgInit:
		// Initialize map
		size := m.grid.Size() // map size: for now the whole window
		m.game.Map = NewMap(size)
		// Initialize entities
		m.game.ECS = NewECS()
		m.game.ECS.PlayerID = m.game.ECS.AddEntity(&Player{}, size.Div(2))
		m.game.ECS.PlayerID = m.game.ECS.AddEntity(&Player{}, m.game.Map.RandomFloor())
		// Initialization: create a player entity centered on the map.
		m.game.ECS.PlayerID = m.game.ECS.AddEntity(&Player{}, size.Div(2))
	case gruid.MsgKeyDown:
		// Update action information on key down.
		m.updateMsgKeyDown(msg)
	}
	// Handle action (if any).
	return m.handleAction()
}

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
	// We draw the map tiles.
	it := m.game.Map.Grid.Iterator()
	for it.Next() {
		m.grid.Set(it.P(), gruid.Cell{Rune: m.game.Map.Rune(it.Cell())})
	}
	// We draw the entities.
	for i, e := range m.game.ECS.Entities {
		m.grid.Set(m.game.ECS.Positions[i], gruid.Cell{
			Rune:  e.Rune(),
			Style: gruid.Style{Fg: e.Color()},
		})
	}
	return m.grid
}
