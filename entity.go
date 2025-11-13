package main

import (
	"codeberg.org/anaseto/gruid"
	"codeberg.org/anaseto/gruid/rl"
)

// TODO: I don't understand almost anything about this file

type ECS struct {
	Entities  []Entity            // list of entities
	Positions map[int]gruid.Point // entity index: map position
	PlayerID  int                 // index of Player's entity (for convenience)
	Fighter   map[int]*fighter    // figthing component
	AI        map[int]*AI         // AI component
	Name      map[int]string      // name component
}

func NewECS() *ECS {
	return &ECS{
		Positions: map[int]gruid.Point{},
		Fighter:   map[int]*fighter{},
		AI:        map[int]*AI{},
		Name:      map[int]string{},
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

func (es *ECS) MonsterAt(p gruid.Point) (int, *Monster) {
	for i, q := range es.Positions {
		if p != q || !es.Alive(i) {
			continue
		}
		e := es.Entities[i]
		switch e := e.(type) {
		case *Monster:
			return i, e
		}
	}
	return -1, nil
}

func (es *ECS) NoBlockingEntityAt(p gruid.Point) bool {
	i, _ := es.MonsterAt(p)
	return es.Positions[es.PlayerID] != p && !es.Alive(i)
}

// PlayerDied checks whether the player died.
func (es *ECS) PlayerDied() bool {
	return es.Dead(es.PlayerID)
}

// Alive checks whether an entity is alive.
func (es *ECS) Alive(i int) bool {
	fi := es.Fighter[i]
	return fi != nil && fi.HP > 0
}

// Dead checks whether an entity is dead (was alive).
func (es *ECS) Dead(i int) bool {
	fi := es.Fighter[i]
	return fi != nil && fi.HP <= 0
}

// Style returns the graphical representation (rune and foreground color) of an
// entity.
func (es *ECS) Style(i int) (r rune, c gruid.Color) {
	r = es.Entities[i].Rune()
	c = es.Entities[i].Color()
	if es.Dead(i) {
		// Alternate representation for corpses of dead monsters.
		r = '%'
		c = gruid.ColorDefault
	}
	return r, c
}

// renderOrder is a type representing the priority of an entity rendering.
type renderOrder int

// Those constants represent distinct kinds of rendering priorities. In case
// two entities are at a given position, only the one with the highest priority
// gets displayed.
const (
	RONone renderOrder = iota
	ROCorpse
	ROItem
	ROActor
)

// RenderOrder returns the rendering priority of an entity.
func (es *ECS) RenderOrder(i int) (ro renderOrder) {
	switch es.Entities[i].(type) {
	case *Player:
		ro = ROActor
	case *Monster:
		if es.Dead(i) {
			ro = ROCorpse
		} else {
			ro = ROActor
		}
	}
	return ro
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

// Monster represents a monster. It implements the Entity interface.
type Monster struct {
	Char rune // monster's graphical representation
}

func (m *Monster) Rune() rune {
	return m.Char
}

func (m *Monster) Color() gruid.Color {
	return ColorMonster
}
