package main

import (
	"fmt"

	"codeberg.org/anaseto/gruid"
	"codeberg.org/anaseto/gruid/paths"
)

// action represents information relevant to the last UI action performed.
type action struct {
	Type  actionType  // kind of action (movement, quitting, ...)
	Delta gruid.Point // direction for ActionMovement
}

type actionType int

const (
	NoAction   actionType = iota
	ActionBump            // movement request
	ActionQuit            // quit the game
)

func (m *model) handleAction() gruid.Effect {
	switch m.action.Type {
	case ActionBump:
		np := m.game.ECS.Positions[m.game.ECS.PlayerID].Add(m.action.Delta)
		m.game.Bump(np)
	case ActionQuit:
		// for now, just terminate with gruid End command: this will
		// have to be updated later when implementing saving.
		return gruid.End()
	}
	return nil
}

func (g *game) Bump(to gruid.Point) {
	if !g.Map.Walkable(to) {
		return
	}
	if m := g.ECS.MonsterAt(to); m != nil {
		// We show a message to standard error. Later in the tutorial,
		// we'll put a message in the UI instead.
		fmt.Printf("You kick the %s, much to its annoyance!\n", m.Name)
		return
	}
	// We move the player to the new destination.
	g.ECS.MovePlayer(to)
	// Update FOV.
	g.UpdateFOV()
}

func (g *game) UpdateFOV() {
	player := g.ECS.Player()
	pp := g.ECS.Positions[g.ECS.PlayerID]
	rg := gruid.NewRange(-maxLOS, -maxLOS, maxLOS+1, maxLOS+1)
	player.FOV.SetRange(rg.Add(pp).Intersect(g.Map.Grid.Range()))
	passable := func(p gruid.Point) bool {
		return g.Map.Grid.At(p) != Wall
	}
	for _, p := range player.FOV.SSCVisionMap(pp, maxLOS, passable, false) {
		if paths.DistanceManhattan(p, pp) > maxLOS {
			continue
		}
		if !g.Map.Explored[p] {
			g.Map.Explored[p] = true
		}
	}
}

func (g *game) InFOV(p gruid.Point) bool {
	pp := g.ECS.Positions[g.ECS.PlayerID]
	return g.ECS.Player().FOV.Visible(p) &&
		paths.DistanceManhattan(pp, p) <= maxLOS
}
