package main

import "codeberg.org/anaseto/gruid"

type model struct {
	grid   gruid.Grid // drawing grid
	game   game       // game state
	action action     // UI action
}

type game struct {
	PlayerPos gruid.Point // tracks player position
}

func (m *model) Update(msg gruid.Msg) gruid.Effect {
	m.action = action{} // reset last action information
	switch msg := msg.(type) {
	case gruid.MsgInit:
		// Initialization: set player position in the center.
		m.game.PlayerPos = m.grid.Size().Div(2)
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
	it := m.grid.Iterator()
	for it.Next() {
		switch {
		case it.P() == m.game.PlayerPos:
			it.SetCell(gruid.Cell{Rune: '@'})
		default:
			it.SetCell(gruid.Cell{Rune: ' '})
		}
	}
	return m.grid
}
