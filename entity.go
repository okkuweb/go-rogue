package main

import (
	"codeberg.org/anaseto/gruid"
	"codeberg.org/anaseto/gruid/rl"
)

type ECS struct {
	Entities  []Entity            // list of entities
	Positions map[int]gruid.Point // entity index: map position
	PlayerID  int                 // index of Player's entity (for convenience)
}

func NewECS() *ECS {
	return &ECS{
		Positions: map[int]gruid.Point{},
	}
}

func (es *ECS) AddEntity(e Entity, p gruid.Point) int {
	i := len(es.Entities)
	es.Entities = append(es.Entities, e)
	es.Positions[i] = p
	return i
}

func (es *ECS) MoveEntity(i int, p gruid.Point) {
	es.Positions[i] = p
}

func (es *ECS) MovePlayer(p gruid.Point) {
	es.MoveEntity(es.PlayerID, p)
}

func (es *ECS) Player() *Player {
	return es.Entities[es.PlayerID].(*Player)
}

func (es *ECS) MonsterAt(p gruid.Point) *Monster {
	for i, q := range es.Positions {
		if p != q {
			continue
		}
		e := es.Entities[i]
		switch e := e.(type) {
		case *Monster:
			return e
		}
	}
	return nil
}


func (es *ECS) NoBlockingEntityAt(p gruid.Point) bool {
	return es.Positions[es.PlayerID] != p && es.MonsterAt(p) == nil
}

type Entity interface {
	Rune() rune         // the character representing the entity
	Color() gruid.Color // the character's color
}

type Player struct {
	FOV *rl.FOV // player's field of view
}

const maxLOS = 10

func NewPlayer() *Player {
	player := &Player{}
	player.FOV = rl.NewFOV(gruid.NewRange(-maxLOS, -maxLOS, maxLOS+1, maxLOS+1))
	return player
}

func (p *Player) Rune() rune {
	return '@'
}

func (p *Player) Color() gruid.Color {
	return ColorPlayer
}

type Monster struct {
	Name string
	Char rune
}

func (m *Monster) Rune() rune {
	return m.Char
}

func (m *Monster) Color() gruid.Color {
	return ColorMonster
}
