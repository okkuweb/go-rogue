package main

import "codeberg.org/anaseto/gruid"

// action represents information relevant to the last UI action performed.
type action struct {
	Type  actionType  // kind of action (movement, quitting, ...)
	Delta gruid.Point // direction for ActionMovement
}

type actionType int

const (
	NoAction       actionType = iota
	ActionMovement            // movement request
	ActionQuit                // quit the game
)

func (m *model) handleAction() gruid.Effect {
	switch m.action.Type {
	case ActionMovement:
		np := m.game.ECS.Positions[m.game.ECS.PlayerID]
		np = np.Add(m.action.Delta)
		if m.game.Map.Walkable(np) {
			m.game.ECS.MovePlayer(np)
		}
	case ActionQuit:
		// for now, just terminate with gruid End command: this will
		// have to be updated later when implementing saving.
		return gruid.End()
	}
	return nil
}
