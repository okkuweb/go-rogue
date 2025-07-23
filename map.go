package main

import (
	"codeberg.org/anaseto/gruid"
	"codeberg.org/anaseto/gruid/rl"
)

// These constants represent the different kind of map tiles.
const (
	Wall rl.Cell = iota
	Floor
)

// Map represents the rectangular map of the game's level.
type Map struct {
	Grid rl.Grid
}

// NewMap returns a new map with given size.
func NewMap(size gruid.Point) *Map {
	m := &Map{}
	m.Grid = rl.NewGrid(size.X, size.Y)
	m.Grid.Fill(Floor)

	m.Box(0, size.X, 0, size.Y, Wall)
	m.Box(1, size.X - 1, 1, size.Y - 1, Floor)
	m.Box(30, 33, 11, 14, Wall)

	return m
}

func (m *Map) Box(xMin, xMax, yMin, yMax int, c rl.Cell) {
	for x := xMin; x < xMax; x++ {
		for y := yMin; y < yMax; y++ {
			m.Grid.Set(gruid.Point{x, y}, c)
		}
	}
	return
}

// Walkable returns true if at the given position there is a floor tile.
func (m *Map) Walkable(p gruid.Point) bool {
	return m.Grid.At(p) == Floor
}

// Rune returns the character rune representing a given terrain.
func (m *Map) Rune(c rl.Cell) (r rune) {
	switch c {
	case Wall:
		r = '#'
	case Floor:
		r = '.'
	}
	return r
}
