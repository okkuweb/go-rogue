package main

import (
	"math/rand/v2"

	"codeberg.org/anaseto/gruid"
	"codeberg.org/anaseto/gruid/paths"
	"codeberg.org/anaseto/gruid/rl"
)

// TODO: I don't understand almost anything about this file

// These constants represent the different kind of map tiles.
const (
	Wall rl.Cell = iota
	Floor
)

// Map represents the rectangular map of the game's level.
type Map struct {
	Grid     rl.Grid
	Rand     *rand.Rand // random number generator
	Explored map[gruid.Point]bool
}

// NewMap returns a new map with given size.
func NewMap(size gruid.Point) *Map {
	m := &Map{
		Grid:     rl.NewGrid(size.X, size.Y),
		Rand:     rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
		Explored: make(map[gruid.Point]bool),
	}

	//m.Grid.Fill(Floor)
	//m.Box(0, size.X, 0, size.Y, Wall)
	//m.Box(1, size.X - 1, 1, size.Y - 1, Floor)
	//m.Box(30, 33, 11, 14, Wall)

	m.Generate()
	return m
}

func (m *Map) Box(xMin, xMax, yMin, yMax int, c rl.Cell) {
	for x := xMin; x < xMax; x++ {
		for y := yMin; y < yMax; y++ {
			m.Grid.Set(gruid.Point{X: x, Y: y}, c)
		}
	}
}

// Walkable returns true if at the given position there is a floor tile.
func (m *Map) Walkable(p gruid.Point) bool {
	return m.Grid.At(p) == Floor
}

// Rune returns the character rune representing a given terrain.
func (m *Map) Rune(c rl.Cell) (r rune) {
	switch c {
	case Wall:
		r = '█'
	case Floor:
		r = '.'
	}
	return r
}

type walker struct {
	rand *rand.Rand
}

func (w walker) Neighbor(p gruid.Point) gruid.Point {
	switch w.rand.IntN(6) {
	case 0, 1:
		return p.Shift(1, 0)
	case 2, 3:
		return p.Shift(-1, 0)
	case 4:
		return p.Shift(0, 1)
	default:
		return p.Shift(0, -1)
	}
}

// Generate fills the Grid attribute of m with a procedurally generated map.
func (m *Map) Generate() {
	// map generator using the rl package from gruid
	// cellular automata map generation with rules that give a cave-like map.
	for {
		sz := m.Grid.Size()
		w := ((37 + rand.IntN(6)) * sz.X) / MapWidth
		size := sz.Y * w
		mgen := rl.MapGen{
			Rand: m.Rand,
			Grid: m.Grid,
		}
		wlk := walker{rand: m.Rand}
		walks := ((7 + m.Rand.IntN(3)) * sz.X) / MapWidth
		mgen.RandomWalkCave(wlk, rl.Cell(Floor), float64(size)/float64(sz.X*sz.Y), walks)
		freep := m.RandomFloor()
		// We put walls in floor cells non reachable from freep, to ensure that
		// all the cells are connected (which is not guaranteed by cellular
		// automata map generation).
		pr := paths.NewPathRange(m.Grid.Range())
		pr.CCMap(&path{m: m}, freep)
		ntiles := mgen.KeepCC(pr, freep, Wall)
		const minCaveSize = 400
		if ntiles > minCaveSize {
			break
		}
		// If there were not enough free tiles, we run the map
		// generation again.
	}
}

// RandomFloor returns a random floor cell in the map. It assumes that such a
// floor cell exists (otherwise the function does not end).
func (m *Map) RandomFloor() gruid.Point {
	size := m.Grid.Size()
	for {
		freep := gruid.Point{m.Rand.IntN(size.X), m.Rand.IntN(size.Y)}
		if m.Grid.At(freep) == Floor {
			return freep
		}
	}
}

// path implements the paths.Pather interface and is used to provide pathing
// information in map generation.
type path struct {
	m  *Map
	nb paths.Neighbors
}

// Neighbors returns the list of walkable neighbors of q in the map using 4-way
// movement along cardinal directions.
func (p *path) Neighbors(q gruid.Point) []gruid.Point {
	return p.nb.Cardinal(q,
		func(r gruid.Point) bool { return p.m.Walkable(r) })
}
