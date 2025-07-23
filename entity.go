package main

import "codeberg.org/anaseto/gruid"

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

type Entity interface {
	Rune() rune         // the character representing the entity
	Color() gruid.Color // the character's color
}

type Player struct{}

func (p *Player) Rune() rune {
	return '@'
}

func (p *Player) Color() gruid.Color {
	return gruid.ColorDefault
}

