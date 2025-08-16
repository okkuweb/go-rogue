package main

import (
	"sort"

	"codeberg.org/anaseto/gruid"
	"codeberg.org/anaseto/gruid/paths"
)

type model struct {
	grid   gruid.Grid // drawing grid
	game   *game      // game state
	action action     // UI action
}

func (m *model) Update(msg gruid.Msg) gruid.Effect {
	m.action = action{} // reset last action information
	switch msg := msg.(type) {
	case gruid.MsgInit:
		m.game = &game{}
		// Initialize map
		size := m.grid.Size() // map size: for now the whole window
		m.game.Map = NewMap(size)
		m.game.PR = paths.NewPathRange(gruid.NewRange(0, 0, size.X, size.Y))
		// Initialize entities
		m.game.ECS = NewECS()
		m.game.ECS.PlayerID = m.game.ECS.AddEntity(NewPlayer(), m.game.Map.RandomFloor())
		m.game.ECS.Fighter[m.game.ECS.PlayerID] = &fighter{
			HP: 30, MaxHP: 30, Power: 5, Defense: 2,
		}
		m.game.ECS.Name[m.game.ECS.PlayerID] = "you"
		m.game.UpdateFOV()
		m.game.SpawnMonsters()
	case gruid.MsgKeyDown:
		// Update action information on key down.
		m.updateMsgKeyDown(msg)
	}
	// Handle action (if any).
	return m.handleAction()
}

const (
	ColorFOV gruid.Color = iota + 1
	ColorPlayer
	ColorMonster
)

func (m *model) updateMsgKeyDown(msg gruid.MsgKeyDown) {
	pdelta := gruid.Point{}
	switch msg.Key {

	case "h":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(-1, 0)}
	case "j":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(0, 1)}
	case "k":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(0, -1)}
	case "l":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(1, 0)}
	case "y":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(-1, -1)}
	case "u":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(1, -1)}
	case "b":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(-1, 1)}
	case "n":
		m.action = action{Type: ActionBump, Delta: pdelta.Shift(1, 1)}
	case ".":
		m.action = action{Type: ActionWait}


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
	// We sort entity indexes using the render ordering.
	sortedEntities := make([]int, 0, len(g.ECS.Entities))
	for i := range g.ECS.Entities {
		sortedEntities = append(sortedEntities, i)
	}
	sort.Slice(sortedEntities, func(i, j int) bool {
		return g.ECS.RenderOrder(sortedEntities[i]) < g.ECS.RenderOrder(sortedEntities[j])
	})
	// We draw the sorted entities.
	for _, i := range sortedEntities {
		p := g.ECS.Positions[i]
		if !g.Map.Explored[p] || !g.InFOV(p) {
			continue
		}
		c := m.grid.At(p)
		c.Rune, c.Style.Fg = g.ECS.Style(i)
		m.grid.Set(p, c)
	}
	return m.grid
}
